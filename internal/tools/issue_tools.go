package tools

import (
	"context"

	mcp_golang "github.com/metoro-io/mcp-golang"
)

type CreateIssueArgs struct {
	ProjectID   string   `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	Title       string   `json:"title" jsonschema:"required,description=Issue title"`
	Description *string  `json:"description,omitempty" jsonschema:"description=Issue description"`
	AssigneeIDs []int    `json:"assignee_ids,omitempty" jsonschema:"description=Array of user IDs to assign"`
	Labels      []string `json:"labels,omitempty" jsonschema:"description=Array of label names"`
	MilestoneID *int     `json:"milestone_id,omitempty" jsonschema:"description=Milestone ID to assign"`
}

type ListIssuesArgs struct {
	ProjectID        string   `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	AssigneeID       *int     `json:"assignee_id,omitempty" jsonschema:"description=Return issues assigned to the given user ID"`
	AssigneeUsername []string `json:"assignee_username,omitempty" jsonschema:"description=Return issues assigned to the given username"`
	AuthorID         *int     `json:"author_id,omitempty" jsonschema:"description=Return issues created by the given user ID"`
	AuthorUsername   *string  `json:"author_username,omitempty" jsonschema:"description=Return issues created by the given username"`
	Confidential     *bool    `json:"confidential,omitempty" jsonschema:"description=Filter confidential or public issues"`
	CreatedAfter     *string  `json:"created_after,omitempty" jsonschema:"description=Return issues created after the given time (ISO 8601)"`
	CreatedBefore    *string  `json:"created_before,omitempty" jsonschema:"description=Return issues created before the given time (ISO 8601)"`
	DueDate          *string  `json:"due_date,omitempty" jsonschema:"description=Return issues that have the due date"`
	Labels           []string `json:"labels,omitempty" jsonschema:"description=Array of label names"`
	Milestone        *string  `json:"milestone,omitempty" jsonschema:"description=Milestone title"`
	Scope            *string  `json:"scope,omitempty" jsonschema:"description=Return issues from a specific scope (created_by_me, assigned_to_me, all)"`
	Search           *string  `json:"search,omitempty" jsonschema:"description=Search for specific terms"`
	State            *string  `json:"state,omitempty" jsonschema:"description=Return issues with a specific state (opened, closed, all)"`
	UpdatedAfter     *string  `json:"updated_after,omitempty" jsonschema:"description=Return issues updated after the given time (ISO 8601)"`
	UpdatedBefore    *string  `json:"updated_before,omitempty" jsonschema:"description=Return issues updated before the given time (ISO 8601)"`
	Page             *int     `json:"page,omitempty" jsonschema:"description=Page number"`
	PerPage          *int     `json:"per_page,omitempty" jsonschema:"description=Number of items per page"`
}

type GetIssueArgs struct {
	ProjectID string `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	IssueIID  int    `json:"issue_iid" jsonschema:"required,description=The internal ID of the project issue"`
}

type UpdateIssueArgs struct {
	ProjectID       string   `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	IssueIID        int      `json:"issue_iid" jsonschema:"required,description=The internal ID of the project issue"`
	Title           *string  `json:"title,omitempty" jsonschema:"description=The title of the issue"`
	Description     *string  `json:"description,omitempty" jsonschema:"description=The description of the issue"`
	AssigneeIDs     []int    `json:"assignee_ids,omitempty" jsonschema:"description=Array of user IDs to assign issue to"`
	Confidential    *bool    `json:"confidential,omitempty" jsonschema:"description=Set the issue to be confidential"`
	DiscussionLocked *bool   `json:"discussion_locked,omitempty" jsonschema:"description=Flag to lock discussions"`
	DueDate         *string  `json:"due_date,omitempty" jsonschema:"description=Date the issue is due (YYYY-MM-DD)"`
	Labels          []string `json:"labels,omitempty" jsonschema:"description=Array of label names"`
	MilestoneID     *int     `json:"milestone_id,omitempty" jsonschema:"description=Milestone ID to assign"`
	StateEvent      *string  `json:"state_event,omitempty" jsonschema:"description=Update issue state (close/reopen)"`
	Weight          *int     `json:"weight,omitempty" jsonschema:"description=Weight of the issue (0-9)"`
}

type DeleteIssueArgs struct {
	ProjectID string `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	IssueIID  int    `json:"issue_iid" jsonschema:"required,description=The internal ID of the project issue"`
}

type ListIssueLinksArgs struct {
	ProjectID string `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	IssueIID  int    `json:"issue_iid" jsonschema:"required,description=The internal ID of a project issue"`
}

type ListIssueDiscussionsArgs struct {
	ProjectID string `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	IssueIID  int    `json:"issue_iid" jsonschema:"required,description=The internal ID of the project issue"`
	Page      *int   `json:"page,omitempty" jsonschema:"description=Page number"`
	PerPage   *int   `json:"per_page,omitempty" jsonschema:"description=Number of items per page"`
}

type GetIssueLinkArgs struct {
	ProjectID  string `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	IssueIID   int    `json:"issue_iid" jsonschema:"required,description=The internal ID of a project issue"`
	IssueLinkID int   `json:"issue_link_id" jsonschema:"required,description=ID of an issue relationship"`
}

type CreateIssueLinkArgs struct {
	ProjectID       string  `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	IssueIID        int     `json:"issue_iid" jsonschema:"required,description=The internal ID of a project issue"`
	TargetProjectID string  `json:"target_project_id" jsonschema:"required,description=The ID or URL-encoded path of a target project"`
	TargetIssueIID  int     `json:"target_issue_iid" jsonschema:"required,description=The internal ID of a target project issue"`
	LinkType        *string `json:"link_type,omitempty" jsonschema:"description=The type of the relation (relates_to, blocks, is_blocked_by)"`
}

type DeleteIssueLinkArgs struct {
	ProjectID  string `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	IssueIID   int    `json:"issue_iid" jsonschema:"required,description=The internal ID of a project issue"`
	IssueLinkID int   `json:"issue_link_id" jsonschema:"required,description=The ID of an issue relationship"`
}

func registerIssueTools(reg toolReg, client *gitlabClient) {
	reg("create_issue", "Create a new issue in a GitLab project",
		func(ctx context.Context, args CreateIssueArgs) (*mcp_golang.ToolResponse, error) {
			desc := ""
			if args.Description != nil {
				desc = *args.Description
			}
			milestoneID := 0
			if args.MilestoneID != nil {
				milestoneID = *args.MilestoneID
			}
			assignees := args.AssigneeIDs
			if assignees == nil {
				assignees = []int{}
			}
			labels := args.Labels
			if labels == nil {
				labels = []string{}
			}
			result, err := client.CreateIssue(args.ProjectID, args.Title, desc, assignees, labels, milestoneID)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("list_issues", "List issues in a GitLab project with filtering options",
		func(ctx context.Context, args ListIssuesArgs) (*mcp_golang.ToolResponse, error) {
			params := toURLValues(args)
			issues, err := client.ListIssues(args.ProjectID, params)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(issues))), nil
		})

	reg("get_issue", "Get details of a specific issue in a GitLab project",
		func(ctx context.Context, args GetIssueArgs) (*mcp_golang.ToolResponse, error) {
			result, err := client.GetIssue(args.ProjectID, args.IssueIID)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("update_issue", "Update an issue in a GitLab project",
		func(ctx context.Context, args UpdateIssueArgs) (*mcp_golang.ToolResponse, error) {
			body := map[string]any{}
			if args.Title != nil {
				body["title"] = *args.Title
			}
			if args.Description != nil {
				body["description"] = *args.Description
			}
			if len(args.AssigneeIDs) > 0 {
				body["assignee_ids"] = args.AssigneeIDs
			}
			if args.Confidential != nil {
				body["confidential"] = *args.Confidential
			}
			if args.DiscussionLocked != nil {
				body["discussion_locked"] = *args.DiscussionLocked
			}
			if args.DueDate != nil {
				body["due_date"] = *args.DueDate
			}
			if len(args.Labels) > 0 {
				body["labels"] = args.Labels
			}
			if args.MilestoneID != nil {
				body["milestone_id"] = *args.MilestoneID
			}
			if args.StateEvent != nil {
				body["state_event"] = *args.StateEvent
			}
			if args.Weight != nil {
				body["weight"] = *args.Weight
			}
			result, err := client.UpdateIssue(args.ProjectID, args.IssueIID, body)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("delete_issue", "Delete an issue from a GitLab project",
		func(ctx context.Context, args DeleteIssueArgs) (*mcp_golang.ToolResponse, error) {
			if err := client.DeleteIssue(args.ProjectID, args.IssueIID); err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(
				toJSON(map[string]string{"status": "success", "message": "Issue deleted successfully"}))), nil
		})

	reg("list_issue_links", "List all issue links for a specific issue",
		func(ctx context.Context, args ListIssueLinksArgs) (*mcp_golang.ToolResponse, error) {
			result, err := client.ListIssueLinks(args.ProjectID, args.IssueIID)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("list_issue_discussions", "List discussions for an issue in a GitLab project",
		func(ctx context.Context, args ListIssueDiscussionsArgs) (*mcp_golang.ToolResponse, error) {
			page := 0
			if args.Page != nil {
				page = *args.Page
			}
			perPage := 0
			if args.PerPage != nil {
				perPage = *args.PerPage
			}
			result, err := client.ListIssueDiscussions(args.ProjectID, args.IssueIID, page, perPage)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("get_issue_link", "Get a specific issue link",
		func(ctx context.Context, args GetIssueLinkArgs) (*mcp_golang.ToolResponse, error) {
			result, err := client.GetIssueLink(args.ProjectID, args.IssueIID, args.IssueLinkID)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("create_issue_link", "Create an issue link between two issues",
		func(ctx context.Context, args CreateIssueLinkArgs) (*mcp_golang.ToolResponse, error) {
			linkType := "relates_to"
			if args.LinkType != nil {
				linkType = *args.LinkType
			}
			result, err := client.CreateIssueLink(args.ProjectID, args.IssueIID, args.TargetProjectID, args.TargetIssueIID, linkType)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("delete_issue_link", "Delete an issue link",
		func(ctx context.Context, args DeleteIssueLinkArgs) (*mcp_golang.ToolResponse, error) {
			if err := client.DeleteIssueLink(args.ProjectID, args.IssueIID, args.IssueLinkID); err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(
				toJSON(map[string]string{"status": "success", "message": "Issue link deleted successfully"}))), nil
		})
}
