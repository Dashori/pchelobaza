package models

type Review struct {
	ReviewId  uint64
	Owner     string
	Timestamp string
	Text      string
}

type AllReviews struct {
	Reviews []Review
}
