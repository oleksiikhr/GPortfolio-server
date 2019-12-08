package routes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/GPortfolio/server/config"
)

// githubDomain for main website
const githubDomain = "https://github.com"

// githubOauthResponse after get user access token
type githubOauthResponse struct {
	AccessToken      string `json:"access_token"`
	Scope            string `json:"scope"`
	TokenType        string `json:"token_type"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	ErrorUri         string `json:"error_uri"`
}

// githubRoutes register routes for Github
func (h *Handlers) githubRoutes() {
	http.HandleFunc("/api/github/oauth/redirect", h.handleGithubRedirect)
	http.HandleFunc("/api/github/oauth/accept", h.handleGithubAccept)
}

// handleRedirect user to Github oauth page
func (h *Handlers) handleGithubRedirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, githubDomain+"/login/oauth/authorize"+
		"?client_id="+config.Env("GITHUB_APP_ID", "")+
		"&redirect_uri="+url.QueryEscape(config.GetUrlWebsite()+"/github/oauth/accept")+
		"&state="+r.Header.Get("APP_TOKEN"), http.StatusMovedPermanently)
}

// handleGithubAccept user redirected after oauth page on Github website
func (*Handlers) handleGithubAccept(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := json.Marshal(map[string]string{
		"client_secret": config.Env("GITHUB_APP_SECRET", ""),
		"client_id":     config.Env("GITHUB_APP_ID", ""),
		"code":          r.URL.Query().Get("code"),
	})

	// TODO Validation + state

	// Send request to get access token from received code
	req, err := http.NewRequest(http.MethodPost, githubDomain+"/login/oauth/access_token", bytes.NewBuffer(requestBody))
	if err != nil {
		response := response(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		response := response(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	defer resp.Body.Close()

	var oauthResp githubOauthResponse
	err = json.NewDecoder(resp.Body).Decode(&oauthResp)
	if err != nil {
		response := response(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	if oauthResp.Error != "" {
		response := response(w, oauthResp.ErrorDescription, http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// TODO Store on Redis to the current user
	// oauthResp.AccessToken
	w.Write([]byte("Token received, you can close the tab"))
}
