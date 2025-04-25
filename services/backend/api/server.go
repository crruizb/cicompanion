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

func (rt *Router) getUserRepos(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(ContextUser).(*data.User)
	repos, err := rt.githubClient.GetRepos(user.GithubPAT)
	if err != nil {
		ServerErrorResponse(w, r, err)
		return
	}

	WriteJSON(w, http.StatusOK, envelope{"repos": repos}, nil)
}

func (rt *Router) createRepo(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func (rt *Router) deleteRepo(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func (rt *Router) getDashboard(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(ContextUser).(*data.User)
	// repoIdStr := r.PathValue("repoId")
	// repoId, err := strconv.Atoi(repoIdStr)
	// if err != nil {
	// 	BadRequestResponse(w, r, err)
	// 	return
	// }
	actions, err := rt.githubClient.GetActions(user.GithubPAT, 344513894)
	if err != nil {
		ServerErrorResponse(w, r, err)
		return
	}

	WriteJSON(w, http.StatusOK, envelope{"actions": actions}, nil)
}
