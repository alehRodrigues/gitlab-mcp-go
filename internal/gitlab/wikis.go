package gitlab

import (
	"net/url"
	"strconv"
)

// ListWikiPages lists wiki pages in a project
func (c *Client) ListWikiPages(projectID string, withContent bool, page, perPage int) ([]WikiPage, error) {
	projectID = c.DecodeProjectID(projectID)
	params := url.Values{}
	if withContent {
		params.Set("with_content", "true")
	}
	if page > 0 {
		params.Set("page", strconv.Itoa(page))
	}
	if perPage > 0 {
		params.Set("per_page", strconv.Itoa(perPage))
	}

	var pages []WikiPage
	if err := c.Get("/projects/"+url.PathEscape(projectID)+"/wikis", params, &pages); err != nil {
		return nil, err
	}
	return pages, nil
}

// GetWikiPage gets a specific wiki page
func (c *Client) GetWikiPage(projectID, slug string) (*WikiPage, error) {
	projectID = c.DecodeProjectID(projectID)
	path := "/projects/" + url.PathEscape(projectID) + "/wikis/" + url.PathEscape(slug)
	var page WikiPage
	if err := c.Get(path, nil, &page); err != nil {
		return nil, err
	}
	return &page, nil
}

// CreateWikiPage creates a new wiki page
func (c *Client) CreateWikiPage(projectID, title, content, format string) (*WikiPage, error) {
	projectID = c.DecodeProjectID(projectID)
	body := map[string]any{
		"title":   title,
		"content": content,
	}
	if format != "" {
		body["format"] = format
	}

	var page WikiPage
	if err := c.Post("/projects/"+url.PathEscape(projectID)+"/wikis", body, &page); err != nil {
		return nil, err
	}
	return &page, nil
}

// UpdateWikiPage updates an existing wiki page
func (c *Client) UpdateWikiPage(projectID, slug, title, content, format string) (*WikiPage, error) {
	projectID = c.DecodeProjectID(projectID)
	body := map[string]any{}
	if title != "" {
		body["title"] = title
	}
	if content != "" {
		body["content"] = content
	}
	if format != "" {
		body["format"] = format
	}

	var page WikiPage
	if err := c.Put("/projects/"+url.PathEscape(projectID)+"/wikis/"+url.PathEscape(slug), body, &page); err != nil {
		return nil, err
	}
	return &page, nil
}

// DeleteWikiPage deletes a wiki page
func (c *Client) DeleteWikiPage(projectID, slug string) error {
	projectID = c.DecodeProjectID(projectID)
	return c.Delete("/projects/" + url.PathEscape(projectID) + "/wikis/" + url.PathEscape(slug))
}
