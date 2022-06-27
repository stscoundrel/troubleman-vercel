package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/stscoundrel/troubleman/issues"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	repositoriesResult, repositoriesError := issues.GetRepositories("stscoundrel")

	if repositoriesError == nil {
		jsonResult, _ := json.Marshal(repositoriesResult)

		fmt.Fprintf(w, string(jsonResult))
	} else {
		fmt.Fprintf(w, "Could not fetch repositories for user stscoundrel")
	}
}
