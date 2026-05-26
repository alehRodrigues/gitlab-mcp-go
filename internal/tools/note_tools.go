package tools

import (
	"context"

	mcp_golang "github.com/metoro-io/mcp-golang"

	"github.com/user/gitlab-mcp-go/internal/gitlab"
)

type CreateNoteArgs struct {
	ProjectID    string `json:"project_id" jsonschema:"required,description=Project ID or namespace/project_path"`
	NoteableType string `json:"noteable_type" jsonschema:"required,description=Type of noteable (issue, merge_request)"`
	NoteableIID  int    `json:"noteable_iid" jsonschema:"required,description=IID of the issue or merge request"`
	Body         string `json:"body" jsonschema:"required,description=Note content"`
}

type CreateMergeRequestThreadArgs struct {
	ProjectID       string              `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	MergeRequestIID int                 `json:"merge_request_iid" jsonschema:"required,description=The IID of a merge request"`
	Body            string              `json:"body" jsonschema:"required,description=The content of the thread"`
	Position        *NotePositionArg    `json:"position,omitempty" jsonschema:"description=Position when creating a diff note"`
	CreatedAt       *string             `json:"created_at,omitempty" jsonschema:"description=Date the thread was created at (ISO 8601 format)"`
}

type NotePositionArg struct {
	BaseSHA      string  `json:"base_sha" jsonschema:"required,description=Base commit SHA in the source branch"`
	HeadSHA      string  `json:"head_sha" jsonschema:"required,description=SHA referencing HEAD of the source branch"`
	StartSHA     string  `json:"start_sha" jsonschema:"required,description=SHA referencing the start commit of the source branch"`
	PositionType string  `json:"position_type" jsonschema:"required,description=Type of position reference (text, image, file)"`
	NewPath      *string `json:"new_path,omitempty" jsonschema:"description=File path after change"`
	OldPath      *string `json:"old_path,omitempty" jsonschema:"description=File path before change"`
	NewLine      *int    `json:"new_line,omitempty" jsonschema:"description=Line number after change"`
	OldLine      *int    `json:"old_line,omitempty" jsonschema:"description=Line number before change"`
}

type ListMRDiscussionsArgs struct {
	ProjectID       string `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	MergeRequestIID int    `json:"merge_request_iid" jsonschema:"required,description=The IID of a merge request"`
	Page            *int   `json:"page,omitempty" jsonschema:"description=Page number"`
	PerPage         *int   `json:"per_page,omitempty" jsonschema:"description=Number of items per page"`
}

type UpdateMergeRequestNoteArgs struct {
	ProjectID       string  `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	MergeRequestIID int     `json:"merge_request_iid" jsonschema:"required,description=The IID of a merge request"`
	DiscussionID    string  `json:"discussion_id" jsonschema:"required,description=The ID of a thread"`
	NoteID          int     `json:"note_id" jsonschema:"required,description=The ID of a thread note"`
	Body            *string `json:"body,omitempty" jsonschema:"description=The content of the note or reply"`
	Resolved        *bool   `json:"resolved,omitempty" jsonschema:"description=Resolve or unresolve the note"`
}

type CreateMergeRequestNoteArgs struct {
	ProjectID       string  `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	MergeRequestIID int     `json:"merge_request_iid" jsonschema:"required,description=The IID of a merge request"`
	DiscussionID    string  `json:"discussion_id" jsonschema:"required,description=The ID of a thread"`
	Body            string  `json:"body" jsonschema:"required,description=The content of the note or reply"`
	CreatedAt       *string `json:"created_at,omitempty" jsonschema:"description=Date the note was created at (ISO 8601 format)"`
}

type UpdateIssueNoteArgs struct {
	ProjectID    string `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	IssueIID     int    `json:"issue_iid" jsonschema:"required,description=The IID of an issue"`
	DiscussionID string `json:"discussion_id" jsonschema:"required,description=The ID of a thread"`
	NoteID       int    `json:"note_id" jsonschema:"required,description=The ID of a thread note"`
	Body         string `json:"body" jsonschema:"required,description=The content of the note or reply"`
}

type CreateIssueNoteArgs struct {
	ProjectID    string  `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	IssueIID     int     `json:"issue_iid" jsonschema:"required,description=The IID of an issue"`
	DiscussionID string  `json:"discussion_id" jsonschema:"required,description=The ID of a thread"`
	Body         string  `json:"body" jsonschema:"required,description=The content of the note or reply"`
	CreatedAt    *string `json:"created_at,omitempty" jsonschema:"description=Date the note was created at (ISO 8601 format)"`
}

func registerNoteTools(reg toolReg, client *gitlabClient) {
	reg("create_note", "Create a new note (comment) to an issue or merge request",
		func(ctx context.Context, args CreateNoteArgs) (*mcp_golang.ToolResponse, error) {
			result, err := client.CreateNote(args.ProjectID, args.NoteableType, args.NoteableIID, args.Body)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("create_merge_request_thread", "Create a new thread on a merge request",
		func(ctx context.Context, args CreateMergeRequestThreadArgs) (*mcp_golang.ToolResponse, error) {
			var pos *gitlab.NotePosition
			if args.Position != nil {
				pos = &gitlab.NotePosition{
					BaseSHA:      args.Position.BaseSHA,
					HeadSHA:      args.Position.HeadSHA,
					StartSHA:     args.Position.StartSHA,
					PositionType: args.Position.PositionType,
				}
				if args.Position.NewPath != nil {
					pos.NewPath = *args.Position.NewPath
				}
				if args.Position.OldPath != nil {
					pos.OldPath = *args.Position.OldPath
				}
				if args.Position.NewLine != nil {
					pos.NewLine = args.Position.NewLine
				}
				if args.Position.OldLine != nil {
					pos.OldLine = args.Position.OldLine
				}
			}
			createdAt := ""
			if args.CreatedAt != nil {
				createdAt = *args.CreatedAt
			}
			result, err := client.CreateMergeRequestThread(args.ProjectID, args.MergeRequestIID, args.Body, pos, createdAt)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("mr_discussions", "List discussion items for a merge request",
		func(ctx context.Context, args ListMRDiscussionsArgs) (*mcp_golang.ToolResponse, error) {
			page := 0
			if args.Page != nil {
				page = *args.Page
			}
			perPage := 0
			if args.PerPage != nil {
				perPage = *args.PerPage
			}
			result, err := client.ListMergeRequestDiscussions(args.ProjectID, args.MergeRequestIID, page, perPage)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("update_merge_request_note", "Modify an existing merge request thread note",
		func(ctx context.Context, args UpdateMergeRequestNoteArgs) (*mcp_golang.ToolResponse, error) {
			body := ""
			if args.Body != nil {
				body = *args.Body
			}
			result, err := client.UpdateMergeRequestNote(args.ProjectID, args.MergeRequestIID, args.DiscussionID, args.NoteID, body, args.Resolved)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("create_merge_request_note", "Add a new note to an existing merge request thread",
		func(ctx context.Context, args CreateMergeRequestNoteArgs) (*mcp_golang.ToolResponse, error) {
			createdAt := ""
			if args.CreatedAt != nil {
				createdAt = *args.CreatedAt
			}
			result, err := client.CreateMergeRequestNote(args.ProjectID, args.MergeRequestIID, args.DiscussionID, args.Body, createdAt)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("update_issue_note", "Modify an existing issue thread note",
		func(ctx context.Context, args UpdateIssueNoteArgs) (*mcp_golang.ToolResponse, error) {
			result, err := client.UpdateIssueNote(args.ProjectID, args.IssueIID, args.DiscussionID, args.NoteID, args.Body)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("create_issue_note", "Add a new note to an existing issue thread",
		func(ctx context.Context, args CreateIssueNoteArgs) (*mcp_golang.ToolResponse, error) {
			createdAt := ""
			if args.CreatedAt != nil {
				createdAt = *args.CreatedAt
			}
			result, err := client.CreateIssueNote(args.ProjectID, args.IssueIID, args.DiscussionID, args.Body, createdAt)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})
}
