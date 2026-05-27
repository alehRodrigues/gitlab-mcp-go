# AGENTS.md

## Build

```shell
go build -o build/gitlab-mcp-go .
go vet ./...
go mod tidy
go test ./...  # no tests exist yet
```

Output goes to `build/`. No linter/formatter config exists. No `golangci-lint` in use.

## Config

All config via env vars (`.env` loaded by godotenv at startup):

| Var | Required | Default |
|---|---|---|
| `GITLAB_PERSONAL_ACCESS_TOKEN` | yes | — |
| `GITLAB_API_URL` | no | `https://gitlab.com/api/v4` |
| `GITLAB_READ_ONLY_MODE` | no | `false` |
| `USE_GITLAB_WIKI` | no | `false` |
| `USE_MILESTONE` | no | `false` |
| `USE_PIPELINE` | no | `false` |

Feature-toggled tools (wiki, milestone, pipeline) are hidden from the MCP tool list unless the corresponding env var is `"true"`.

`GITLAB_API_URL` auto-appends `/api/v4` if missing. `GITLAB_READ_ONLY_MODE=true` exposes only read-only tools (see `readOnlyTools` list in `internal/tools/registry.go:78`).

## Architecture

- `main.go` — entrypoint. Creates MCP server over stdio transport, registers all tools, blocks on signal.
- `internal/config/config.go` — loads env vars into `Config` struct.
- `internal/gitlab/client.go` — thin HTTP client for GitLab REST API v4. No official GitLab SDK; raw `net/http` with JSON marshal/unmarshal.
- `internal/gitlab/*.go` — one file per domain (issues, merge_requests, pipelines, etc.), plus `types.go` for all response structs.
- `internal/tools/registry.go` — `RegisterAll()` wires tools to MCP server. Has a closure-based `enabled()` check that skips feature-gated tools. `toURLValues()` converts args structs to query params via reflection (skips `project_id` field intentionally).
- `internal/tools/*_tools.go` — one file per domain, each contains arg structs + `register*Tools()` function.

MCP framework: `github.com/metoro-io/mcp-golang` v0.16.1, transport `stdio.NewStdioServerTransport()`. Tool responses always return `toJSON(result)` pretty-printed JSON.

## Project quirks

- Module path in `go.mod` is `github.com/user/gitlab-mcp-go` — a placeholder. Do not treat it as the canonical import path without confirming.
- `internal/tools/toURLValues()` explicitly skips fields tagged `json:"project_id"` — project IDs are passed as positional args, not query params.
- Project IDs are passed as strings and URL-decoded via `client.DecodeProjectID()` to handle URL-encoded paths like `group%2Fsubgroup%2Fproject`.
- No tests exist yet. Unit tests go in `*_test.go` files alongside the package they test.
- Go 1.26.2 — verify compatibility if introducing new stdlib features.

## Workflow

After any code change:
1. Update `README.md` if new tools, config vars, or features were added
2. Run `go build -o build/gitlab-mcp-go .` to verify compilation
3. Run `go vet ./...` to check for issues
4. Run `go mod tidy` to keep go.mod/go.sum clean
5. Stage changes (`git add -A`) and commit with a descriptive message
