package githubapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GithubHttpClient struct{}

func NewGithubHttpClient() *GithubHttpClient {
	return &GithubHttpClient{}
}

type GithubRepo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type GithubRepoActions struct {
	TotalCount   int                  `json:"total_count"`
	WorkflowRuns []GithubWorkflowRuns `json:"workflow_runs"`
}

type GithubWorkflowRuns struct {
	Branch       string `json:"head_branch"`
	DisplayTitle string `json:"display_title"`
	Status       string `json:"status"`
	Conclusion   string `json:"conclusion"`
	// CreatedAt    time.Time `json:"created_at"`
}

func (c *GithubHttpClient) GetRepos(githubPat string) ([]GithubRepo, error) {
	endpoint := "https://api.github.com/user/repos"
	httpClient := http.Client{}
	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		println(err.Error())
		return nil, err
	}
	request.Header.Add("Authorization", fmt.Sprintf("token %s", githubPat))
	resp, err := httpClient.Do(request)
	if err != nil {
		println(err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	githubRepos := []GithubRepo{}
	json.Unmarshal(body, &githubRepos)
	return githubRepos, nil
}

func (c *GithubHttpClient) GetActions(githubPat string, repoId int) (*GithubRepoActions, error) {
	endpoint := fmt.Sprintf("https://api.github.com/repositories/%d/actions/runs", repoId)
	httpClient := http.Client{}
	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		println(err.Error())
		return nil, err
	}
	request.Header.Add("Authorization", fmt.Sprintf("token %s", githubPat))
	resp, err := httpClient.Do(request)
	if err != nil {
		println(err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	githubRepoActions := &GithubRepoActions{}
	json.Unmarshal(body, &githubRepoActions)
	return githubRepoActions, nil
}
