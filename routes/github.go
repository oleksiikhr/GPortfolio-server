package routes

import (
	"net/http"
	"strings"
	"time"

	"github.com/GPortfolio/server/services/github"
)

// githubRoutes register routes for Github
func (h *Handlers) githubRoutes() {
	http.HandleFunc("/api/github/oauth/redirect", h.handleGithubRedirect())
	http.HandleFunc("/api/github/oauth/accept", h.handleGithubAccept())
	// TODO
	// /api/github/profile
	// /api/github/repositories
}

// handleGithubRedirect user to Github oauth page
func (h *Handlers) handleGithubRedirect() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, github.GetOauthRedirectUrl(), http.StatusMovedPermanently)
	}
}

// handleGithubAccept user redirected after oauth page on Github website
func (h *Handlers) handleGithubAccept() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := strings.TrimSpace(r.URL.Query().Get("code"))
		if code == "" {
			responseQuick(w, "Code is empty", http.StatusBadRequest)
			return
		}

		requestBody, err := github.GenerateOauthBody(code)
		if err != nil {
			h.Logger.Println(err)
			responseQuick(w, "Something went wrong", http.StatusBadRequest)
			return
		}

		resp, err := github.FetchAccessToken(requestBody)
		if err != nil {
			responseQuick(w, err.Error(), http.StatusBadRequest)
			return
		}

		defer resp.Body.Close()

		oauthResponse, err := github.GetOauthAnswerFromResponse(resp)
		if err != nil {
			responseQuick(w, err.Error(), http.StatusBadRequest)
			return
		}

		keyPass, err := h.Redis.RndSet(oauthResponse.AccessToken, time.Hour*24)
		if err != nil {
			h.Logger.Println(err)
			responseQuick(w, err.Error(), http.StatusBadRequest)
			return
		}

		responseKeyPass(w, keyPass.Key, keyPass.Pass)
	}
}
