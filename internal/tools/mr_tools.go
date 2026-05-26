package tools

import (
	"context"

	mcp_golang "github.com/metoro-io/mcp-golang"
)

type CreateMergeRequestArgs struct {
	ProjectID          string   `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	Title              string   `json:"title" jsonschema:"required,description=Merge request title"`
	Description        *string  `json:"description,omitempty" jsonschema:"description=Merge request description"`
	SourceBranch       string   `json:"source_branch" jsonschema:"required,description=Branch containing changes"`
	TargetBranch       string   `json:"target_branch" jsonschema:"required,description=Branch to merge into"`
	AssigneeIDs        []int    `json:"assignee_ids,omitempty" jsonschema:"description=Array of user IDs to assign the MR to"`
	ReviewerIDs        []int    `json:"reviewer_ids,omitempty" jsonschema:"description=Array of user IDs to assign as reviewers"`
	Labels             []string `json:"labels,omitempty" jsonschema:"description=Labels for the MR"`
	Draft              *bool    `json:"draft,omitempty" jsonschema:"description=Create as draft merge request"`
	AllowCollaboration *bool    `json:"allow_collaboration,omitempty" jsonschema:"description=Allow commits from upstream members"`
}

type GetMergeRequestArgs struct {
	ProjectID       string  `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	MergeRequestIID *int    `json:"merge_request_iid,omitempty" jsonschema:"description=The IID of a merge request"`
	SourceBranch    *string `json:"source_branch,omitempty" jsonschema:"description=Source branch name"`
}

type GetMergeRequestDiffsArgs struct {
	ProjectID       string  `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	MergeRequestIID *int    `json:"merge_request_iid,omitempty" jsonschema:"description=The IID of a merge request"`
	SourceBranch    *string `json:"source_branch,omitempty" jsonschema:"description=Source branch name"`
	View            *string `json:"view,omitempty" jsonschema:"description=Diff view type (inline, parallel)"`
}

type UpdateMergeRequestArgs struct {
	ProjectID        string   `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	MergeRequestIID  *int     `json:"merge_request_iid,omitempty" jsonschema:"description=The IID of a merge request"`
	SourceBranch     *string  `json:"source_branch,omitempty" jsonschema:"description=Source branch name"`
	Title            *string  `json:"title,omitempty" jsonschema:"description=The title of the merge request"`
	Description      *string  `json:"description,omitempty" jsonschema:"description=The description of the merge request"`
	TargetBranch     *string  `json:"target_branch,omitempty" jsonschema:"description=The target branch"`
	AssigneeIDs      []int    `json:"assignee_ids,omitempty" jsonschema:"description=Array of user IDs to assign"`
	Labels           []string `json:"labels,omitempty" jsonschema:"description=Labels for the MR"`
	StateEvent       *string  `json:"state_event,omitempty" jsonschema:"description=New state (close/reopen) for the MR"`
	RemoveSourceBranch *bool  `json:"remove_source_branch,omitempty" jsonschema:"description=Flag indicating if the source branch should be removed"`
	Squash           *bool    `json:"squash,omitempty" jsonschema:"description=Squash commits into a single commit when merging"`
	Draft            *bool    `json:"draft,omitempty" jsonschema:"description=Draft status"`
}

type ListMergeRequestsArgs struct {
	ProjectID        string   `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	AssigneeID       *int     `json:"assignee_id,omitempty" jsonschema:"description=Returns MRs assigned to the given user ID"`
	AssigneeUsername *string  `json:"assignee_username,omitempty" jsonschema:"description=Returns MRs assigned to the given username"`
	AuthorID         *int     `json:"author_id,omitempty" jsonschema:"description=Returns MRs created by the given user ID"`
	AuthorUsername   *string  `json:"author_username,omitempty" jsonschema:"description=Returns MRs created by the given username"`
	ReviewerID       *int     `json:"reviewer_id,omitempty" jsonschema:"description=Returns MRs which have the user as a reviewer"`
	ReviewerUsername *string  `json:"reviewer_username,omitempty" jsonschema:"description=Returns MRs which have the user as a reviewer"`
	CreatedAfter     *string  `json:"created_after,omitempty" jsonschema:"description=Return MRs created after the given time (ISO 8601)"`
	CreatedBefore    *string  `json:"created_before,omitempty" jsonschema:"description=Return MRs created before the given time (ISO 8601)"`
	UpdatedAfter     *string  `json:"updated_after,omitempty" jsonschema:"description=Return MRs updated after the given time"`
	UpdatedBefore    *string  `json:"updated_before,omitempty" jsonschema:"description=Return MRs updated before the given time"`
	Labels           []string `json:"labels,omitempty" jsonschema:"description=Array of label names"`
	Milestone        *string  `json:"milestone,omitempty" jsonschema:"description=Milestone title"`
	Scope            *string  `json:"scope,omitempty" jsonschema:"description=Scope (created_by_me, assigned_to_me, all)"`
	Search           *string  `json:"search,omitempty" jsonschema:"description=Search for specific terms"`
	State            *string  `json:"state,omitempty" jsonschema:"description=State (opened, closed, locked, merged, all)"`
	OrderBy          *string  `json:"order_by,omitempty" jsonschema:"description=Order by (created_at, updated_at, priority)"`
	Sort             *string  `json:"sort,omitempty" jsonschema:"description=Sort direction (asc, desc)"`
	TargetBranch     *string  `json:"target_branch,omitempty" jsonschema:"description=Filter by target branch"`
	SourceBranch     *string  `json:"source_branch,omitempty" jsonschema:"description=Filter by source branch"`
	Page             *int     `json:"page,omitempty" jsonschema:"description=Page number"`
	PerPage          *int     `json:"per_page,omitempty" jsonschema:"description=Number of items per page"`
}

func registerMRTools(reg toolReg, client *gitlabClient) {
	reg("create_merge_request", "Create a new merge request in a GitLab project",
		func(ctx context.Context, args CreateMergeRequestArgs) (*mcp_golang.ToolResponse, error) {
			desc := ""
			if args.Description != nil {
				desc = *args.Description
			}
			draft := false
			if args.Draft != nil {
				draft = *args.Draft
			}
			allowCollab := false
			if args.AllowCollaboration != nil {
				allowCollab = *args.AllowCollaboration
			}
			if args.AssigneeIDs == nil {
				args.AssigneeIDs = []int{}
			}
			if args.ReviewerIDs == nil {
				args.ReviewerIDs = []int{}
			}
			if args.Labels == nil {
				args.Labels = []string{}
			}
			result, err := client.CreateMergeRequest(args.ProjectID, args.Title, desc,
				args.SourceBranch, args.TargetBranch, args.AssigneeIDs, args.ReviewerIDs,
				args.Labels, allowCollab, draft)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("get_merge_request", "Get details of a merge request (Either mergeRequestIid or sourceBranch must be provided)",
		func(ctx context.Context, args GetMergeRequestArgs) (*mcp_golang.ToolResponse, error) {
			mrIID := 0
			if args.MergeRequestIID != nil {
				mrIID = *args.MergeRequestIID
			}
			branch := ""
			if args.SourceBranch != nil {
				branch = *args.SourceBranch
			}
			result, err := client.GetMergeRequest(args.ProjectID, mrIID, branch)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("get_merge_request_diffs", "Get the changes/diffs of a merge request (Either mergeRequestIid or sourceBranch must be provided)",
		func(ctx context.Context, args GetMergeRequestDiffsArgs) (*mcp_golang.ToolResponse, error) {
			mrIID := 0
			if args.MergeRequestIID != nil {
				mrIID = *args.MergeRequestIID
			}
			branch := ""
			if args.SourceBranch != nil {
				branch = *args.SourceBranch
			}
			view := ""
			if args.View != nil {
				view = *args.View
			}
			result, err := client.GetMergeRequestDiffs(args.ProjectID, mrIID, branch, view)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("update_merge_request", "Update a merge request (Either mergeRequestIid or sourceBranch must be provided)",
		func(ctx context.Context, args UpdateMergeRequestArgs) (*mcp_golang.ToolResponse, error) {
			mrIID := 0
			if args.MergeRequestIID != nil {
				mrIID = *args.MergeRequestIID
			}
			branch := ""
			if args.SourceBranch != nil {
				branch = *args.SourceBranch
			}
			body := map[string]any{}
			if args.Title != nil {
				body["title"] = *args.Title
			}
			if args.Description != nil {
				body["description"] = *args.Description
			}
			if args.TargetBranch != nil {
				body["target_branch"] = *args.TargetBranch
			}
			if len(args.AssigneeIDs) > 0 {
				body["assignee_ids"] = args.AssigneeIDs
			}
			if len(args.Labels) > 0 {
				body["labels"] = args.Labels
			}
			if args.StateEvent != nil {
				body["state_event"] = *args.StateEvent
			}
			if args.RemoveSourceBranch != nil {
				body["remove_source_branch"] = *args.RemoveSourceBranch
			}
			if args.Squash != nil {
				body["squash"] = *args.Squash
			}
			if args.Draft != nil {
				body["draft"] = *args.Draft
			}
			result, err := client.UpdateMergeRequest(args.ProjectID, mrIID, branch, body)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("list_merge_requests", "List merge requests in a GitLab project with filtering options",
		func(ctx context.Context, args ListMergeRequestsArgs) (*mcp_golang.ToolResponse, error) {
			params := toURLValues(args)
			result, err := client.ListMergeRequests(args.ProjectID, params)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})
}
