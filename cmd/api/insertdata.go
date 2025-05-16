package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"strings"
	"time"

	"github.com/soumayg9673/inshorts-assessment/internal/llm/gemini"
)

type newsItem struct {
	ID              string   `json:"id"`
	Title           string   `json:"title"`
	Description     string   `json:"description"`
	URL             string   `json:"url"`
	PublicationDate string   `json:"publication_date"`
	SourceName      string   `json:"source_name"`
	Category        []string `json:"category"`
	RelevanceScore  float64  `json:"relevance_score"`
	Latitude        float64  `json:"latitude"`
	Longitude       float64  `json:"longitude"`
}

func slugify(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, " ", "_"))
}

func (app *application) insertInitialData(db *sql.DB, ctx context.Context) {
	// Read JSON
	file, err := os.ReadFile("scripts/data/news_data.json")
	if err != nil {
		log.Fatal("Cannot read news_data.json:", err)
	}

	var items []newsItem
	if err := json.Unmarshal(file, &items); err != nil {
		log.Fatal("Failed to parse JSON:", err)
	}

	sourceCache := make(map[string]int)
	categoryCache := make(map[string]int)

	for _, item := range items {
		// Insert/Get Source
		sourceID, ok := sourceCache[item.SourceName]
		if !ok {
			err = db.QueryRow(`
				INSERT INTO news_sources (source_name)
				VALUES ($1)
				RETURNING id;
			`, item.SourceName).Scan(&sourceID)

			if err == sql.ErrNoRows || sourceID == 0 {
				err = db.QueryRow(`SELECT id FROM news_sources WHERE source_name = $1;`, item.SourceName).Scan(&sourceID)
			}
			if err != nil {
				log.Fatalf("Source insert/fetch error: %v", err)
			}
			sourceCache[item.SourceName] = sourceID
		}

		parsedTime, err := time.Parse("2006-01-02T15:04:05", item.PublicationDate)
		if err != nil {
			log.Fatalf("Failed to parse publication date for ID %s: %v", item.ID, err)
		}

		// Insert Article
		_, err = db.Exec(`
			INSERT INTO news_articles (
				id, title, description, url, publication_date,
				source_id, relevance_score, location
			)
			VALUES (
				$1, $2, $3, $4, $5,
				$6, $7,
				ST_SetSRID(ST_MakePoint($8, $9), 4326)::geography
			);
		`, item.ID, item.Title, item.Description, item.URL, parsedTime,
			sourceID, item.RelevanceScore, item.Longitude, item.Latitude)

		if err != nil {
			log.Fatalf("Article insert error: %v", err)
		}

		go app.llm.GeminiAi.QueryText(gemini.GeminiMdl.Gemini_2_0_Flash, item.URL)

		// Insert Categories and Map
		for _, cat := range item.Category {
			catID, ok := categoryCache[cat]
			if !ok {
				catIdentifier := slugify(cat)
				err = db.QueryRow(`
					INSERT INTO news_categories (category_name, category_identifier)
					VALUES ($1, $2)
					ON CONFLICT (category_identifier) DO NOTHING
					RETURNING id;
				`, cat, catIdentifier).Scan(&catID)

				if err == sql.ErrNoRows || catID == 0 {
					err = db.QueryRow(`SELECT id FROM news_categories WHERE category_identifier = $1;`, catIdentifier).Scan(&catID)
				}
				if err != nil {
					log.Fatalf("Category insert/fetch error: %v", err)
				}
				categoryCache[cat] = catID
			}

			// Insert Mapping
			_, err = db.Exec(`
				INSERT INTO news_article_categories (article_id, category_id)
				VALUES ($1, $2)
				ON CONFLICT DO NOTHING;
			`, item.ID, catID)

			if err != nil {
				log.Fatalf("Mapping insert error: %v", err)
			}
		}
	}

}
