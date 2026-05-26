package gitlab

import (
	"net/url"
	"strconv"
)

// CreateIssue creates a new issue
func (c *Client) CreateIssue(projectID, title, description string, assigneeIDs []int, labels []string, milestoneID int) (*Issue, error) {
	projectID = c.DecodeProjectID(projectID)

	body := map[string]any{
		"title": title,
	}
	if description != "" {
		body["description"] = description
	}
	if len(assigneeIDs) > 0 {
		body["assignee_ids"] = assigneeIDs
	}
	if len(labels) > 0 {
		body["labels"] = labels
	}
	if milestoneID > 0 {
		body["milestone_id"] = milestoneID
	}

	var issue Issue
	if err := c.Post("/projects/"+url.PathEscape(projectID)+"/issues", body, &issue); err != nil {
		return nil, err
	}
	return &issue, nil
}

// ListIssues lists issues with filtering
func (c *Client) ListIssues(projectID string, params url.Values) ([]Issue, error) {
	projectID = c.DecodeProjectID(projectID)
	path := "/projects/" + url.PathEscape(projectID) + "/issues"
	var issues []Issue
	if err := c.Get(path, params, &issues); err != nil {
		return nil, err
	}
	return issues, nil
}

// GetIssue gets a single issue
func (c *Client) GetIssue(projectID string, issueIID int) (*Issue, error) {
	projectID = c.DecodeProjectID(projectID)
	path := "/projects/" + url.PathEscape(projectID) + "/issues/" + strconv.Itoa(issueIID)
	var issue Issue
	if err := c.Get(path, nil, &issue); err != nil {
		return nil, err
	}
	return &issue, nil
}

// UpdateIssue updates an issue
func (c *Client) UpdateIssue(projectID string, issueIID int, body map[string]any) (*Issue, error) {
	projectID = c.DecodeProjectID(projectID)
	path := "/projects/" + url.PathEscape(projectID) + "/issues/" + strconv.Itoa(issueIID)

	// Convert labels array to comma-separated if present
	if labels, ok := body["labels"]; ok {
		if arr, ok := labels.([]string); ok {
			joined := ""
			for i, l := range arr {
				if i > 0 {
					joined += ","
				}
				joined += l
			}
			body["labels"] = joined
		}
	}

	var issue Issue
	if err := c.Put(path, body, &issue); err != nil {
		return nil, err
	}
	return &issue, nil
}

// DeleteIssue deletes an issue
func (c *Client) DeleteIssue(projectID string, issueIID int) error {
	projectID = c.DecodeProjectID(projectID)
	path := "/projects/" + url.PathEscape(projectID) + "/issues/" + strconv.Itoa(issueIID)
	return c.Delete(path)
}

// ListIssueLinks lists all issue links for an issue
func (c *Client) ListIssueLinks(projectID string, issueIID int) ([]IssueWithLinkDetails, error) {
	projectID = c.DecodeProjectID(projectID)
	path := "/projects/" + url.PathEscape(projectID) + "/issues/" + strconv.Itoa(issueIID) + "/links"
	var links []IssueWithLinkDetails
	if err := c.Get(path, nil, &links); err != nil {
		return nil, err
	}
	return links, nil
}

// GetIssueLink gets a specific issue link
func (c *Client) GetIssueLink(projectID string, issueIID, issueLinkID int) (*IssueLink, error) {
	projectID = c.DecodeProjectID(projectID)
	path := "/projects/" + url.PathEscape(projectID) + "/issues/" + strconv.Itoa(issueIID) + "/links/" + strconv.Itoa(issueLinkID)
	var link IssueLink
	if err := c.Get(path, nil, &link); err != nil {
		return nil, err
	}
	return &link, nil
}

// CreateIssueLink creates a link between two issues
func (c *Client) CreateIssueLink(projectID string, issueIID int, targetProjectID string, targetIssueIID int, linkType string) (*IssueLink, error) {
	projectID = c.DecodeProjectID(projectID)
	targetProjectID = c.DecodeProjectID(targetProjectID)

	body := map[string]any{
		"target_project_id": targetProjectID,
		"target_issue_iid":  targetIssueIID,
		"link_type":         linkType,
	}

	path := "/projects/" + url.PathEscape(projectID) + "/issues/" + strconv.Itoa(issueIID) + "/links"
	var link IssueLink
	if err := c.Post(path, body, &link); err != nil {
		return nil, err
	}
	return &link, nil
}

// DeleteIssueLink deletes an issue link
func (c *Client) DeleteIssueLink(projectID string, issueIID, issueLinkID int) error {
	projectID = c.DecodeProjectID(projectID)
	path := "/projects/" + url.PathEscape(projectID) + "/issues/" + strconv.Itoa(issueIID) + "/links/" + strconv.Itoa(issueLinkID)
	return c.Delete(path)
}
