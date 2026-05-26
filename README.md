# GitLab MCP Server (Go)

MCP server for interacting with the GitLab API. Port of [@zereight/mcp-gitlab](https://github.com/zereight/mcp-gitlab) to Go.

## Usage

### Configuration

| Variable | Default | Description |
|---|---|---|
| `GITLAB_PERSONAL_ACCESS_TOKEN` | — | GitLab personal access token (required) |
| `GITLAB_API_URL` | `https://gitlab.com/api/v4` | GitLab instance URL |
| `GITLAB_READ_ONLY_MODE` | `false` | Only expose read-only tools |
| `USE_GITLAB_WIKI` | `false` | Enable wiki tools |
| `USE_MILESTONE` | `false` | Enable milestone tools |
| `USE_PIPELINE` | `false` | Enable pipeline tools |

### Claude Desktop / Cline / Cursor

```json
{
  "mcpServers": {
    "gitlab": {
      "command": "path/to/gitlab-mcp-go",
      "env": {
        "GITLAB_PERSONAL_ACCESS_TOKEN": "your-token",
        "GITLAB_API_URL": "https://gitlab.com/api/v4"
      }
    }
  }
}
```

## Build

```shell
go build -o build/gitlab-mcp-go .
```

## Tools

- **Repository**: create_or_update_file, search_repositories, create_repository, get_file_contents, push_files, get_repository_tree, fork_repository
- **Branches**: create_branch, get_branch_diffs
- **Merge Requests**: create_merge_request, get_merge_request, get_merge_request_diffs, update_merge_request, list_merge_requests, create_merge_request_thread, mr_discussions, update_merge_request_note, create_merge_request_note, create_note
- **Issues**: create_issue, list_issues, get_issue, update_issue, delete_issue, list_issue_links, list_issue_discussions, get_issue_link, create_issue_link, delete_issue_link, update_issue_note, create_issue_note
- **Namespaces**: list_namespaces, get_namespace, verify_namespace
- **Projects**: get_project, list_projects, list_group_projects
- **Labels**: list_labels, get_label, create_label, update_label, delete_label
- **Wikis**: list_wiki_pages, get_wiki_page, create_wiki_page, update_wiki_page, delete_wiki_page
- **Pipelines**: list_pipelines, get_pipeline, list_pipeline_jobs, get_pipeline_job, get_pipeline_job_output, create_pipeline, retry_pipeline, cancel_pipeline
- **Milestones**: list_milestones, get_milestone, create_milestone, edit_milestone, delete_milestone, get_milestone_issue, get_milestone_merge_requests, promote_milestone, get_milestone_burndown_events
- **Users**: get_users

## License

MIT
