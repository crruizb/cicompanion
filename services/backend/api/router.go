package api

import (
	"net/http"

	"golang.org/x/oauth2"
)

type Router struct {
	*http.ServeMux
	oauthConfig *oauth2.Config
	gc          githubClient
	rs          reposStore
}

func NewRouter(oauthConfig *oauth2.Config, gc githubClient, rs reposStore) *Router {
	rt := &Router{
		ServeMux:    http.NewServeMux(),
		oauthConfig: oauthConfig,
		gc:          gc,
		rs:          rs,
	}

	rt.HandleFunc("GET /api/repos", rt.getUserRepos)
	rt.HandleFunc("POST /api/repos", rt.createRepo)
	rt.HandleFunc("DELETE /api/repos/{id}", rt.deleteRepo)
	rt.HandleFunc("GET /api/dashboard/repos", rt.getDashboard)

	rt.HandleFunc("GET /auth/github/login", rt.oauthLogin)
	rt.HandleFunc("GET /auth/callback", rt.oauthCallback)

	return rt
}
