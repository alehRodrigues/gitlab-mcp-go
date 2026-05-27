package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Token          string
	APIURL         string
	ReadOnly       bool
	UseWiki        bool
	UseMilestone   bool
	UsePipeline    bool
}

func Load() *Config {
	loadEnv()
	return &Config{
		Token:          os.Getenv("GITLAB_PERSONAL_ACCESS_TOKEN"),
		APIURL:         normalizeURL(os.Getenv("GITLAB_API_URL")),
		ReadOnly:       os.Getenv("GITLAB_READ_ONLY_MODE") == "true",
		UseWiki:        os.Getenv("USE_GITLAB_WIKI") == "true",
		UseMilestone:   os.Getenv("USE_MILESTONE") == "true",
		UsePipeline:    os.Getenv("USE_PIPELINE") == "true",
	}
}

func loadEnv() {
	exe, err := os.Executable()
	if err == nil {
		godotenv.Load(filepath.Join(filepath.Dir(exe), ".env"))
	}
}

func normalizeURL(url string) string {
	if url == "" {
		return "https://gitlab.com/api/v4"
	}
	url = strings.TrimRight(url, "/")
	if !strings.HasSuffix(url, "/api/v4") {
		url += "/api/v4"
	}
	return url
}
