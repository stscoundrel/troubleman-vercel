package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/stscoundrel/troubleman/issues"
)

type Response struct {
	ReposCount   int
	IssuesCount  int
	Repositories []issues.Repository
}

func Handler(w http.ResponseWriter, r *http.Request) {
	repositoriesResult, repositoriesError := issues.GetRepositories("stscoundrel")

	if repositoriesError == nil {
		reposCount := len(repositoriesResult)
		issuesCount := 0

		for _, repo := range repositoriesResult {
			issuesCount = issuesCount + repo.Count
		}

		response := Response{reposCount, issuesCount, repositoriesResult}

		jsonResult, _ := json.Marshal(response)

		fmt.Fprintf(w, string(jsonResult))
	} else {
		fmt.Fprintf(w, "Could not fetch repositories for user stscoundrel")
	}
}
