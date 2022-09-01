package legitscore

type LegitScorePredicate struct {
	Repository string  `json:"repository"`
	Timestamp  string  `json:"timestamp"`
	Score      float64 `json:"score"`
}
