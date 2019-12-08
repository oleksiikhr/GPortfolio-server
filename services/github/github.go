package github

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/GPortfolio/server/config"
)

const (
	// Domain of main website
	Domain = "https://github.com"
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
func FetchAccessToken(body []byte) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, Domain+"/login/oauth/access_token", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetOauthAnswerFromResponse
func GetOauthAnswerFromResponse(resp *http.Response) (OauthResponse, error) {
	var oauthResp OauthResponse
	err := json.NewDecoder(resp.Body).Decode(&oauthResp)
	if err != nil {
		return oauthResp, err
	}

	if oauthResp.Error != "" {
		return oauthResp, errors.New(oauthResp.ErrorDescription)
	}

	return oauthResp, nil
}
