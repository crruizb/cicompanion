package api

import (
	"net/http"

	"golang.org/x/oauth2"
)

type Router struct {
	*http.ServeMux
	oauthConfig  *oauth2.Config
	githubClient githubClient
}

func NewRouter(oauthConfig *oauth2.Config, githubClient githubClient) *Router {
	rt := &Router{
		ServeMux:     http.NewServeMux(),
		oauthConfig:  oauthConfig,
		githubClient: githubClient,
	}

	rt.HandleFunc("GET /repos", rt.getUserRepos)
	rt.HandleFunc("POST /repos/{id}", rt.createRepo)
	rt.HandleFunc("DELETE /repos/{id}", rt.deleteRepo)
	rt.HandleFunc("GET /dashboard/repos", rt.getDashboard)

	rt.HandleFunc("GET /auth/github/login", rt.oauthLogin)
	rt.HandleFunc("GET /auth/callback", rt.oauthCallback)

	return rt
}
