package v1db

import (
	"database/sql"

	"go.uber.org/zap"
)

func (ds *V1Db) GetNewsByCategory(q []string) (*sql.Rows, error) {
	const query = `
	SELECT 
		na.id,
		na.title,
		na.description,
		na.publication_date,
		na.relevance_score,
		ns.source_name,
		nc.category_name,
		ST_X(na."location"::geometry) AS longitude,
		ST_Y(na."location"::geometry) AS latitude
	FROM news_article_categories nac 
	JOIN news_categories nc ON nc.id  = nac.category_id AND nc.id IN ($1)
	JOIN news_articles na ON na.id = nac.article_id
	JOIN news_sources ns ON ns.id = na.source_id
	ORDER BY na.publication_date DESC
	LIMIT 5;
	`

	rows, err := ds.DB.Query(query, q)
	if err != nil {
		ds.LOG.Debug(err.Error(),
			zap.String("env", ds.ENV),
		)
		return nil, err
	}

	return rows, nil
}

func (ds *V1Db) GetNewsByScore() {

}

func (ds *V1Db) GetNewsBySearch() {

}

func (ds *V1Db) GetNewsBySource() {

}

func (ds *V1Db) GetNewsByNearby() {

}
