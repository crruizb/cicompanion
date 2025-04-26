package api

import (
	"net/http"

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
}

func (rt *Router) getUserRepos(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(ContextUser).(*data.User)
	repos, err := rt.gc.GetRepos(user.GithubPAT)
	if err != nil {
		ServerErrorResponse(w, r, err)
		return
	}

	WriteJSON(w, http.StatusOK, envelope{"repos": repos}, nil)
}

func (rt *Router) createRepo(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(ContextUser).(*data.User)
	repo := &data.Repo{}
	err := ReadJSON(w, r, repo)
	if err != nil {
		BadRequestResponse(w, r, err)
		return
	}
	err = rt.rs.AddRepo(*repo, user.Id)
	if err != nil {
		ServerErrorResponse(w, r, err)
		return
	}

	WriteJSON(w, http.StatusAccepted, nil, nil)
}

func (rt *Router) deleteRepo(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

type DashboardDTO struct {
	Id      int                         `json:"id"`
	Name    string                      `json:"name"`
	Actions githubapi.GithubRepoActions `json:"actions"`
}

func (rt *Router) getDashboard(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(ContextUser).(*data.User)
	repos, err := rt.rs.GetRepos(user.Id)
	if err != nil {
		ServerErrorResponse(w, r, err)
		return
	}

	dashboardRepos := []DashboardDTO{}
	for _, repo := range repos {
		actions, err := rt.gc.GetActions(user.GithubPAT, repo.Id)
		if err != nil {
			ServerErrorResponse(w, r, err)
			return
		}

		dashboardDto := DashboardDTO{
			Id:      repo.Id,
			Name:    repo.Name,
			Actions: *actions,
		}
		dashboardRepos = append(dashboardRepos, dashboardDto)
	}

	WriteJSON(w, http.StatusOK, envelope{"repos": dashboardRepos}, nil)
}
