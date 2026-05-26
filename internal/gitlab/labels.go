package gitlab

import (
	"net/url"
	"strconv"
)

// ListLabels lists labels for a project
func (c *Client) ListLabels(projectID string, withCounts *bool, includeAncestorGroups *bool, search string) ([]Label, error) {
	projectID = c.DecodeProjectID(projectID)
	params := url.Values{}
	if withCounts != nil {
		params.Set("with_counts", strconv.FormatBool(*withCounts))
	}
	if includeAncestorGroups != nil {
		params.Set("include_ancestor_groups", strconv.FormatBool(*includeAncestorGroups))
	}
	if search != "" {
		params.Set("search", search)
	}

	var labels []Label
	if err := c.Get("/projects/"+url.PathEscape(projectID)+"/labels", params, &labels); err != nil {
		return nil, err
	}
	return labels, nil
}

// GetLabel gets a single label
func (c *Client) GetLabel(projectID, labelID string, includeAncestorGroups bool) (*Label, error) {
	projectID = c.DecodeProjectID(projectID)
	params := url.Values{}
	if includeAncestorGroups {
		params.Set("include_ancestor_groups", "true")
	}

	var label Label
	if err := c.Get("/projects/"+url.PathEscape(projectID)+"/labels/"+url.PathEscape(labelID), params, &label); err != nil {
		return nil, err
	}
	return &label, nil
}

// CreateLabel creates a new label
func (c *Client) CreateLabel(projectID, name, color, description string, priority *int) (*Label, error) {
	projectID = c.DecodeProjectID(projectID)

	body := map[string]any{
		"name":  name,
		"color": color,
	}
	if description != "" {
		body["description"] = description
	}
	if priority != nil {
		body["priority"] = *priority
	}

	var label Label
	if err := c.Post("/projects/"+url.PathEscape(projectID)+"/labels", body, &label); err != nil {
		return nil, err
	}
	return &label, nil
}

// UpdateLabel updates an existing label
func (c *Client) UpdateLabel(projectID, labelID, newName, color, description string, priority *int) (*Label, error) {
	projectID = c.DecodeProjectID(projectID)

	body := map[string]any{}
	if newName != "" {
		body["new_name"] = newName
	}
	if color != "" {
		body["color"] = color
	}
	if description != "" {
		body["description"] = description
	}
	if priority != nil {
		body["priority"] = *priority
	}

	var label Label
	if err := c.Put("/projects/"+url.PathEscape(projectID)+"/labels/"+url.PathEscape(labelID), body, &label); err != nil {
		return nil, err
	}
	return &label, nil
}

// DeleteLabel deletes a label
func (c *Client) DeleteLabel(projectID, labelID string) error {
	projectID = c.DecodeProjectID(projectID)
	return c.Delete("/projects/" + url.PathEscape(projectID) + "/labels/" + url.PathEscape(labelID))
}
