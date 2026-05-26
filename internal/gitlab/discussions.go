package gitlab

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"strings"
)

// ListMergeRequestDiscussions lists discussions for a merge request
func (c *Client) ListMergeRequestDiscussions(projectID string, mrIID int, page, perPage int) (*PaginatedDiscussions, error) {
	return c.listDiscussions(projectID, "merge_requests", mrIID, page, perPage)
}

// ListIssueDiscussions lists discussions for an issue
func (c *Client) ListIssueDiscussions(projectID string, issueIID int, page, perPage int) (*PaginatedDiscussions, error) {
	return c.listDiscussions(projectID, "issues", issueIID, page, perPage)
}

func (c *Client) listDiscussions(projectID string, resourceType string, resourceIID int, page, perPage int) (*PaginatedDiscussions, error) {
	projectID = c.DecodeProjectID(projectID)
	params := url.Values{}
	if page > 0 {
		params.Set("page", strconv.Itoa(page))
	}
	if perPage > 0 {
		params.Set("per_page", strconv.Itoa(perPage))
	}

	path := fmt.Sprintf("/projects/%s/%s/%d/discussions",
		url.PathEscape(projectID), resourceType, resourceIID)

	req, err := c.NewRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if err := checkError(resp); err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	var discussions []Discussion
	if err := json.Unmarshal(body, &discussions); err != nil {
		return nil, err
	}

	pagination := Pagination{
		NextPage:   parseIntHeader(resp.Header.Get("x-next-page")),
		Page:       parseIntHeaderDefault(resp.Header.Get("x-page")),
		PerPage:    parseIntHeaderDefault(resp.Header.Get("x-per-page")),
		PrevPage:   parseIntHeader(resp.Header.Get("x-prev-page")),
		Total:      parseIntHeader(resp.Header.Get("x-total")),
		TotalPages: parseIntHeader(resp.Header.Get("x-total-pages")),
	}

	return &PaginatedDiscussions{
		Items:      discussions,
		Pagination: pagination,
	}, nil
}

func parseIntHeader(s string) *int {
	if s == "" {
		return nil
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return nil
	}
	return &v
}

func parseIntHeaderDefault(s string) int {
	if s == "" {
		return 0
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return v
}

// CreateMergeRequestThread creates a new thread on a merge request
func (c *Client) CreateMergeRequestThread(projectID string, mrIID int, body string, position *NotePosition, createdAt string) (*Discussion, error) {
	projectID = c.DecodeProjectID(projectID)

	payload := map[string]any{
		"body": body,
	}
	if position != nil {
		payload["position"] = position
	}
	if createdAt != "" {
		payload["created_at"] = createdAt
	}

	path := fmt.Sprintf("/projects/%s/merge_requests/%d/discussions",
		url.PathEscape(projectID), mrIID)

	var discussion Discussion
	req, err := c.NewRequest("POST", path, nil, strings.NewReader(toJSON(payload)))
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()
	if err := checkError(resp); err != nil {
		return nil, err
	}
	if err := json.NewDecoder(resp.Body).Decode(&discussion); err != nil {
		return nil, err
	}
	return &discussion, nil
}

// CreateNote creates a simple note on an issue or merge request
func (c *Client) CreateNote(projectID, noteableType string, noteableIID int, noteBody string) (map[string]any, error) {
	projectID = c.DecodeProjectID(projectID)
	path := fmt.Sprintf("/projects/%s/%ss/%d/notes",
		url.PathEscape(projectID), noteableType, noteableIID)

	payload := map[string]string{"body": noteBody}

	var result map[string]any
	if err := c.Post(path, payload, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateMergeRequestNote updates a note in a merge request thread
func (c *Client) UpdateMergeRequestNote(projectID string, mrIID int, discussionID string, noteID int, body string, resolved *bool) (*Note, error) {
	projectID = c.DecodeProjectID(projectID)
	path := fmt.Sprintf("/projects/%s/merge_requests/%d/discussions/%s/notes/%d",
		url.PathEscape(projectID), mrIID, url.PathEscape(discussionID), noteID)

	payload := map[string]any{}
	if body != "" {
		payload["body"] = body
	}
	if resolved != nil {
		payload["resolved"] = *resolved
	}

	var note Note
	if err := c.Put(path, payload, &note); err != nil {
		return nil, err
	}
	return &note, nil
}

// CreateMergeRequestNote adds a note to an existing thread
func (c *Client) CreateMergeRequestNote(projectID string, mrIID int, discussionID string, noteBody, createdAt string) (*Note, error) {
	projectID = c.DecodeProjectID(projectID)
	path := fmt.Sprintf("/projects/%s/merge_requests/%d/discussions/%s/notes",
		url.PathEscape(projectID), mrIID, url.PathEscape(discussionID))

	payload := map[string]any{"body": noteBody}
	if createdAt != "" {
		payload["created_at"] = createdAt
	}

	var note Note
	if err := c.Post(path, payload, &note); err != nil {
		return nil, err
	}
	return &note, nil
}

// UpdateIssueNote updates a note in an issue discussion
func (c *Client) UpdateIssueNote(projectID string, issueIID int, discussionID string, noteID int, noteBody string) (*Note, error) {
	projectID = c.DecodeProjectID(projectID)
	path := fmt.Sprintf("/projects/%s/issues/%d/discussions/%s/notes/%d",
		url.PathEscape(projectID), issueIID, url.PathEscape(discussionID), noteID)

	payload := map[string]string{"body": noteBody}

	var note Note
	if err := c.Put(path, payload, &note); err != nil {
		return nil, err
	}
	return &note, nil
}

// CreateIssueNote adds a note to an issue discussion
func (c *Client) CreateIssueNote(projectID string, issueIID int, discussionID string, noteBody, createdAt string) (*Note, error) {
	projectID = c.DecodeProjectID(projectID)
	path := fmt.Sprintf("/projects/%s/issues/%d/discussions/%s/notes",
		url.PathEscape(projectID), issueIID, url.PathEscape(discussionID))

	payload := map[string]any{"body": noteBody}
	if createdAt != "" {
		payload["created_at"] = createdAt
	}

	var note Note
	if err := c.Post(path, payload, &note); err != nil {
		return nil, err
	}
	return &note, nil
}

func toJSON(v any) string {
	b, _ := json.Marshal(v)
	return string(b)
}
