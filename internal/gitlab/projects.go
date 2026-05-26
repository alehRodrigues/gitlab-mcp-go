package gitlab

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

// SearchProjects searches GitLab projects
func (c *Client) SearchProjects(query string, page, perPage int) (*SearchResponse, error) {
	params := url.Values{}
	params.Set("search", query)
	params.Set("page", strconv.Itoa(page))
	params.Set("per_page", strconv.Itoa(perPage))
	params.Set("order_by", "id")
	params.Set("sort", "desc")

	var projects []Project
	resp, err := c.doWithHeaders("/projects", params, &projects)
	if err != nil {
		return nil, err
	}

	totalCount := resp.Header.Get("x-total")
	totalPages := resp.Header.Get("x-total-pages")

	count := len(projects)
	if totalCount != "" {
		count, _ = strconv.Atoi(totalCount)
	}
	tp := 1
	if totalPages != "" {
		tp, _ = strconv.Atoi(totalPages)
	} else if perPage > 0 {
		tp = (count + perPage - 1) / perPage
	}

	return &SearchResponse{
		Count:       count,
		TotalPages:  tp,
		CurrentPage: page,
		Items:       projects,
	}, nil
}

type rawResponse struct {
	*http.Response
}

func (c *Client) doWithHeaders(path string, params url.Values, out any) (*rawResponse, error) {
	u := c.baseURL + path
	if len(params) > 0 {
		u += "?" + params.Encode()
	}
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if err := checkError(resp); err != nil {
		return nil, err
	}

	if out != nil {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("read response: %w", err)
		}
		if err := json.Unmarshal(body, out); err != nil {
			return nil, fmt.Errorf("decode response: %w", err)
		}
	}

	return &rawResponse{Response: resp}, nil
}

// CreateRepository creates a new GitLab project
func (c *Client) CreateRepository(name, description, visibility string, initReadme bool) (*Project, error) {
	body := map[string]any{
		"name":                 name,
		"default_branch":       "main",
		"path":                 name,
		"initialize_with_readme": initReadme,
	}
	if description != "" {
		body["description"] = description
	}
	if visibility != "" {
		body["visibility"] = visibility
	}
	var project Project
	if err := c.Post("/projects", body, &project); err != nil {
		return nil, err
	}
	return &project, nil
}

// GetProject gets a single project by ID
func (c *Client) GetProject(projectID string) (*Project, error) {
	projectID = c.DecodeProjectID(projectID)
	path := "/projects/" + url.PathEscape(projectID)
	var project Project
	if err := c.Get(path, nil, &project); err != nil {
		return nil, err
	}
	return &project, nil
}

// ListProjects lists projects accessible by the user
func (c *Client) ListProjects(params url.Values) ([]Project, error) {
	var projects []Project
	if err := c.Get("/projects", params, &projects); err != nil {
		return nil, err
	}
	return projects, nil
}

// ForkProject forks a project to a namespace
func (c *Client) ForkProject(projectID, namespace string) (*Fork, error) {
	projectID = c.DecodeProjectID(projectID)
	params := url.Values{}
	if namespace != "" {
		params.Set("namespace", namespace)
	}
	var fork Fork
	if err := c.PostForm("/projects/"+url.PathEscape(projectID)+"/fork", params, nil, &fork); err != nil {
		return nil, err
	}
	return &fork, nil
}

// ListGroupProjects lists projects in a group
func (c *Client) ListGroupProjects(groupID string, params url.Values) ([]Project, error) {
	groupID = c.DecodeProjectID(groupID)
	path := "/groups/" + url.PathEscape(groupID) + "/projects"
	var projects []Project
	if err := c.Get(path, params, &projects); err != nil {
		return nil, err
	}
	return projects, nil
}
