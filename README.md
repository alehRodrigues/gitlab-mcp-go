# GitLab MCP Server (Go)

Servidor MCP para interagir com a API do GitLab. Port em Go do [@zereight/mcp-gitlab](https://github.com/zereight/mcp-gitlab).

## Configuração

| Variável | Padrão | Descrição |
|---|---|---|
| `GITLAB_PERSONAL_ACCESS_TOKEN` | — | Token de acesso pessoal do GitLab (obrigatório) |
| `GITLAB_API_URL` | `https://gitlab.com/api/v4` | URL da instância GitLab (self-hosted) |
| `GITLAB_READ_ONLY_MODE` | `false` | Expor apenas ferramentas de leitura |
| `USE_GITLAB_WIKI` | `false` | Habilitar ferramentas de wiki |
| `USE_MILESTONE` | `false` | Habilitar ferramentas de milestone |
| `USE_PIPELINE` | `false` | Habilitar ferramentas de pipeline |

## Build

Compilar para a plataforma atual:

```shell
go build -o build/gitlab-mcp-go .
```

### Cross-compilation

**Linux (amd64):**

```shell
GOOS=linux GOARCH=amd64 go build -o build/gitlab-mcp-go-linux .
```

**macOS (Intel):**

```shell
GOOS=darwin GOARCH=amd64 go build -o build/gitlab-mcp-go-darwin .
```

**macOS (Apple Silicon M1/M2/M3):**

```shell
GOOS=darwin GOARCH=arm64 go build -o build/gitlab-mcp-go-darwin-arm64 .
```

**Windows:**

```shell
go build -o build/gitlab-mcp-go.exe .
```

## Compatibilidade

Compatível com qualquer cliente MCP via stdio: **Visual Studio**, **VS Code**, **opencode**, **Claude Desktop**, **Cline**, **Cursor**, entre outros.

```json
{
  "mcpServers": {
    "gitlab": {
      "command": "caminho/para/gitlab-mcp-go",
      "env": {
        "GITLAB_PERSONAL_ACCESS_TOKEN": "seu-token",
        "GITLAB_API_URL": "https://gitlab.com/api/v4"
      }
    }
  }
}
```

## Ferramentas

- **Repositório**: create_or_update_file, search_repositories, create_repository, get_file_contents, push_files, get_repository_tree, fork_repository
- **Branches**: create_branch, get_branch_diffs
- **Merge Requests**: create_merge_request, get_merge_request, get_merge_request_diffs, update_merge_request, list_merge_requests, create_merge_request_thread, mr_discussions, update_merge_request_note, create_merge_request_note, create_note
- **Issues**: create_issue, list_issues, get_issue, update_issue, delete_issue, list_issue_links, list_issue_discussions, get_issue_link, create_issue_link, delete_issue_link, update_issue_note, create_issue_note
- **Namespaces**: list_namespaces, get_namespace, verify_namespace
- **Projetos**: get_project, list_projects, list_group_projects
- **Labels**: list_labels, get_label, create_label, update_label, delete_label
- **Wikis**: list_wiki_pages, get_wiki_page, create_wiki_page, update_wiki_page, delete_wiki_page
- **Pipelines**: list_pipelines, get_pipeline, list_pipeline_jobs, get_pipeline_job, get_pipeline_job_output, create_pipeline, retry_pipeline, cancel_pipeline
- **Milestones**: list_milestones, get_milestone, create_milestone, edit_milestone, delete_milestone, get_milestone_issue, get_milestone_merge_requests, promote_milestone, get_milestone_burndown_events
- **Usuários**: get_users

## Licença

MIT
