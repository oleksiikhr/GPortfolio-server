package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/GPortfolio/server/services/github"
)

// githubRoutes register routes for Github
func (h *Handlers) githubRoutes() {
	http.HandleFunc("/api/github/oauth/redirect", h.handleGithubRedirect())
	http.HandleFunc("/api/github/oauth/accept", h.handleGithubAccept())
	http.HandleFunc("/api/github/profile", h.handleGithubProfile())
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

		oauthResponse, err := github.FetchAccessToken(requestBody)
		if err != nil {
			responseQuick(w, err.Error(), http.StatusBadRequest)
			return
		}

		keyPass, err := h.Redis.SecSet(oauthResponse.AccessToken, time.Hour*24)
		if err != nil {
			h.Logger.Println(err)
			responseQuick(w, err.Error(), http.StatusBadRequest)
			return
		}

		responseQuick(w, map[string]string{
			"key": keyPass.Key,
			"pass": keyPass.Pass,
		}, http.StatusOK)
	}
}

// handleGithubProfile get profile for auth user
func (h *Handlers) handleGithubProfile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		value := h.tryAccessToken(r)
		if value == nil {
			responseQuick(w, "No access", http.StatusBadRequest)
			return
		}

		token := fmt.Sprintf("%v", value)

		profile, err := github.FetchProfile(token)
		if err != nil {
			responseQuick(w, "Profile not found", http.StatusBadRequest)
			return
		}

		response := response(w, "Profile received", http.StatusOK)
		response["data"] = profile

		json.NewEncoder(w).Encode(response)
	}
}
