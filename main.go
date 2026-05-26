package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	mcp_golang "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"

	"github.com/user/gitlab-mcp-go/internal/config"
	"github.com/user/gitlab-mcp-go/internal/gitlab"
	"github.com/user/gitlab-mcp-go/internal/tools"
)

func main() {
	cfg := config.Load()

	fmt.Fprintf(os.Stderr, "GitLab MCP Server starting...\n")
	fmt.Fprintf(os.Stderr, "API URL: %s\n", cfg.APIURL)
	if cfg.ReadOnly {
		fmt.Fprintf(os.Stderr, "Mode: read-only\n")
	}

	if cfg.Token == "" {
		fmt.Fprintf(os.Stderr, "GITLAB_PERSONAL_ACCESS_TOKEN is not set\n")
		os.Exit(1)
	}

	client := gitlab.NewClient(cfg.APIURL, cfg.Token)
	server := mcp_golang.NewServer(stdio.NewStdioServerTransport())
	tools.RegisterAll(server, client, cfg)

	go func() {
		if err := server.Serve(); err != nil {
			fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
			os.Exit(1)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
