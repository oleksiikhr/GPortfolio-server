package api

import (
	"bytes"
	"encoding/json"
	"github.com/go-redis/redis/v7"
	"net/http"
	"net/url"

	"github.com/GPortfolio/server/config"
)

// githubDomain for main website
const githubDomain = "https://github.com"

type githubOauthSuccess struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

type githubOauthFailed struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	ErrorUri         string `json:"error_uri"`
}

// GithubRoutes register routes for Github
func GithubRoutes(redis redis.Client) {
	http.HandleFunc("/github/oauth/redirect", handleGithubRedirect)
	http.HandleFunc("/github/oauth/accept", func(w http.ResponseWriter, r *http.Request) {
		handleGithubAccept(w, r, redis)
	})
}

// handleRedirect user to Github oauth page
func handleGithubRedirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, githubDomain+"/login/oauth/authorize"+
		"?client_id="+config.Env("GITHUB_APP_ID", "")+
		"&redirect_uri="+url.QueryEscape(config.GetUrlWebsite()+"/github/oauth/accept")+
		"&state="+r.Header.Get("APP_TOKEN"), http.StatusMovedPermanently)
}

// handleGithubAccept user redirected after oauth page on Github website
func handleGithubAccept(w http.ResponseWriter, r *http.Request, redis redis.Client) {
	requestBody, _ := json.Marshal(map[string]string{
		"client_id":     config.Env("GITHUB_APP_ID", ""),
		"client_secret": config.Env("GITHUB_APP_SECRET", ""),
		"code":          r.URL.Query().Get("code"),
	})

	// TODO Validation

	// Send request to get access token from received code
	req, err := http.NewRequest(http.MethodPost, githubDomain+"/login/oauth/access_token", bytes.NewBuffer(requestBody))
	if err != nil {
		respJsonFailed(w, err.Error())
		return
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		respJsonFailed(w, err.Error())
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		var oauthResp githubOauthFailed
		err = json.NewDecoder(resp.Body).Decode(&oauthResp)
		if err != nil {
			respJsonFailed(w, err.Error())
			return
		}

		respJsonFailed(w, oauthResp.ErrorDescription)
		return
	}

	var oauthResp githubOauthSuccess
	err = json.NewDecoder(resp.Body).Decode(&oauthResp)
	if err != nil {
		respJsonFailed(w, err.Error())
		return
	}

	// TODO Store on Redis
	w.Write([]byte("Token received, you can close the tab"))
}
