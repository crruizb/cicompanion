package server

import (
	"github.com/cicompanion/data"
	"github.com/cicompanion/githubapi"
)

type contextUserKey string

const ContextUser = contextUserKey("ctxUser")

type githubClient interface {
	GetRepos(githubPat string) ([]githubapi.GithubRepo, error)
	GetActions(githubPat string, repoId int) (*githubapi.GithubRepoActions, error)
}

type reposStore interface {
	AddRepo(repo data.Repo, userId string) error
	GetRepos(userId string) ([]data.Repo, error)
	Deleterepo(repoId int, userId string) error
}

type monitoringStore interface {
	AddMonitoringURLs(monitoringURLs []data.MonitoringURL, userId string) error
	GetMonitoringURLsByUserId(userId string) ([]data.MonitoringURL, error)
	GetAllMonitoringURLs() ([]data.MonitoringURL, error)
	UpdateMonitoringURL(id int, userId string, mURL data.MonitoringURL) (*data.MonitoringURL, error)
	DeleteMonitoringURL(id int, userId string) error
}
