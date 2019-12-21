package github

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"time"

	"github.com/GPortfolio/server/config"
)

const (
	// Domain of main website
	Domain = "https://github.com"
	// ApiDomain for send json requests
	ApiDomain = "https://api.github.com"
)

// OauthResponse after receiving access token
type OauthResponse struct {
	AccessToken      string `json:"access_token"`
	Scope            string `json:"scope"`
	TokenType        string `json:"token_type"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	ErrorUri         string `json:"error_uri"`
}

// ProfileResponse - User Profile from Github API
// https://developer.github.com/v3/users/#response-with-public-and-private-profile-information
type ProfileResponse struct {
	Login             string    `json:"login"`
	Id                int64     `json:"id"`
	NodeId            string    `json:"node_id"`
	AvatarUrl         string    `json:"avatar_url"`
	HtmlUrl           string    `json:"html_url"`
	Type              string    `json:"type"`
	Name              string    `json:"name"`
	Company           string    `json:"company"`
	Location          string    `json:"location"`
	Email             string    `json:"email"`
	Hireable          bool      `json:"hireable"`
	Bio               string    `json:"bio"`
	PublicRepos       int       `json:"public_repos"`
	PublicGists       int       `json:"public_gists"`
	Followers         int       `json:"followers"`
	Following         int       `json:"following"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	PrivateGists      int       `json:"private_gists"`
	TotalPrivateRepos int       `json:"total_private_repos"`
	OwnedPrivateRepos int       `json:"owned_private_repos"`
	DiskUsage         int       `json:"disk_usage"`
	Collaborators     int       `json:"collaborators"`
}

// GetOauthRedirectUrl of website
func GetOauthRedirectUrl() string {
	return Domain + "/login/oauth/authorize" +
		"?client_id=" + config.Env("GITHUB_APP_ID", "") +
		"&redirect_uri=" + url.QueryEscape(config.GetUrlWebsite()+"/api/github/oauth/accept")
}

// GenerateOauthBody for get access token
func GenerateOauthBody(code string) ([]byte, error) {
	return json.Marshal(map[string]string{
		"client_secret": config.Env("GITHUB_APP_SECRET", ""),
		"client_id":     config.Env("GITHUB_APP_ID", ""),
		"code":          code,
	})
}

// FetchAccessToken send request to get access token from received code
// from oauth body (code)
func FetchAccessToken(body []byte) (*OauthResponse, error) {
	req, err := http.NewRequest(http.MethodPost, Domain+"/login/oauth/access_token", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	setJsonHeader(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var oauthResp *OauthResponse
	err = json.NewDecoder(resp.Body).Decode(&oauthResp)
	if err != nil {
		return nil, err
	}

	if oauthResp.Error != "" {
		return nil, errors.New(oauthResp.ErrorDescription)
	}

	return oauthResp, nil
}

// FetchProfile from Github API
func FetchProfile(token string) (*ProfileResponse, error) {
	req, err := http.NewRequest(http.MethodGet, ApiDomain+"/user", nil)
	if err != nil {
		return nil, err
	}

	setJsonHeader(req)
	req.Header.Set("Authorization", "token "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var profileResp *ProfileResponse
	err = json.NewDecoder(resp.Body).Decode(&profileResp)
	if err != nil {
		return nil, err
	}

	return profileResp, nil
}

// setJsonHeader accept and send application/json format
func setJsonHeader(req *http.Request) {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
}
