package server

import (
	"net/http"

	"golang.org/x/oauth2"
)

type Router struct {
	*http.ServeMux
	oauthConfig   *oauth2.Config
	gc            githubClient
	rs            reposStore
	ms            monitoringStore
	frontendURL   string
	backendDomain string
}

func NewRouter(oauthConfig *oauth2.Config, gc githubClient, rs reposStore, ms monitoringStore, frontendURL, backendDomain string) *Router {
	rt := &Router{
		ServeMux:      http.NewServeMux(),
		oauthConfig:   oauthConfig,
		gc:            gc,
		rs:            rs,
		ms:            ms,
		frontendURL:   frontendURL,
		backendDomain: backendDomain,
	}

	rt.HandleFunc("GET /api/repos", rt.getUserRepos)
	rt.HandleFunc("POST /api/repos", rt.createRepo)
	rt.HandleFunc("DELETE /api/repos/{id}", rt.deleteRepo)
	rt.HandleFunc("GET /api/dashboard/repos", rt.getDashboard)

	rt.HandleFunc("POST /api/monitoring", rt.addMonitoring)
	rt.HandleFunc("GET /api/monitoring", rt.getMonitoringUrls)
	rt.HandleFunc("PUT /api/monitoring/{id}", rt.updateMonitoringURL)
	rt.HandleFunc("DELETE /api/monitoring/{id}", rt.deleteMonitoringURL)
	rt.HandleFunc("GET /api/monitoring/ping", rt.getAllMonitoringUrlsToPing)

	rt.HandleFunc("GET /auth/github/login", rt.oauthLogin)
	rt.HandleFunc("GET /auth/callback", rt.oauthCallback)

	return rt
}
