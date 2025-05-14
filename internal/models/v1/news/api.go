package newsv1

type NewsApi struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Url         string  `json:"url"`
	PubDate     string  `json:"publication_date"`
	Source      string  `json:"source_name"`
	Category    string  `json:"category"`
	RevScore    float32 `json:"relevance_score"`
	LlmSummary  string  `json:"llm_summary"`
	Latitude    float32 `json:"latitude"`
	Longitude   float32 `json:"longitude"`
}
