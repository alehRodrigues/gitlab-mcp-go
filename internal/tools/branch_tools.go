package tools

import (
	"context"

	mcp_golang "github.com/metoro-io/mcp-golang"
)

type CreateBranchArgs struct {
	ProjectID string  `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	Branch    string  `json:"branch" jsonschema:"required,description=Name for the new branch"`
	Ref       *string `json:"ref,omitempty" jsonschema:"description=Source branch/commit for new branch"`
}

type GetBranchDiffsArgs struct {
	ProjectID          string   `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	From               string   `json:"from" jsonschema:"required,description=The base branch or commit SHA to compare from"`
	To                 string   `json:"to" jsonschema:"required,description=The target branch or commit SHA to compare to"`
	Straight           *bool    `json:"straight,omitempty" jsonschema:"description=Comparison method: false for '...' (default), true for '--'"`
	ExcludedPatterns   []string `json:"excluded_file_patterns,omitempty" jsonschema:"description=Array of regex patterns to exclude files from diff results"`
}

func registerBranchTools(reg toolReg, client *gitlabClient) {
	reg("create_branch", "Create a new branch in a GitLab project",
		func(ctx context.Context, args CreateBranchArgs) (*mcp_golang.ToolResponse, error) {
			ref := ""
			if args.Ref != nil {
				ref = *args.Ref
			}
			result, err := client.CreateBranch(args.ProjectID, args.Branch, ref)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("get_branch_diffs", "Get the changes/diffs between two branches or commits in a GitLab project",
		func(ctx context.Context, args GetBranchDiffsArgs) (*mcp_golang.ToolResponse, error) {
			straight := false
			if args.Straight != nil {
				straight = *args.Straight
			}
			excluded := args.ExcludedPatterns
			if excluded == nil {
				excluded = []string{}
			}
			result, err := client.GetBranchDiffs(args.ProjectID, args.From, args.To, straight, excluded)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})
}
