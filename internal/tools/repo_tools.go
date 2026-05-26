package tools

import (
	"context"

	mcp_golang "github.com/metoro-io/mcp-golang"

	"github.com/user/gitlab-mcp-go/internal/gitlab"
)

type CreateOrUpdateFileArgs struct {
	ProjectID     string  `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	FilePath      string  `json:"file_path" jsonschema:"required,description=Path where to create/update the file"`
	Content       string  `json:"content" jsonschema:"required,description=Content of the file"`
	CommitMessage string  `json:"commit_message" jsonschema:"required,description=Commit message"`
	Branch        string  `json:"branch" jsonschema:"required,description=Branch to create/update the file in"`
	PreviousPath  *string `json:"previous_path,omitempty" jsonschema:"description=Path of the file to move/rename"`
	LastCommitID  *string `json:"last_commit_id,omitempty" jsonschema:"description=Last known file commit ID"`
	CommitID      *string `json:"commit_id,omitempty" jsonschema:"description=Current file commit ID (for update operations)"`
}

type SearchRepositoriesArgs struct {
	Search  string `json:"search" jsonschema:"required,description=Search query"`
	Page    *int   `json:"page,omitempty" jsonschema:"description=Page number"`
	PerPage *int   `json:"per_page,omitempty" jsonschema:"description=Number of items per page"`
}

type CreateRepositoryArgs struct {
	Name              string  `json:"name" jsonschema:"required,description=Repository name"`
	Description       *string `json:"description,omitempty" jsonschema:"description=Repository description"`
	Visibility        *string `json:"visibility,omitempty" jsonschema:"description=Repository visibility level (private, internal, public)"`
	InitWithReadme    *bool   `json:"initialize_with_readme,omitempty" jsonschema:"description=Initialize with README.md"`
}

type GetFileContentsArgs struct {
	ProjectID string  `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	FilePath  string  `json:"file_path" jsonschema:"required,description=Path to the file or directory"`
	Ref       *string `json:"ref,omitempty" jsonschema:"description=Branch/tag/commit to get contents from"`
}

type PushFilesArgs struct {
	ProjectID     string     `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	Branch        string     `json:"branch" jsonschema:"required,description=Branch to push to"`
	Files         []FileArg  `json:"files" jsonschema:"required,description=Array of files to push"`
	CommitMessage string     `json:"commit_message" jsonschema:"required,description=Commit message"`
}

type FileArg struct {
	FilePath string `json:"file_path" jsonschema:"required,description=Path where to create the file"`
	Content  string `json:"content" jsonschema:"required,description=Content of the file"`
}

type ForkRepositoryArgs struct {
	ProjectID string  `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	Namespace *string `json:"namespace,omitempty" jsonschema:"description=Namespace to fork to (full path)"`
}

type GetRepositoryTreeArgs struct {
	ProjectID  string  `json:"project_id" jsonschema:"required,description=The ID or URL-encoded path of the project"`
	Path       *string `json:"path,omitempty" jsonschema:"description=The path inside the repository"`
	Ref        *string `json:"ref,omitempty" jsonschema:"description=The name of a repository branch or tag"`
	Recursive  *bool   `json:"recursive,omitempty" jsonschema:"description=Boolean value to get a recursive tree"`
	PerPage    *int    `json:"per_page,omitempty" jsonschema:"description=Number of results to show per page"`
	PageToken  *string `json:"page_token,omitempty" jsonschema:"description=The tree record ID for pagination"`
	Pagination *string `json:"pagination,omitempty" jsonschema:"description=Pagination method (keyset)"`
}

func registerRepoAndFileTools(reg toolReg, client *gitlabClient) {
	reg("create_or_update_file", "Create or update a single file in a GitLab project",
		func(ctx context.Context, args CreateOrUpdateFileArgs) (*mcp_golang.ToolResponse, error) {
			result, err := client.CreateOrUpdateFile(args.ProjectID, args.FilePath, args.Content,
				args.CommitMessage, args.Branch,
				ptrStr(args.PreviousPath), ptrStr(args.LastCommitID), ptrStr(args.CommitID))
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("search_repositories", "Search for GitLab projects",
		func(ctx context.Context, args SearchRepositoriesArgs) (*mcp_golang.ToolResponse, error) {
			page := 1
			perPage := 20
			if args.Page != nil {
				page = *args.Page
			}
			if args.PerPage != nil {
				perPage = *args.PerPage
			}
			result, err := client.SearchProjects(args.Search, page, perPage)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("create_repository", "Create a new GitLab project",
		func(ctx context.Context, args CreateRepositoryArgs) (*mcp_golang.ToolResponse, error) {
			vis := ""
			if args.Visibility != nil {
				vis = *args.Visibility
			}
			initReadme := false
			if args.InitWithReadme != nil {
				initReadme = *args.InitWithReadme
			}
			desc := ""
			if args.Description != nil {
				desc = *args.Description
			}
			result, err := client.CreateRepository(args.Name, desc, vis, initReadme)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("get_file_contents", "Get the contents of a file or directory from a GitLab project",
		func(ctx context.Context, args GetFileContentsArgs) (*mcp_golang.ToolResponse, error) {
			ref := ""
			if args.Ref != nil {
				ref = *args.Ref
			}
			result, err := client.GetFileContents(args.ProjectID, args.FilePath, ref)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("push_files", "Push multiple files to a GitLab project in a single commit",
		func(ctx context.Context, args PushFilesArgs) (*mcp_golang.ToolResponse, error) {
			ops := make([]gitlab.FileOperation, len(args.Files))
			for i, f := range args.Files {
				ops[i] = gitlab.FileOperation{FilePath: f.FilePath, Content: f.Content}
			}
			result, err := client.PushFiles(args.ProjectID, args.Branch, args.CommitMessage, ops)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("fork_repository", "Fork a GitLab project to your account or specified namespace",
		func(ctx context.Context, args ForkRepositoryArgs) (*mcp_golang.ToolResponse, error) {
			ns := ""
			if args.Namespace != nil {
				ns = *args.Namespace
			}
			result, err := client.ForkProject(args.ProjectID, ns)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("get_repository_tree", "Get the repository tree for a GitLab project (list files and directories)",
		func(ctx context.Context, args GetRepositoryTreeArgs) (*mcp_golang.ToolResponse, error) {
			path := ""
			if args.Path != nil {
				path = *args.Path
			}
			ref := ""
			if args.Ref != nil {
				ref = *args.Ref
			}
			recursive := false
			if args.Recursive != nil {
				recursive = *args.Recursive
			}
			perPage := 0
			if args.PerPage != nil {
				perPage = *args.PerPage
			}
			pageToken := ""
			if args.PageToken != nil {
				pageToken = *args.PageToken
			}
			pagination := ""
			if args.Pagination != nil {
				pagination = *args.Pagination
			}
			result, err := client.GetRepositoryTree(args.ProjectID, path, ref, recursive, perPage, pageToken, pagination)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})
}

func ptrStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
