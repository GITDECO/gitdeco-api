package pkg

import (
	"encoding/json"
	"gitdeco-api/configs"
	"gitdeco-api/internal/exception"
	"gitdeco-api/internal/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
)

func GithubLogin(c *fiber.Ctx, gitHubState string) string {
	return configs.GithubOAuthConfig.AuthCodeURL(gitHubState, oauth2.AccessTypeOnline)
}

func GithubCallback(c *fiber.Ctx) *models.User {
	token, err := configs.GithubOAuthConfig.Exchange(c.Context(), c.Query("code"))
	if err != nil {
		panic(&exception.Error{Key: "OAUTH_ERROR", Data: err.Error()})
	}

	client := configs.GithubOAuthConfig.Client(c.Context(), token)
	userInfo, err := client.Get("https://api.github.com/user")
	if err != nil {
		panic(&exception.Error{Key: "OAUTH_ERROR", Data: err.Error()})
	}
	defer userInfo.Body.Close()

	var githubUser map[string]interface{}
	err = json.NewDecoder(userInfo.Body).Decode(&githubUser)
	if err != nil {
		panic(&exception.Error{Key: "OAUTH_ERROR", Data: err.Error()})
	}

	user := new(models.User)
	user.Username = githubUser["login"].(string)
	if githubUser["email"] != nil {
		user.Email = githubUser["email"].(string)
	}
	if githubUser["name"] != nil {
		user.Name = githubUser["name"].(string)
	}
	if githubUser["bio"] != nil {
		user.Bio = githubUser["bio"].(string)
	}
	if githubUser["avatar_url"] != nil {
		user.Avatar = githubUser["avatar_url"].(string)
	}
	return user
}
