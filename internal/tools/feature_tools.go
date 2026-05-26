package tools

import (
	"context"
	"strconv"

	mcp_golang "github.com/metoro-io/mcp-golang"
)

// Wiki tool args
type ListWikiPagesArgs struct {
	ProjectID   string `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	WithContent *bool  `json:"with_content,omitempty" jsonschema:"description=Include content of the wiki pages"`
	Page        *int   `json:"page,omitempty" jsonschema:"description=Page number"`
	PerPage     *int   `json:"per_page,omitempty" jsonschema:"description=Number of items per page"`
}

type GetWikiPageArgs struct {
	ProjectID string `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	Slug      string `json:"slug" jsonschema:"required,description=URL-encoded slug of the wiki page"`
}

type CreateWikiPageArgs struct {
	ProjectID string  `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	Title     string  `json:"title" jsonschema:"required,description=Title of the wiki page"`
	Content   string  `json:"content" jsonschema:"required,description=Content of the wiki page"`
	Format    *string `json:"format,omitempty" jsonschema:"description=Content format, e.g., markdown"`
}

type UpdateWikiPageArgs struct {
	ProjectID string  `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	Slug      string  `json:"slug" jsonschema:"required,description=URL-encoded slug of the wiki page"`
	Title     *string `json:"title,omitempty" jsonschema:"description=New title of the wiki page"`
	Content   *string `json:"content,omitempty" jsonschema:"description=New content of the wiki page"`
	Format    *string `json:"format,omitempty" jsonschema:"description=Content format, e.g., markdown"`
}

type DeleteWikiPageArgs struct {
	ProjectID string `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	Slug      string `json:"slug" jsonschema:"required,description=URL-encoded slug of the wiki page"`
}

// Pipeline tool args
type ListPipelinesArgs struct {
	ProjectID    string  `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	Scope        *string `json:"scope,omitempty" jsonschema:"description=The scope of pipelines (running, pending, finished, branches, tags)"`
	Status       *string `json:"status,omitempty" jsonschema:"description=The status of pipelines"`
	Ref          *string `json:"ref,omitempty" jsonschema:"description=The ref of pipelines"`
	SHA          *string `json:"sha,omitempty" jsonschema:"description=The SHA of pipelines"`
	YamlErrors   *bool   `json:"yaml_errors,omitempty" jsonschema:"description=Returns pipelines with invalid configurations"`
	Username     *string `json:"username,omitempty" jsonschema:"description=The username of the user who triggered pipelines"`
	UpdatedAfter *string `json:"updated_after,omitempty" jsonschema:"description=Return pipelines updated after the specified date"`
	OrderBy      *string `json:"order_by,omitempty" jsonschema:"description=Order pipelines by (id, status, ref, updated_at)"`
	Sort         *string `json:"sort,omitempty" jsonschema:"description=Sort pipelines (asc, desc)"`
	Page         *int    `json:"page,omitempty" jsonschema:"description=Page number"`
	PerPage      *int    `json:"per_page,omitempty" jsonschema:"description=Number of items per page"`
}

type GetPipelineArgs struct {
	ProjectID  string `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	PipelineID int    `json:"pipeline_id" jsonschema:"required,description=The ID of the pipeline"`
}

type ListPipelineJobsArgs struct {
	ProjectID      string `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	PipelineID     int    `json:"pipeline_id" jsonschema:"required,description=The ID of the pipeline"`
	Scope          *string `json:"scope,omitempty" jsonschema:"description=The scope of jobs to show"`
	IncludeRetried *bool   `json:"include_retried,omitempty" jsonschema:"description=Whether to include retried jobs"`
	Page           *int    `json:"page,omitempty" jsonschema:"description=Page number"`
	PerPage        *int    `json:"per_page,omitempty" jsonschema:"description=Number of items per page"`
}

type GetPipelineJobArgs struct {
	ProjectID string `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	JobID     int    `json:"job_id" jsonschema:"required,description=The ID of the job"`
}

type CreatePipelineArgs struct {
	ProjectID string              `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	Ref       string              `json:"ref" jsonschema:"required,description=The branch or tag to run the pipeline on"`
	Variables []map[string]string `json:"variables,omitempty" jsonschema:"description=An array of variables to use for the pipeline"`
}

type RetryPipelineArgs struct {
	ProjectID  string `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	PipelineID int    `json:"pipeline_id" jsonschema:"required,description=The ID of the pipeline to retry"`
}

type CancelPipelineArgs struct {
	ProjectID  string `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	PipelineID int    `json:"pipeline_id" jsonschema:"required,description=The ID of the pipeline to cancel"`
}

// Milestone tool args
type ListMilestonesArgs struct {
	ProjectID        string   `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	IIDs             []int    `json:"iids,omitempty" jsonschema:"description=Return only milestones having the given iid"`
	State            *string  `json:"state,omitempty" jsonschema:"description=Return only active or closed milestones"`
	Title            *string  `json:"title,omitempty" jsonschema:"description=Return only milestones with a matching title"`
	Search           *string  `json:"search,omitempty" jsonschema:"description=Search milestones by title or description"`
	IncludeAncestors *bool    `json:"include_ancestors,omitempty" jsonschema:"description=Include ancestor groups"`
	UpdatedBefore    *string  `json:"updated_before,omitempty" jsonschema:"description=Return milestones updated before (ISO 8601)"`
	UpdatedAfter     *string  `json:"updated_after,omitempty" jsonschema:"description=Return milestones updated after (ISO 8601)"`
	Page             *int     `json:"page,omitempty" jsonschema:"description=Page number"`
	PerPage          *int     `json:"per_page,omitempty" jsonschema:"description=Number of items per page"`
}

type GetMilestoneArgs struct {
	ProjectID   string `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	MilestoneID int    `json:"milestone_id" jsonschema:"required,description=The ID of a project milestone"`
}

type CreateMilestoneArgs struct {
	ProjectID   string  `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	Title       string  `json:"title" jsonschema:"required,description=The title of the milestone"`
	Description *string `json:"description,omitempty" jsonschema:"description=The description of the milestone"`
	DueDate     *string `json:"due_date,omitempty" jsonschema:"description=The due date of the milestone (YYYY-MM-DD)"`
	StartDate   *string `json:"start_date,omitempty" jsonschema:"description=The start date of the milestone (YYYY-MM-DD)"`
}

type EditMilestoneArgs struct {
	ProjectID   string  `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	MilestoneID int     `json:"milestone_id" jsonschema:"required,description=The ID of a project milestone"`
	Title       *string `json:"title,omitempty" jsonschema:"description=The title of the milestone"`
	Description *string `json:"description,omitempty" jsonschema:"description=The description of the milestone"`
	DueDate     *string `json:"due_date,omitempty" jsonschema:"description=The due date of the milestone (YYYY-MM-DD)"`
	StartDate   *string `json:"start_date,omitempty" jsonschema:"description=The start date of the milestone (YYYY-MM-DD)"`
	StateEvent  *string `json:"state_event,omitempty" jsonschema:"description=The state event of the milestone (close, activate)"`
}

type DeleteMilestoneArgs struct {
	ProjectID   string `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	MilestoneID int    `json:"milestone_id" jsonschema:"required,description=The ID of a project milestone"`
}

type GetMilestoneIssuesArgs struct {
	ProjectID   string `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	MilestoneID int    `json:"milestone_id" jsonschema:"required,description=The ID of the milestone"`
}

type GetMilestoneMergeRequestsArgs struct {
	ProjectID   string `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	MilestoneID int    `json:"milestone_id" jsonschema:"required,description=The ID of the milestone"`
}

type PromoteMilestoneArgs struct {
	ProjectID   string `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	MilestoneID int    `json:"milestone_id" jsonschema:"required,description=The ID of the milestone"`
}

type GetMilestoneBurndownArgs struct {
	ProjectID   string `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	MilestoneID int    `json:"milestone_id" jsonschema:"required,description=The ID of the milestone"`
}

func registerWikiTools(reg toolReg, client *gitlabClient) {
	reg("list_wiki_pages", "List wiki pages in a GitLab project",
		func(ctx context.Context, args ListWikiPagesArgs) (*mcp_golang.ToolResponse, error) {
			withContent := false
			if args.WithContent != nil {
				withContent = *args.WithContent
			}
			page := 0
			if args.Page != nil {
				page = *args.Page
			}
			perPage := 0
			if args.PerPage != nil {
				perPage = *args.PerPage
			}
			result, err := client.ListWikiPages(args.ProjectID, withContent, page, perPage)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("get_wiki_page", "Get details of a specific wiki page",
		func(ctx context.Context, args GetWikiPageArgs) (*mcp_golang.ToolResponse, error) {
			result, err := client.GetWikiPage(args.ProjectID, args.Slug)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("create_wiki_page", "Create a new wiki page in a GitLab project",
		func(ctx context.Context, args CreateWikiPageArgs) (*mcp_golang.ToolResponse, error) {
			format := ""
			if args.Format != nil {
				format = *args.Format
			}
			result, err := client.CreateWikiPage(args.ProjectID, args.Title, args.Content, format)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("update_wiki_page", "Update an existing wiki page in a GitLab project",
		func(ctx context.Context, args UpdateWikiPageArgs) (*mcp_golang.ToolResponse, error) {
			title := ""
			if args.Title != nil {
				title = *args.Title
			}
			content := ""
			if args.Content != nil {
				content = *args.Content
			}
			format := ""
			if args.Format != nil {
				format = *args.Format
			}
			result, err := client.UpdateWikiPage(args.ProjectID, args.Slug, title, content, format)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("delete_wiki_page", "Delete a wiki page from a GitLab project",
		func(ctx context.Context, args DeleteWikiPageArgs) (*mcp_golang.ToolResponse, error) {
			if err := client.DeleteWikiPage(args.ProjectID, args.Slug); err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(
				toJSON(map[string]string{"status": "success", "message": "Wiki page deleted successfully"}))), nil
		})
}

func registerPipelineTools(reg toolReg, client *gitlabClient) {
	reg("list_pipelines", "List pipelines in a GitLab project with filtering options",
		func(ctx context.Context, args ListPipelinesArgs) (*mcp_golang.ToolResponse, error) {
			params := toURLValues(args)
			result, err := client.ListPipelines(args.ProjectID, params)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("get_pipeline", "Get details of a specific pipeline in a GitLab project",
		func(ctx context.Context, args GetPipelineArgs) (*mcp_golang.ToolResponse, error) {
			result, err := client.GetPipeline(args.ProjectID, args.PipelineID)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("list_pipeline_jobs", "List all jobs in a specific pipeline",
		func(ctx context.Context, args ListPipelineJobsArgs) (*mcp_golang.ToolResponse, error) {
			params := toURLValues(args)
			result, err := client.ListPipelineJobs(args.ProjectID, args.PipelineID, params)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("get_pipeline_job", "Get details of a GitLab pipeline job",
		func(ctx context.Context, args GetPipelineJobArgs) (*mcp_golang.ToolResponse, error) {
			result, err := client.GetPipelineJob(args.ProjectID, args.JobID)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("get_pipeline_job_output", "Get the output/trace of a GitLab pipeline job",
		func(ctx context.Context, args GetPipelineJobArgs) (*mcp_golang.ToolResponse, error) {
			result, err := client.GetPipelineJobOutput(args.ProjectID, args.JobID)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(result)), nil
		})

	reg("create_pipeline", "Create a new pipeline for a branch or tag",
		func(ctx context.Context, args CreatePipelineArgs) (*mcp_golang.ToolResponse, error) {
			vars := args.Variables
			if vars == nil {
				vars = []map[string]string{}
			}
			result, err := client.CreatePipeline(args.ProjectID, args.Ref, vars)
			if err != nil {
				return nil, err
			}
			msg := "Created pipeline #" + strconv.Itoa(result.ID) + " for " + args.Ref + ". Status: " + result.Status
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(msg)), nil
		})

	reg("retry_pipeline", "Retry a failed or canceled pipeline",
		func(ctx context.Context, args RetryPipelineArgs) (*mcp_golang.ToolResponse, error) {
			result, err := client.RetryPipeline(args.ProjectID, args.PipelineID)
			if err != nil {
				return nil, err
			}
			msg := "Retried pipeline #" + strconv.Itoa(result.ID) + ". Status: " + result.Status
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(msg)), nil
		})

	reg("cancel_pipeline", "Cancel a running pipeline",
		func(ctx context.Context, args CancelPipelineArgs) (*mcp_golang.ToolResponse, error) {
			result, err := client.CancelPipeline(args.ProjectID, args.PipelineID)
			if err != nil {
				return nil, err
			}
			msg := "Canceled pipeline #" + strconv.Itoa(result.ID) + ". Status: " + result.Status
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(msg)), nil
		})
}

func registerMilestoneTools(reg toolReg, client *gitlabClient) {
	reg("list_milestones", "List milestones in a GitLab project with filtering options",
		func(ctx context.Context, args ListMilestonesArgs) (*mcp_golang.ToolResponse, error) {
			params := toURLValues(args)
			result, err := client.ListProjectMilestones(args.ProjectID, params)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("get_milestone", "Get details of a specific milestone",
		func(ctx context.Context, args GetMilestoneArgs) (*mcp_golang.ToolResponse, error) {
			result, err := client.GetProjectMilestone(args.ProjectID, args.MilestoneID)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("create_milestone", "Create a new milestone in a GitLab project",
		func(ctx context.Context, args CreateMilestoneArgs) (*mcp_golang.ToolResponse, error) {
			desc := ""
			if args.Description != nil {
				desc = *args.Description
			}
			dueDate := ""
			if args.DueDate != nil {
				dueDate = *args.DueDate
			}
			startDate := ""
			if args.StartDate != nil {
				startDate = *args.StartDate
			}
			result, err := client.CreateProjectMilestone(args.ProjectID, args.Title, desc, dueDate, startDate)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("edit_milestone", "Edit an existing milestone in a GitLab project",
		func(ctx context.Context, args EditMilestoneArgs) (*mcp_golang.ToolResponse, error) {
			title := ""
			if args.Title != nil {
				title = *args.Title
			}
			desc := ""
			if args.Description != nil {
				desc = *args.Description
			}
			dueDate := ""
			if args.DueDate != nil {
				dueDate = *args.DueDate
			}
			startDate := ""
			if args.StartDate != nil {
				startDate = *args.StartDate
			}
			stateEvent := ""
			if args.StateEvent != nil {
				stateEvent = *args.StateEvent
			}
			result, err := client.EditProjectMilestone(args.ProjectID, args.MilestoneID, title, desc, dueDate, startDate, stateEvent)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("delete_milestone", "Delete a milestone from a GitLab project",
		func(ctx context.Context, args DeleteMilestoneArgs) (*mcp_golang.ToolResponse, error) {
			if err := client.DeleteProjectMilestone(args.ProjectID, args.MilestoneID); err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(
				toJSON(map[string]string{"status": "success", "message": "Milestone deleted successfully"}))), nil
		})

	reg("get_milestone_issue", "Get issues associated with a specific milestone",
		func(ctx context.Context, args GetMilestoneIssuesArgs) (*mcp_golang.ToolResponse, error) {
			result, err := client.GetMilestoneIssues(args.ProjectID, args.MilestoneID)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("get_milestone_merge_requests", "Get merge requests associated with a specific milestone",
		func(ctx context.Context, args GetMilestoneMergeRequestsArgs) (*mcp_golang.ToolResponse, error) {
			result, err := client.GetMilestoneMergeRequests(args.ProjectID, args.MilestoneID)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("promote_milestone", "Promote a milestone to the next stage",
		func(ctx context.Context, args PromoteMilestoneArgs) (*mcp_golang.ToolResponse, error) {
			result, err := client.PromoteProjectMilestone(args.ProjectID, args.MilestoneID)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("get_milestone_burndown_events", "Get burndown events for a specific milestone",
		func(ctx context.Context, args GetMilestoneBurndownArgs) (*mcp_golang.ToolResponse, error) {
			result, err := client.GetMilestoneBurndownEvents(args.ProjectID, args.MilestoneID)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})
}
