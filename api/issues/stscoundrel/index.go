package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/stscoundrel/troubleman/issues"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	issuesResult, issuesError := issues.GetIssues("stscoundrel")

	if issuesError == nil {
		jsonResult, _ := json.Marshal(issuesResult)

		fmt.Fprintf(w, string(jsonResult))
	} else {
		fmt.Fprintf(w, "Could not fetch issues for stscoundrel")
	}
}
