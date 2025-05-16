package v1db

import (
	"database/sql"
	"strings"

	"go.uber.org/zap"
)

func (ds *V1Db) GetNewsByCategory(q []string) (*sql.Rows, error) {
	const query = `
	SELECT 
		na.id,
		na.title,
		na.description,
		na.url,
		na.publication_date,
		ns.source_name,
		nc.category_name,
		na.relevance_score,
		na.llm_summary,
		ST_Y(na."location"::geometry) AS latitude,
		ST_X(na."location"::geometry) AS longitude
	FROM news_article_categories nac 
	JOIN news_categories nc ON nc.id  = nac.category_id AND nc.id IN ($1)
	JOIN news_articles na ON na.id = nac.article_id
	JOIN news_sources ns ON ns.id = na.source_id
	ORDER BY na.publication_date DESC
	LIMIT 5;
	`

	rows, err := ds.DB.Query(query, strings.Join(q, ","))
	if err != nil {
		ds.LOG.Debug(err.Error(),
			zap.String("env", ds.ENV),
		)
		return nil, err
	}

	return rows, nil
}

func (ds *V1Db) GetNewsByScore() (*sql.Rows, error) {
	const query = `
	SELECT
		na.id,
		na.title,
		na.description,
		na.url,
		na.publication_date,
		ns.source_name,
		STRING_AGG(nc.category_name, ', ') AS categories,
		na.relevance_score,
		na.llm_summary,
		ST_Y(na.location::geometry) AS latitude,
		ST_X(na.location::geometry) AS longitude
	FROM news_articles na
	JOIN news_sources ns ON ns.id = na.source_id
	JOIN news_article_categories nac ON nac.article_id = na.id
	JOIN news_categories nc ON nc.id = nac.category_id
	WHERE na.relevance_score > 0.7
	GROUP BY na.id, ns.source_name
	ORDER BY na.publication_date DESC, na.relevance_score desc
	LIMIT 5;
	`

	rows, err := ds.DB.Query(query)
	if err != nil {
		ds.LOG.Debug(err.Error(),
			zap.String("env", ds.ENV),
		)
		return nil, err
	}

	return rows, nil
}

func (ds *V1Db) GetNewsBySearch() {

}

func (ds *V1Db) GetNewsBySource(s int) (*sql.Rows, error) {
	const query = `
	SELECT
		na.id,
		na.title,
		na.description,
		na.url,
		na.publication_date,
		ns.source_name,
		STRING_AGG(nc.category_name, ', ') AS categories,
		na.relevance_score,
		na.llm_summary,
		ST_Y(na.location::geometry) AS latitude,
		ST_X(na.location::geometry) AS longitude
	FROM news_articles na
	JOIN news_sources ns ON ns.id = na.source_id
	JOIN news_article_categories nac ON nac.article_id = na.id
	JOIN news_categories nc ON nc.id = nac.category_id
	WHERE na.source_id = $1
	GROUP BY na.id, ns.source_name
	ORDER BY na.publication_date DESC
	LIMIT 5;
	`

	rows, err := ds.DB.Query(query, s)
	if err != nil {
		ds.LOG.Debug(err.Error(),
			zap.String("env", ds.ENV),
		)
		return nil, err
	}

	return rows, nil
}

func (ds *V1Db) GetNewsByNearby() {

}

func (ds *V1Db) PatchLlmSummary(id, ls string) *sql.Row {
	const query = `
	UPDATE news_articles
	SET llm_summary = $1
	WHERE id = $2
	RETURNING 1
	`
	rows := ds.DB.QueryRow(query, ls, id)
	return rows
}
