package main

import (
	"fmt"
	"os"

	mcp_golang "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"

	"github.com/user/gitlab-mcp-go/internal/config"
	"github.com/user/gitlab-mcp-go/internal/gitlab"
	"github.com/user/gitlab-mcp-go/internal/tools"
)

func main() {
	cfg := config.Load()

	if cfg.Token == "" {
		fmt.Fprintf(os.Stderr, "GITLAB_PERSONAL_ACCESS_TOKEN environment variable is not set\n")
		os.Exit(1)
	}

	client := gitlab.NewClient(cfg.APIURL, cfg.Token)

	server := mcp_golang.NewServer(stdio.NewStdioServerTransport())

	tools.RegisterAll(server, client, cfg)

	if err := server.Serve(); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
		os.Exit(1)
	}

	select {}
}
