package gitlab

import (
	"net/url"
)

func (c *Client) ListCommits(projectID string, params url.Values) ([]Commit, error) {
	projectID = c.DecodeProjectID(projectID)
	path := "/projects/" + url.PathEscape(projectID) + "/repository/commits"
	var commits []Commit
	if err := c.Get(path, params, &commits); err != nil {
		return nil, err
	}
	return commits, nil
}

func (c *Client) GetCommit(projectID, sha string, params url.Values) (*Commit, error) {
	projectID = c.DecodeProjectID(projectID)
	path := "/projects/" + url.PathEscape(projectID) + "/repository/commits/" + url.PathEscape(sha)
	var commit Commit
	if err := c.Get(path, params, &commit); err != nil {
		return nil, err
	}
	return &commit, nil
}

func (c *Client) GetCommitDiff(projectID, sha string, params url.Values) ([]Diff, error) {
	projectID = c.DecodeProjectID(projectID)
	path := "/projects/" + url.PathEscape(projectID) + "/repository/commits/" + url.PathEscape(sha) + "/diff"
	var diffs []Diff
	if err := c.Get(path, params, &diffs); err != nil {
		return nil, err
	}
	return diffs, nil
}

func (c *Client) GetCommitRefs(projectID, sha string, refType string) ([]CommitRef, error) {
	projectID = c.DecodeProjectID(projectID)
	params := url.Values{}
	if refType != "" {
		params.Set("type", refType)
	}
	path := "/projects/" + url.PathEscape(projectID) + "/repository/commits/" + url.PathEscape(sha) + "/refs"
	var refs []CommitRef
	if err := c.Get(path, params, &refs); err != nil {
		return nil, err
	}
	return refs, nil
}

func (c *Client) GetCommitComments(projectID, sha string) ([]CommitComment, error) {
	projectID = c.DecodeProjectID(projectID)
	path := "/projects/" + url.PathEscape(projectID) + "/repository/commits/" + url.PathEscape(sha) + "/comments"
	var comments []CommitComment
	if err := c.Get(path, nil, &comments); err != nil {
		return nil, err
	}
	return comments, nil
}

func (c *Client) CreateCommitComment(projectID, sha, note, path string, line *int, lineType string) (*CommitComment, error) {
	projectID = c.DecodeProjectID(projectID)
	body := map[string]any{"note": note}
	if path != "" {
		body["path"] = path
	}
	if line != nil {
		body["line"] = *line
	}
	if lineType != "" {
		body["line_type"] = lineType
	}
	apiPath := "/projects/" + url.PathEscape(projectID) + "/repository/commits/" + url.PathEscape(sha) + "/comments"
	var comment CommitComment
	if err := c.Post(apiPath, body, &comment); err != nil {
		return nil, err
	}
	return &comment, nil
}

func (c *Client) GetCommitMergeRequests(projectID, sha string, params url.Values) ([]MergeRequest, error) {
	projectID = c.DecodeProjectID(projectID)
	path := "/projects/" + url.PathEscape(projectID) + "/repository/commits/" + url.PathEscape(sha) + "/merge_requests"
	var mrs []MergeRequest
	if err := c.Get(path, params, &mrs); err != nil {
		return nil, err
	}
	return mrs, nil
}
