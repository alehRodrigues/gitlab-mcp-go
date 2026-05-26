package gitlab

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// CreateMergeRequest creates a new merge request
func (c *Client) CreateMergeRequest(projectID, title, description, sourceBranch, targetBranch string, assigneeIDs, reviewerIDs []int, labels []string, allowCollaboration, draft bool) (*MergeRequest, error) {
	projectID = c.DecodeProjectID(projectID)

	body := map[string]any{
		"title":        title,
		"source_branch": sourceBranch,
		"target_branch": targetBranch,
	}
	if description != "" {
		body["description"] = description
	}
	if len(assigneeIDs) > 0 {
		body["assignee_ids"] = assigneeIDs
	}
	if len(reviewerIDs) > 0 {
		body["reviewer_ids"] = reviewerIDs
	}
	if len(labels) > 0 {
		body["labels"] = strings.Join(labels, ",")
	}
	if allowCollaboration {
		body["allow_collaboration"] = true
	}
	if draft {
		body["draft"] = true
	}

	var mr MergeRequest
	if err := c.Post("/projects/"+url.PathEscape(projectID)+"/merge_requests", body, &mr); err != nil {
		return nil, err
	}
	return &mr, nil
}

// ListMergeRequests lists merge requests with filtering
func (c *Client) ListMergeRequests(projectID string, params url.Values) ([]MergeRequest, error) {
	projectID = c.DecodeProjectID(projectID)
	path := "/projects/" + url.PathEscape(projectID) + "/merge_requests"
	var mrs []MergeRequest
	if err := c.Get(path, params, &mrs); err != nil {
		return nil, err
	}
	return mrs, nil
}

// GetMergeRequest gets a merge request by IID or source branch
func (c *Client) GetMergeRequest(projectID string, mrIID int, sourceBranch string) (*MergeRequest, error) {
	projectID = c.DecodeProjectID(projectID)

	if mrIID > 0 {
		path := "/projects/" + url.PathEscape(projectID) + "/merge_requests/" + strconv.Itoa(mrIID)
		var mr MergeRequest
		if err := c.Get(path, nil, &mr); err != nil {
			return nil, err
		}
		return &mr, nil
	}

	if sourceBranch != "" {
		params := url.Values{}
		params.Set("source_branch", sourceBranch)
		var mrs []MergeRequest
		path := "/projects/" + url.PathEscape(projectID) + "/merge_requests"
		if err := c.Get(path, params, &mrs); err != nil {
			return nil, err
		}
		if len(mrs) > 0 {
			return &mrs[0], nil
		}
		return nil, fmt.Errorf("no merge request found for branch %s", sourceBranch)
	}

	return nil, fmt.Errorf("either mergeRequestIid or sourceBranch must be provided")
}

// GetMergeRequestDiffs gets the diffs of a merge request
func (c *Client) GetMergeRequestDiffs(projectID string, mrIID int, sourceBranch, view string) ([]Diff, error) {
	projectID = c.DecodeProjectID(projectID)

	if mrIID <= 0 && sourceBranch != "" {
		mr, err := c.GetMergeRequest(projectID, 0, sourceBranch)
		if err != nil {
			return nil, err
		}
		mrIID = mr.IID
	}

	params := url.Values{}
	if view != "" {
		params.Set("view", view)
	}

	type changesResponse struct {
		Changes []Diff `json:"changes"`
	}

	var resp changesResponse
	if err := c.Get("/projects/"+url.PathEscape(projectID)+"/merge_requests/"+strconv.Itoa(mrIID)+"/changes", params, &resp); err != nil {
		return nil, err
	}
	return resp.Changes, nil
}

// UpdateMergeRequest updates a merge request
func (c *Client) UpdateMergeRequest(projectID string, mrIID int, sourceBranch string, body map[string]any) (*MergeRequest, error) {
	projectID = c.DecodeProjectID(projectID)

	if mrIID <= 0 && sourceBranch != "" {
		mr, err := c.GetMergeRequest(projectID, 0, sourceBranch)
		if err != nil {
			return nil, err
		}
		mrIID = mr.IID
	}

	path := "/projects/" + url.PathEscape(projectID) + "/merge_requests/" + strconv.Itoa(mrIID)
	var mr MergeRequest
	if err := c.Put(path, body, &mr); err != nil {
		return nil, err
	}
	return &mr, nil
}
