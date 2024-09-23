package configs

import (
	"fmt"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var GithubOAuthConfig = &oauth2.Config{
	ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
	ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
	RedirectURL:  fmt.Sprintf("%s://%s:%s/%s/api%s", os.Getenv("SERVER_PROTOCOL"), os.Getenv("SERVER_DOMAIN"), os.Getenv("SERVER_PORT"), os.Getenv("SERVER_VERSION"), os.Getenv("GITHUB_REDIRECT_URL")),
	Scopes:       []string{"user:email"},
	Endpoint:     github.Endpoint,
}
