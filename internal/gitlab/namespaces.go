package gitlab

import (
	"net/url"
	"strconv"
)

// ListNamespaces lists all namespaces
func (c *Client) ListNamespaces(search string, owned bool, page, perPage int) ([]Namespace, error) {
	params := url.Values{}
	if search != "" {
		params.Set("search", search)
	}
	if owned {
		params.Set("owned", "true")
	}
	if page > 0 {
		params.Set("page", strconv.Itoa(page))
	}
	if perPage > 0 {
		params.Set("per_page", strconv.Itoa(perPage))
	}

	var namespaces []Namespace
	if err := c.Get("/namespaces", params, &namespaces); err != nil {
		return nil, err
	}
	return namespaces, nil
}

// GetNamespace gets a namespace by ID or path
func (c *Client) GetNamespace(id string) (*Namespace, error) {
	path := "/namespaces/" + url.PathEscape(id)
	var namespace Namespace
	if err := c.Get(path, nil, &namespace); err != nil {
		return nil, err
	}
	return &namespace, nil
}

// VerifyNamespace checks if a namespace path exists
func (c *Client) VerifyNamespace(namespacePath string, parentID int) (*NamespaceExistsResponse, error) {
	params := url.Values{}
	if parentID > 0 {
		params.Set("parent_id", strconv.Itoa(parentID))
	}

	path := "/namespaces/" + url.PathEscape(namespacePath) + "/exists"
	var resp NamespaceExistsResponse
	if err := c.Get(path, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
