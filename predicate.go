package legitscore

import (
	"fmt"
	"time"
)

type LegitScorePredicate struct {
	Repository string  `json:"repository"`
	Timestamp  string  `json:"timestamp"`
	Score      float64 `json:"score"`
}

const Schema = "https://raw.githubusercontent.com/Legit-Labs/legit-score/cb1dcd92f893d71dc32b1ed729023b803647532e/legit-score-predicate.json.schema"

func FetchScore(repo string, api_token string) *LegitScorePredicate {
	return &LegitScorePredicate{
		Repository: repo,
		Timestamp:  time.Now().Format(time.RFC1123),
		Score:      7, // TODO query Legit
	}
}

func (l *LegitScorePredicate) Verify(repo string, min_score float64) error {
	if l.Repository != repo {
		return fmt.Errorf("predicate repository (%v) does not match expected repository (%v)", l.Repository, repo)
	}

	if l.Score < min_score {
		return fmt.Errorf("predicate score (%v) is below the min score (%v)", l.Score, min_score)
	}

	return nil
}
