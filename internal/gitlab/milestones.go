package gitlab

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// ListProjectMilestones lists milestones in a project
func (c *Client) ListProjectMilestones(projectID string, params url.Values) ([]Milestone, error) {
	projectID = c.DecodeProjectID(projectID)
	path := "/projects/" + url.PathEscape(projectID) + "/milestones"
	var milestones []Milestone
	if err := c.Get(path, params, &milestones); err != nil {
		return nil, err
	}
	return milestones, nil
}

// GetProjectMilestone gets a specific milestone
func (c *Client) GetProjectMilestone(projectID string, milestoneID int) (*Milestone, error) {
	projectID = c.DecodeProjectID(projectID)
	path := "/projects/" + url.PathEscape(projectID) + "/milestones/" + strconv.Itoa(milestoneID)
	var milestone Milestone
	if err := c.Get(path, nil, &milestone); err != nil {
		return nil, err
	}
	return &milestone, nil
}

// CreateProjectMilestone creates a new milestone
func (c *Client) CreateProjectMilestone(projectID, title, description, dueDate, startDate string) (*Milestone, error) {
	projectID = c.DecodeProjectID(projectID)
	body := map[string]any{
		"title": title,
	}
	if description != "" {
		body["description"] = description
	}
	if dueDate != "" {
		body["due_date"] = dueDate
	}
	if startDate != "" {
		body["start_date"] = startDate
	}

	var milestone Milestone
	if err := c.Post("/projects/"+url.PathEscape(projectID)+"/milestones", body, &milestone); err != nil {
		return nil, err
	}
	return &milestone, nil
}

// EditProjectMilestone edits an existing milestone
func (c *Client) EditProjectMilestone(projectID string, milestoneID int, title, description, dueDate, startDate, stateEvent string) (*Milestone, error) {
	projectID = c.DecodeProjectID(projectID)
	body := map[string]any{}
	if title != "" {
		body["title"] = title
	}
	if description != "" {
		body["description"] = description
	}
	if dueDate != "" {
		body["due_date"] = dueDate
	}
	if startDate != "" {
		body["start_date"] = startDate
	}
	if stateEvent != "" {
		body["state_event"] = stateEvent
	}

	var milestone Milestone
	path := "/projects/" + url.PathEscape(projectID) + "/milestones/" + strconv.Itoa(milestoneID)
	if err := c.Put(path, body, &milestone); err != nil {
		return nil, err
	}
	return &milestone, nil
}

// DeleteProjectMilestone deletes a milestone
func (c *Client) DeleteProjectMilestone(projectID string, milestoneID int) error {
	projectID = c.DecodeProjectID(projectID)
	path := "/projects/" + url.PathEscape(projectID) + "/milestones/" + strconv.Itoa(milestoneID)
	return c.Delete(path)
}

// GetMilestoneIssues gets issues for a milestone
func (c *Client) GetMilestoneIssues(projectID string, milestoneID int) ([]Issue, error) {
	projectID = c.DecodeProjectID(projectID)
	path := "/projects/" + url.PathEscape(projectID) + "/milestones/" + strconv.Itoa(milestoneID) + "/issues"
	var issues []Issue
	if err := c.Get(path, nil, &issues); err != nil {
		return nil, err
	}
	return issues, nil
}

// GetMilestoneMergeRequests gets merge requests for a milestone
func (c *Client) GetMilestoneMergeRequests(projectID string, milestoneID int) ([]MergeRequest, error) {
	projectID = c.DecodeProjectID(projectID)
	path := "/projects/" + url.PathEscape(projectID) + "/milestones/" + strconv.Itoa(milestoneID) + "/merge_requests"
	var mrs []MergeRequest
	if err := c.Get(path, nil, &mrs); err != nil {
		return nil, err
	}
	return mrs, nil
}

// PromoteProjectMilestone promotes a milestone to group level
func (c *Client) PromoteProjectMilestone(projectID string, milestoneID int) (*Milestone, error) {
	projectID = c.DecodeProjectID(projectID)
	path := "/projects/" + url.PathEscape(projectID) + "/milestones/" + strconv.Itoa(milestoneID) + "/promote"

	b, _, err := c.rawPost(path, nil)
	if err != nil {
		return nil, err
	}

	var milestone Milestone
	if err := json.Unmarshal(b, &milestone); err != nil {
		return nil, err
	}
	return &milestone, nil
}

// GetMilestoneBurndownEvents gets burndown chart events for a milestone
func (c *Client) GetMilestoneBurndownEvents(projectID string, milestoneID int) ([]any, error) {
	projectID = c.DecodeProjectID(projectID)
	path := "/projects/" + url.PathEscape(projectID) + "/milestones/" + strconv.Itoa(milestoneID) + "/burndown_events"

	var events []any
	if err := c.Get(path, nil, &events); err != nil {
		return nil, err
	}
	return events, nil
}
