package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/Legit-Labs/legit-score/pkg/legit_score"
)

var (
	repository string
	apiToken   string
)

func main() {
	flag.StringVar(&repository, "repository", "", "The repository for which to fetch the score")
	flag.StringVar(&apiToken, "api-token", "", "The API token to use to fetch the score")

	flag.Parse()

	if repository == "" {
		log.Panicf("please provide a repository")
	} else if apiToken == "" {
		log.Panicf("please provide an API token")
	}

	predicate, err := legit_score.FetchScore(repository, apiToken)
	if err != nil {
		log.Panicf("failed to fetch score: %v", err)
	}

	jsonized, err := json.Marshal(predicate)

	fmt.Printf("%v", string(jsonized))
}
