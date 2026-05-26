package gitlab

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

// GetFileContents gets the contents of a file or directory
func (c *Client) GetFileContents(projectID, filePath, ref string) (any, error) {
	projectID = c.DecodeProjectID(projectID)
	encodedPath := url.PathEscape(filePath)

	params := url.Values{}
	if ref != "" {
		params.Set("ref", ref)
	} else {
		def, err := c.GetDefaultBranch(projectID)
		if err != nil {
			return nil, err
		}
		params.Set("ref", def)
	}

	path := "/projects/" + url.PathEscape(projectID) + "/repository/files/" + encodedPath
	req, err := c.NewRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("File not found: %s", filePath)
	}
	if err := checkError(resp); err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	// Try directory listing first
	var dir []DirectoryEntry
	if err := json.Unmarshal(bodyBytes, &dir); err == nil && len(dir) > 0 {
		return dir, nil
	}

	// Then try file content
	var file FileContent
	if err := json.Unmarshal(bodyBytes, &file); err != nil {
		return nil, fmt.Errorf("unmarshal file content: %w", err)
	}
	decoded, err := base64.StdEncoding.DecodeString(file.Content)
	if err == nil {
		file.Content = string(decoded)
		file.Encoding = "utf8"
	}
	return file, nil
}

// CreateOrUpdateFile creates or updates a single file
func (c *Client) CreateOrUpdateFile(projectID, filePath, content, commitMessage, branch, previousPath, lastCommitID, commitID string) (*CreateUpdateFileResponse, error) {
	projectID = c.DecodeProjectID(projectID)
	encodedPath := url.PathEscape(filePath)
	basePath := "/projects/" + url.PathEscape(projectID) + "/repository/files/" + encodedPath

	body := map[string]any{
		"branch":         branch,
		"content":        content,
		"commit_message": commitMessage,
		"encoding":       "text",
	}
	if previousPath != "" {
		body["previous_path"] = previousPath
	}

	// Check if file exists to determine PUT vs POST
	method := "POST"
	existing, err := c.GetFileContents(projectID, filePath, branch)
	if err == nil {
		method = "PUT"
		if f, ok := existing.(FileContent); ok {
			if commitID == "" && f.CommitID != "" {
				body["commit_id"] = f.CommitID
			}
			if lastCommitID == "" && f.LastCommitID != "" {
				body["last_commit_id"] = f.LastCommitID
			}
		}
	}
	if commitID != "" {
		body["commit_id"] = commitID
	}
	if lastCommitID != "" {
		body["last_commit_id"] = lastCommitID
	}

	b, _ := json.Marshal(body)
	req, err := c.NewRequest(method, basePath, nil, strings.NewReader(string(b)))
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		msg, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("GitLab API error: %d %s\n%s", resp.StatusCode, resp.Status, string(msg))
	}

	var result CreateUpdateFileResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PushFiles pushes multiple files in a single commit
func (c *Client) PushFiles(projectID, branch, commitMessage string, files []FileOperation) (*Commit, error) {
	projectID = c.DecodeProjectID(projectID)
	path := "/projects/" + url.PathEscape(projectID) + "/repository/commits"

	type action struct {
		Action   string `json:"action"`
		FilePath string `json:"file_path"`
		Content  string `json:"content"`
		Encoding string `json:"encoding"`
	}

	var actions []action
	for _, f := range files {
		actions = append(actions, action{
			Action:   "create",
			FilePath: f.FilePath,
			Content:  f.Content,
			Encoding: "text",
		})
	}

	body := map[string]any{
		"branch":         branch,
		"commit_message": commitMessage,
		"actions":        actions,
	}

	var commit Commit
	if err := c.Post(path, body, &commit); err != nil {
		return nil, err
	}
	return &commit, nil
}

// GetRepositoryTree gets the repository tree
func (c *Client) GetRepositoryTree(projectID, path, ref string, recursive bool, perPage int, pageToken, pagination string) ([]TreeItem, error) {
	projectID = c.DecodeProjectID(projectID)
	params := url.Values{}
	if path != "" {
		params.Set("path", path)
	}
	if ref != "" {
		params.Set("ref", ref)
	}
	if recursive {
		params.Set("recursive", "true")
	}
	if perPage > 0 {
		params.Set("per_page", strconv.Itoa(perPage))
	}
	if pageToken != "" {
		params.Set("page_token", pageToken)
	}
	if pagination != "" {
		params.Set("pagination", pagination)
	}

	var tree []TreeItem
	if err := c.Get("/projects/"+url.PathEscape(projectID)+"/repository/tree", params, &tree); err != nil {
		return nil, err
	}
	return tree, nil
}

// GetDefaultBranch returns the default branch name for a project
func (c *Client) GetDefaultBranch(projectID string) (string, error) {
	projectID = c.DecodeProjectID(projectID)
	var project Project
	if err := c.Get("/projects/"+url.PathEscape(projectID), nil, &project); err != nil {
		return "main", nil
	}
	if project.DefaultBranch != "" {
		return project.DefaultBranch, nil
	}
	return "main", nil
}

// CreateBranch creates a new branch
func (c *Client) CreateBranch(projectID, name, ref string) (*Reference, error) {
	projectID = c.DecodeProjectID(projectID)

	if ref == "" {
		def, err := c.GetDefaultBranch(projectID)
		if err != nil {
			return nil, err
		}
		ref = def
	}

	body := map[string]string{
		"branch": name,
		"ref":    ref,
	}

	var reference Reference
	if err := c.Post("/projects/"+url.PathEscape(projectID)+"/repository/branches", body, &reference); err != nil {
		return nil, err
	}
	return &reference, nil
}

// GetBranchDiffs compares two branches/commits
func (c *Client) GetBranchDiffs(projectID, from, to string, straight bool, excludedPatterns []string) (*CompareResult, error) {
	projectID = c.DecodeProjectID(projectID)
	params := url.Values{}
	params.Set("from", from)
	params.Set("to", to)
	if straight {
		params.Set("straight", "true")
	}

	path := "/projects/" + url.PathEscape(projectID) + "/repository/compare"
	req, err := c.NewRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		msg, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("GitLab API error: %d %s\n%s", resp.StatusCode, resp.Status, string(msg))
	}

	var result CompareResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if len(excludedPatterns) > 0 {
		var filtered []Diff
		for _, d := range result.Diffs {
			exclude := false
			for _, pattern := range excludedPatterns {
				matched, _ := regexp.MatchString(pattern, d.NewPath)
				if matched {
					exclude = true
					break
				}
			}
			if !exclude {
				filtered = append(filtered, d)
			}
		}
		result.Diffs = filtered
	}

	return &result, nil
}
