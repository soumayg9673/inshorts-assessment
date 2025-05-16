package newsv1

import "database/sql"

type NewsSql struct {
	Id          string
	Title       string
	Description string
	Url         string
	PubDate     string
	Source      string
	Category    string
	RevScore    float32
	LlmSummary  sql.NullString
	Latitude    float32
	Longitude   float32
}
