package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/stscoundrel/troubleman/issues"
)

type Response struct {
	Count  int
	Issues []issues.Issue
}

func Handler(w http.ResponseWriter, r *http.Request) {
	issuesResult, issuesError := issues.GetIssues("stscoundrel")

	if issuesError == nil {
		count := len(issuesResult)
		response := Response{count, issuesResult}
		jsonResult, _ := json.Marshal(response)

		fmt.Fprintf(w, string(jsonResult))
	} else {
		fmt.Fprintf(w, "Could not fetch issues for stscoundrel")
	}
}
