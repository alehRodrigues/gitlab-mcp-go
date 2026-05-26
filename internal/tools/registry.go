package tools

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	mcp_golang "github.com/metoro-io/mcp-golang"

	"github.com/user/gitlab-mcp-go/internal/config"
	"github.com/user/gitlab-mcp-go/internal/gitlab"
)

func RegisterAll(server *mcp_golang.Server, client *gitlab.Client, cfg *config.Config) {
	wikiTools := []string{"list_wiki_pages", "get_wiki_page", "create_wiki_page", "update_wiki_page", "delete_wiki_page"}
	milestoneTools := []string{"list_milestones", "get_milestone", "create_milestone", "edit_milestone",
		"delete_milestone", "get_milestone_issue", "get_milestone_merge_requests",
		"promote_milestone", "get_milestone_burndown_events"}
	pipelineTools := []string{"list_pipelines", "get_pipeline", "list_pipeline_jobs",
		"get_pipeline_job", "get_pipeline_job_output", "create_pipeline", "retry_pipeline", "cancel_pipeline"}

	enabled := func(name string) bool {
		if cfg.ReadOnly {
			for _, r := range readOnlyTools {
				if r == name {
					goto checkToggles
				}
			}
			return false
		}
	checkToggles:
		for _, w := range wikiTools {
			if w == name {
				return cfg.UseWiki
			}
		}
		for _, m := range milestoneTools {
			if m == name {
				return cfg.UseMilestone
			}
		}
		for _, p := range pipelineTools {
			if p == name {
				return cfg.UsePipeline
			}
		}
		return true
	}

	reg := func(name, desc string, handler any) {
		if !enabled(name) {
			return
		}
		if err := server.RegisterTool(name, desc, handler); err != nil {
			panic(fmt.Sprintf("register tool %s: %v", name, err))
		}
	}

	registerRepoAndFileTools(reg, client)
	registerBranchTools(reg, client)
	registerProjectTools(reg, client)
	registerIssueTools(reg, client)
	registerMRTools(reg, client)
	registerNoteTools(reg, client)
	registerNamespaceTools(reg, client)
	registerLabelTools(reg, client)
	registerWikiTools(reg, client)
	registerPipelineTools(reg, client)
	registerMilestoneTools(reg, client)
	registerUserTools(reg, client)
}

type gitlabClient = gitlab.Client

var readOnlyTools = []string{
	"search_repositories", "get_file_contents", "get_merge_request", "get_merge_request_diffs",
	"get_branch_diffs", "mr_discussions", "list_issues", "list_merge_requests", "get_issue",
	"list_issue_links", "list_issue_discussions", "get_issue_link", "list_namespaces",
	"get_namespace", "verify_namespace", "get_project", "get_pipeline", "list_pipelines",
	"list_pipeline_jobs", "get_pipeline_job", "get_pipeline_job_output", "list_projects",
	"list_labels", "get_label", "list_group_projects", "get_repository_tree",
	"list_milestones", "get_milestone", "get_milestone_issue", "get_milestone_merge_requests",
	"get_milestone_burndown_events", "list_wiki_pages", "get_wiki_page", "get_users",
}

type toolReg func(name, desc string, handler any)

func toJSON(v any) string {
	b, _ := json.MarshalIndent(v, "", "  ")
	return string(b)
}

func ptr[T any](v T) *T {
	return &v
}

func toURLValues(v any) url.Values {
	params := url.Values{}
	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.Ptr && field.IsNil() {
			continue
		}

		jsonTag := typ.Field(i).Tag.Get("json")
		name := strings.Split(jsonTag, ",")[0]
		if name == "" || name == "-" || name == "project_id" {
			continue
		}

		if field.Kind() == reflect.Ptr {
			field = field.Elem()
		}

		switch field.Kind() {
		case reflect.String:
			params.Set(name, field.String())
		case reflect.Bool:
			params.Set(name, strconv.FormatBool(field.Bool()))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			params.Set(name, strconv.FormatInt(field.Int(), 10))
		case reflect.Slice:
			if field.Type().Elem().Kind() == reflect.String {
				joined := ""
				for j := 0; j < field.Len(); j++ {
					if j > 0 {
						joined += ","
					}
					joined += field.Index(j).String()
				}
				if joined != "" {
					params.Set(name, joined)
				}
			}
		}
	}
	return params
}

