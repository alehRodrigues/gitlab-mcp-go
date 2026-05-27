package tools

import (
	"context"

	mcp_golang "github.com/metoro-io/mcp-golang"
)

type ListCommitsArgs struct {
	ProjectID   string   `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	Ref         *string  `json:"ref_name,omitempty" jsonschema:"description=The branch or tag name"`
	SHA         *string  `json:"sha,omitempty" jsonschema:"description=List commits up to this SHA"`
	Path        *string  `json:"path,omitempty" jsonschema:"description=File path to filter commits"`
	Author      *string  `json:"author,omitempty" jsonschema:"description=Filter by author email or username"`
	Since       *string  `json:"since,omitempty" jsonschema:"description=Commits after this date (ISO 8601)"`
	Until       *string  `json:"until,omitempty" jsonschema:"description=Commits before this date (ISO 8601)"`
	All         *bool    `json:"all,omitempty" jsonschema:"description=Show all commits from all branches"`
	WithStats   *bool    `json:"with_stats,omitempty" jsonschema:"description=Include commit statistics"`
	FirstParent *bool    `json:"first_parent,omitempty" jsonschema:"description=Follow only the first parent of merge commits"`
	Order       *string  `json:"order,omitempty" jsonschema:"description=Order commits by (default, author_date)"`
	Trailers    *bool    `json:"trailers,omitempty" jsonschema:"description=Include commit trailers"`
	Page        *int     `json:"page,omitempty" jsonschema:"description=Page number"`
	PerPage     *int     `json:"per_page,omitempty" jsonschema:"description=Number of items per page"`
}

type GetCommitArgs struct {
	ProjectID string  `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	SHA       string  `json:"sha" jsonschema:"required,description=The commit SHA"`
	Stats     *bool   `json:"stats,omitempty" jsonschema:"description=Include commit statistics"`
}

type GetCommitDiffArgs struct {
	ProjectID string  `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	SHA       string  `json:"sha" jsonschema:"required,description=The commit SHA"`
	Unidiff   *bool   `json:"unidiff,omitempty" jsonschema:"description=Present diffs in unified diff format"`
}

type GetCommitRefsArgs struct {
	ProjectID string  `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	SHA       string  `json:"sha" jsonschema:"required,description=The commit SHA"`
	Type      *string `json:"type,omitempty" jsonschema:"description=Filter refs by type (branch, tag, all)"`
}

type GetCommitCommentsArgs struct {
	ProjectID string  `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	SHA       string  `json:"sha" jsonschema:"required,description=The commit SHA"`
}

type CreateCommitCommentArgs struct {
	ProjectID string  `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	SHA       string  `json:"sha" jsonschema:"required,description=The commit SHA"`
	Note      string  `json:"note" jsonschema:"required,description=The comment text"`
	Path      *string `json:"path,omitempty" jsonschema:"description=File path relative to the repository"`
	Line      *int    `json:"line,omitempty" jsonschema:"description=The line number where the comment should be placed"`
	LineType  *string `json:"line_type,omitempty" jsonschema:"description=The line type (new, old)"`
}

type GetCommitMergeRequestsArgs struct {
	ProjectID string  `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	SHA       string  `json:"sha" jsonschema:"required,description=The commit SHA"`
	State     *string `json:"state,omitempty" jsonschema:"description=Filter by state (opened, closed, locked, merged)"`
}

func registerCommitTools(reg toolReg, client *gitlabClient) {
	reg("list_commits", "List commits in a GitLab project with filtering options",
		func(ctx context.Context, args ListCommitsArgs) (*mcp_golang.ToolResponse, error) {
			params := toURLValues(args)
			result, err := client.ListCommits(args.ProjectID, params)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("get_commit", "Get details of a specific commit",
		func(ctx context.Context, args GetCommitArgs) (*mcp_golang.ToolResponse, error) {
			params := toURLValues(args)
			result, err := client.GetCommit(args.ProjectID, args.SHA, params)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("get_commit_diff", "Get the diff of a specific commit",
		func(ctx context.Context, args GetCommitDiffArgs) (*mcp_golang.ToolResponse, error) {
			params := toURLValues(args)
			result, err := client.GetCommitDiff(args.ProjectID, args.SHA, params)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("get_commit_refs", "Get branches and tags that contain a specific commit",
		func(ctx context.Context, args GetCommitRefsArgs) (*mcp_golang.ToolResponse, error) {
			refType := ""
			if args.Type != nil {
				refType = *args.Type
			}
			result, err := client.GetCommitRefs(args.ProjectID, args.SHA, refType)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("get_commit_comments", "Get comments on a specific commit",
		func(ctx context.Context, args GetCommitCommentsArgs) (*mcp_golang.ToolResponse, error) {
			result, err := client.GetCommitComments(args.ProjectID, args.SHA)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("create_commit_comment", "Add a comment to a specific commit",
		func(ctx context.Context, args CreateCommitCommentArgs) (*mcp_golang.ToolResponse, error) {
			path := ""
			if args.Path != nil {
				path = *args.Path
			}
			lineType := ""
			if args.LineType != nil {
				lineType = *args.LineType
			}
			result, err := client.CreateCommitComment(args.ProjectID, args.SHA, args.Note, path, args.Line, lineType)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("get_commit_merge_requests", "Get merge requests associated with a specific commit",
		func(ctx context.Context, args GetCommitMergeRequestsArgs) (*mcp_golang.ToolResponse, error) {
			params := toURLValues(args)
			result, err := client.GetCommitMergeRequests(args.ProjectID, args.SHA, params)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})
}
