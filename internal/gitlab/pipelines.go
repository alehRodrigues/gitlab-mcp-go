package gitlab

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"strings"
)

// ListPipelines lists pipelines in a project
func (c *Client) ListPipelines(projectID string, params url.Values) ([]Pipeline, error) {
	projectID = c.DecodeProjectID(projectID)
	path := "/projects/" + url.PathEscape(projectID) + "/pipelines"
	var pipelines []Pipeline
	if err := c.Get(path, params, &pipelines); err != nil {
		return nil, err
	}
	return pipelines, nil
}

// GetPipeline gets a specific pipeline
func (c *Client) GetPipeline(projectID string, pipelineID int) (*Pipeline, error) {
	projectID = c.DecodeProjectID(projectID)
	path := "/projects/" + url.PathEscape(projectID) + "/pipelines/" + strconv.Itoa(pipelineID)
	var pipeline Pipeline
	if err := c.Get(path, nil, &pipeline); err != nil {
		return nil, fmt.Errorf("pipeline not found: %w", err)
	}
	return &pipeline, nil
}

// ListPipelineJobs lists jobs in a pipeline
func (c *Client) ListPipelineJobs(projectID string, pipelineID int, params url.Values) ([]PipelineJob, error) {
	projectID = c.DecodeProjectID(projectID)
	path := "/projects/" + url.PathEscape(projectID) + "/pipelines/" + strconv.Itoa(pipelineID) + "/jobs"
	var jobs []PipelineJob
	if err := c.Get(path, params, &jobs); err != nil {
		return nil, fmt.Errorf("pipeline not found: %w", err)
	}
	return jobs, nil
}

// GetPipelineJob gets a specific job
func (c *Client) GetPipelineJob(projectID string, jobID int) (*PipelineJob, error) {
	projectID = c.DecodeProjectID(projectID)
	path := "/projects/" + url.PathEscape(projectID) + "/jobs/" + strconv.Itoa(jobID)
	var job PipelineJob
	if err := c.Get(path, nil, &job); err != nil {
		return nil, fmt.Errorf("job not found: %w", err)
	}
	return &job, nil
}

// GetPipelineJobOutput gets the trace/output of a job
func (c *Client) GetPipelineJobOutput(projectID string, jobID int) (string, error) {
	projectID = c.DecodeProjectID(projectID)
	path := "/projects/" + url.PathEscape(projectID) + "/jobs/" + strconv.Itoa(jobID) + "/trace"
	return c.GetRaw(path, nil)
}

// CreatePipeline creates a new pipeline for a branch/tag
func (c *Client) CreatePipeline(projectID, ref string, variables []map[string]string) (*Pipeline, error) {
	projectID = c.DecodeProjectID(projectID)

	payload := map[string]any{
		"ref": ref,
	}
	if len(variables) > 0 {
		payload["variables"] = variables
	}

	b, _ := json.Marshal(payload)
	req, err := c.NewRequest("POST", "/projects/"+url.PathEscape(projectID)+"/pipeline", nil, strings.NewReader(string(b)))
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

	var pipeline Pipeline
	if err := json.Unmarshal(body, &pipeline); err != nil {
		return nil, err
	}
	return &pipeline, nil
}

// RetryPipeline retries a failed or canceled pipeline
func (c *Client) RetryPipeline(projectID string, pipelineID int) (*Pipeline, error) {
	projectID = c.DecodeProjectID(projectID)
	path := "/projects/" + url.PathEscape(projectID) + "/pipelines/" + strconv.Itoa(pipelineID) + "/retry"

	req, err := c.NewRequest("POST", path, nil, nil)
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

	var pipeline Pipeline
	if err := json.NewDecoder(resp.Body).Decode(&pipeline); err != nil {
		return nil, err
	}
	return &pipeline, nil
}

// CancelPipeline cancels a running pipeline
func (c *Client) CancelPipeline(projectID string, pipelineID int) (*Pipeline, error) {
	projectID = c.DecodeProjectID(projectID)
	path := "/projects/" + url.PathEscape(projectID) + "/pipelines/" + strconv.Itoa(pipelineID) + "/cancel"

	req, err := c.NewRequest("POST", path, nil, nil)
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

	var pipeline Pipeline
	if err := json.NewDecoder(resp.Body).Decode(&pipeline); err != nil {
		return nil, err
	}
	return &pipeline, nil
}
