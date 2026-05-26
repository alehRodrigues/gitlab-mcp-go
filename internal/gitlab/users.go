package gitlab

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
)

// GetUsers retrieves user details by usernames
func (c *Client) GetUsers(usernames []string) (UsersResponse, error) {
	result := make(UsersResponse)

	for _, username := range usernames {
		user, err := c.getUser(username)
		if err != nil {
			result[username] = nil
		} else {
			result[username] = user
		}
	}

	return result, nil
}

func (c *Client) getUser(username string) (*User, error) {
	params := url.Values{}
	params.Set("username", username)

	req, err := c.NewRequest("GET", "/users", params, nil)
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

	var users []User
	if err := json.Unmarshal(body, &users); err != nil {
		return nil, err
	}

	for _, u := range users {
		if u.Username == username {
			return &u, nil
		}
	}

	return nil, nil
}
