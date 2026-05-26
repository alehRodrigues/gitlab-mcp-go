package gitlab

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	baseURL    string
	token      string
	httpClient *http.Client
}

func NewClient(baseURL, token string) *Client {
	return &Client{
		baseURL:    strings.TrimRight(baseURL, "/"),
		token:      token,
		httpClient: &http.Client{},
	}
}

func (c *Client) NewRequest(method, path string, query url.Values, body io.Reader) (*http.Request, error) {
	u := c.baseURL + path
	if len(query) > 0 {
		u += "?" + query.Encode()
	}
	req, err := http.NewRequest(method, u, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	return req, nil
}

func (c *Client) Do(req *http.Request, out any) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if err := checkError(resp); err != nil {
		return nil, err
	}

	if out != nil {
		if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
			return nil, fmt.Errorf("decode response: %w", err)
		}
	}
	return resp, nil
}

func (c *Client) Get(path string, query url.Values, out any) error {
	req, err := c.NewRequest(http.MethodGet, path, query, nil)
	if err != nil {
		return err
	}
	_, err = c.Do(req, out)
	return err
}

func (c *Client) Post(path string, body any, out any) error {
	var r io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("marshal body: %w", err)
		}
		r = strings.NewReader(string(b))
	}
	req, err := c.NewRequest(http.MethodPost, path, nil, r)
	if err != nil {
		return err
	}
	_, err = c.Do(req, out)
	return err
}

func (c *Client) PostForm(path string, query url.Values, body any, out any) error {
	var r io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("marshal body: %w", err)
		}
		r = strings.NewReader(string(b))
	}
	req, err := c.NewRequest(http.MethodPost, path, query, r)
	if err != nil {
		return err
	}
	_, err = c.Do(req, out)
	return err
}

func (c *Client) Put(path string, body any, out any) error {
	b, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("marshal body: %w", err)
	}
	req, err := c.NewRequest(http.MethodPut, path, nil, strings.NewReader(string(b)))
	if err != nil {
		return err
	}
	_, err = c.Do(req, out)
	return err
}

func (c *Client) Delete(path string) error {
	req, err := c.NewRequest(http.MethodDelete, path, nil, nil)
	if err != nil {
		return err
	}
	_, err = c.Do(req, nil)
	return err
}

func (c *Client) GetRaw(path string, query url.Values) (string, error) {
	req, err := c.NewRequest(http.MethodGet, path, query, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "text/plain")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()
	if err := checkError(resp); err != nil {
		return "", err
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read body: %w", err)
	}
	return string(b), nil
}

func checkError(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}
	body, _ := io.ReadAll(resp.Body)
	msg := string(body)
	if msg == "" {
		msg = resp.Status
	}
	if resp.StatusCode == 403 && strings.Contains(msg, "User API Key Rate limit exceeded") {
		return fmt.Errorf("GitLab API Rate Limit Exceeded: %s", msg)
	}
	return fmt.Errorf("GitLab API error: %d %s\n%s", resp.StatusCode, resp.Status, msg)
}

func (c *Client) rawPost(path string, body any) ([]byte, *http.Response, error) {
	var r io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, nil, fmt.Errorf("marshal body: %w", err)
		}
		r = strings.NewReader(string(b))
	}
	req, err := c.NewRequest(http.MethodPost, path, nil, r)
	if err != nil {
		return nil, nil, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()
	if err := checkError(resp); err != nil {
		return nil, nil, err
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("read body: %w", err)
	}
	return b, resp, nil
}

func (c *Client) DecodeProjectID(id string) string {
	decoded, err := url.QueryUnescape(id)
	if err != nil {
		return id
	}
	return decoded
}
