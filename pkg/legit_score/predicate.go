package legit_score

import (
	"fmt"
	"time"

	intoto "github.com/in-toto/in-toto-golang/in_toto"
)

type LegitScorePredicate struct {
	Repository string  `json:"repository"`
	Timestamp  string  `json:"timestamp"`
	Score      float64 `json:"score"`
}

type LegitScoreStatement struct {
	intoto.StatementHeader
	Predicate LegitScorePredicate
}

const Schema = "https://raw.githubusercontent.com/Legit-Labs/legit-score/07bf6fa16d2b2543d3280ad0ee5879cfed16a43a/pub/legit-score-predicate.json.schema"

func NewLegitScorePredicate(repo string, score float64) *LegitScorePredicate {
	return &LegitScorePredicate{
		Repository: repo,
		Timestamp:  time.Now().Format(time.RFC1123),
		Score:      score,
	}
}

func fetchScoreForRepo(repo string, apiToken string) (float64, error) {
	return 7, nil // TODO query Legit
}

func FetchScore(repo string, apiToken string) (*LegitScorePredicate, error) {
	score, err := fetchScoreForRepo(repo, apiToken)
	if err != nil {
		return nil, err
	}

	predicate := NewLegitScorePredicate(repo, score)

	return predicate, nil
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
