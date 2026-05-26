package tools

import (
	"context"

	mcp_golang "github.com/metoro-io/mcp-golang"
)

type ListNamespacesArgs struct {
	Search *string `json:"search,omitempty" jsonschema:"description=Search term for namespaces"`
	Owned  *bool   `json:"owned,omitempty" jsonschema:"description=Filter for namespaces owned by current user"`
	Page   *int    `json:"page,omitempty" jsonschema:"description=Page number"`
	PerPage *int   `json:"per_page,omitempty" jsonschema:"description=Number of items per page"`
}

type GetNamespaceArgs struct {
	NamespaceID string `json:"namespace_id" jsonschema:"required,description=Namespace ID or full path"`
}

type VerifyNamespaceArgs struct {
	Path     string `json:"path" jsonschema:"required,description=Namespace path to verify"`
	ParentID *int   `json:"parent_id,omitempty" jsonschema:"description=Parent namespace ID"`
}

type GetProjectArgs struct {
	ProjectID string `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
}

type ListProjectsArgs struct {
	Search                 *string `json:"search,omitempty" jsonschema:"description=Search term for projects"`
	SearchNamespaces       *bool   `json:"search_namespaces,omitempty" jsonschema:"description=Needs to be true if search is full path"`
	Owned                  *bool   `json:"owned,omitempty" jsonschema:"description=Filter for projects owned by current user"`
	Membership             *bool   `json:"membership,omitempty" jsonschema:"description=Filter for projects where current user is a member"`
	Archived               *bool   `json:"archived,omitempty" jsonschema:"description=Filter for archived projects"`
	Visibility             *string `json:"visibility,omitempty" jsonschema:"description=Filter by project visibility (public, internal, private)"`
	OrderBy                *string `json:"order_by,omitempty" jsonschema:"description=Order by (id, name, path, created_at, updated_at, last_activity_at)"`
	Sort                   *string `json:"sort,omitempty" jsonschema:"description=Sort direction (asc, desc)"`
	MinAccessLevel         *int    `json:"min_access_level,omitempty" jsonschema:"description=Filter by minimum access level"`
	Page                   *int    `json:"page,omitempty" jsonschema:"description=Page number"`
	PerPage                *int    `json:"per_page,omitempty" jsonschema:"description=Number of items per page"`
}

type ListGroupProjectsArgs struct {
	GroupID              string  `json:"group_id" jsonschema:"required,description=Group ID or path"`
	IncludeSubgroups     *bool   `json:"include_subgroups,omitempty" jsonschema:"description=Include projects from subgroups"`
	Search               *string `json:"search,omitempty" jsonschema:"description=Search term to filter projects"`
	OrderBy              *string `json:"order_by,omitempty" jsonschema:"description=Field to sort by (name, path, created_at)"`
	Sort                 *string `json:"sort,omitempty" jsonschema:"description=Sort direction (asc, desc)"`
	Archived             *bool   `json:"archived,omitempty" jsonschema:"description=Filter for archived projects"`
	Visibility           *string `json:"visibility,omitempty" jsonschema:"description=Filter by project visibility"`
	Page                 *int    `json:"page,omitempty" jsonschema:"description=Page number"`
	PerPage              *int    `json:"per_page,omitempty" jsonschema:"description=Number of items per page"`
}

type ListLabelsArgs struct {
	ProjectID            string `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	WithCounts           *bool  `json:"with_counts,omitempty" jsonschema:"description=Include issue and merge request counts"`
	IncludeAncestorGroups *bool `json:"include_ancestor_groups,omitempty" jsonschema:"description=Include ancestor groups"`
	Search               *string `json:"search,omitempty" jsonschema:"description=Keyword to filter labels by"`
}

type GetLabelArgs struct {
	ProjectID            string `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	LabelID              string `json:"label_id" jsonschema:"required,description=The ID or title of a project label"`
	IncludeAncestorGroups *bool `json:"include_ancestor_groups,omitempty" jsonschema:"description=Include ancestor groups"`
}

type CreateLabelArgs struct {
	ProjectID   string  `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	Name        string  `json:"name" jsonschema:"required,description=The name of the label"`
	Color       string  `json:"color" jsonschema:"required,description=The color in 6-digit hex notation with leading # sign"`
	Description *string `json:"description,omitempty" jsonschema:"description=The description of the label"`
	Priority    *int    `json:"priority,omitempty" jsonschema:"description=The priority of the label"`
}

type UpdateLabelArgs struct {
	ProjectID   string  `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	LabelID     string  `json:"label_id" jsonschema:"required,description=The ID or title of a project label"`
	NewName     *string `json:"new_name,omitempty" jsonschema:"description=The new name of the label"`
	Color       *string `json:"color,omitempty" jsonschema:"description=The color in 6-digit hex notation with leading # sign"`
	Description *string `json:"description,omitempty" jsonschema:"description=The new description of the label"`
	Priority    *int    `json:"priority,omitempty" jsonschema:"description=The new priority of the label"`
}

type DeleteLabelArgs struct {
	ProjectID string `json:"project_id" jsonschema:"required,description=Project ID or URL-encoded path"`
	LabelID   string `json:"label_id" jsonschema:"required,description=The ID or title of a project label"`
}

type GetUsersArgs struct {
	Usernames []string `json:"usernames" jsonschema:"required,description=Array of usernames to search for"`
}

func registerNamespaceTools(reg toolReg, client *gitlabClient) {
	reg("list_namespaces", "List all namespaces available to the current user",
		func(ctx context.Context, args ListNamespacesArgs) (*mcp_golang.ToolResponse, error) {
			search := ""
			if args.Search != nil {
				search = *args.Search
			}
			owned := false
			if args.Owned != nil {
				owned = *args.Owned
			}
			page := 0
			if args.Page != nil {
				page = *args.Page
			}
			perPage := 0
			if args.PerPage != nil {
				perPage = *args.PerPage
			}
			result, err := client.ListNamespaces(search, owned, page, perPage)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("get_namespace", "Get details of a namespace by ID or path",
		func(ctx context.Context, args GetNamespaceArgs) (*mcp_golang.ToolResponse, error) {
			result, err := client.GetNamespace(args.NamespaceID)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("verify_namespace", "Verify if a namespace path exists",
		func(ctx context.Context, args VerifyNamespaceArgs) (*mcp_golang.ToolResponse, error) {
			parentID := 0
			if args.ParentID != nil {
				parentID = *args.ParentID
			}
			result, err := client.VerifyNamespace(args.Path, parentID)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})
}

func registerProjectTools(reg toolReg, client *gitlabClient) {
	reg("get_project", "Get details of a specific project",
		func(ctx context.Context, args GetProjectArgs) (*mcp_golang.ToolResponse, error) {
			result, err := client.GetProject(args.ProjectID)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("list_projects", "List projects accessible by the current user",
		func(ctx context.Context, args ListProjectsArgs) (*mcp_golang.ToolResponse, error) {
			params := toURLValues(args)
			result, err := client.ListProjects(params)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("list_group_projects", "List projects in a GitLab group with filtering options",
		func(ctx context.Context, args ListGroupProjectsArgs) (*mcp_golang.ToolResponse, error) {
			params := toURLValues(args)
			result, err := client.ListGroupProjects(args.GroupID, params)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})
}

func registerLabelTools(reg toolReg, client *gitlabClient) {
	reg("list_labels", "List labels for a project",
		func(ctx context.Context, args ListLabelsArgs) (*mcp_golang.ToolResponse, error) {
			var withCounts *bool
			var includeAncestor *bool
			if args.WithCounts != nil {
				withCounts = args.WithCounts
			}
			if args.IncludeAncestorGroups != nil {
				includeAncestor = args.IncludeAncestorGroups
			}
			search := ""
			if args.Search != nil {
				search = *args.Search
			}
			result, err := client.ListLabels(args.ProjectID, withCounts, includeAncestor, search)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("get_label", "Get a single label from a project",
		func(ctx context.Context, args GetLabelArgs) (*mcp_golang.ToolResponse, error) {
			include := false
			if args.IncludeAncestorGroups != nil {
				include = *args.IncludeAncestorGroups
			}
			result, err := client.GetLabel(args.ProjectID, args.LabelID, include)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("create_label", "Create a new label in a project",
		func(ctx context.Context, args CreateLabelArgs) (*mcp_golang.ToolResponse, error) {
			desc := ""
			if args.Description != nil {
				desc = *args.Description
			}
			result, err := client.CreateLabel(args.ProjectID, args.Name, args.Color, desc, args.Priority)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("update_label", "Update an existing label in a project",
		func(ctx context.Context, args UpdateLabelArgs) (*mcp_golang.ToolResponse, error) {
			newName := ""
			if args.NewName != nil {
				newName = *args.NewName
			}
			color := ""
			if args.Color != nil {
				color = *args.Color
			}
			desc := ""
			if args.Description != nil {
				desc = *args.Description
			}
			result, err := client.UpdateLabel(args.ProjectID, args.LabelID, newName, color, desc, args.Priority)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})

	reg("delete_label", "Delete a label from a project",
		func(ctx context.Context, args DeleteLabelArgs) (*mcp_golang.ToolResponse, error) {
			if err := client.DeleteLabel(args.ProjectID, args.LabelID); err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(
				toJSON(map[string]string{"status": "success", "message": "Label deleted successfully"}))), nil
		})
}

func registerUserTools(reg toolReg, client *gitlabClient) {
	reg("get_users", "Get GitLab user details by usernames",
		func(ctx context.Context, args GetUsersArgs) (*mcp_golang.ToolResponse, error) {
			result, err := client.GetUsers(args.Usernames)
			if err != nil {
				return nil, err
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(toJSON(result))), nil
		})
}
