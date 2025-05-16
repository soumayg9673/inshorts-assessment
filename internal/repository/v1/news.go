package v1rpo

import (
	newsv1 "github.com/soumayg9673/inshorts-assessment/internal/models/v1/news"
	"go.uber.org/zap"
)

func (rp *V1Rpo) GetNewsByCategory(q []string) ([]newsv1.NewsSql, error) {
	rows, err := rp.DB.GetNewsByCategory(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	news := []newsv1.NewsSql{}

	for rows.Next() {
		itr := newsv1.NewsSql{}
		if err := rows.Scan(
			&itr.Id,
			&itr.Title,
			&itr.Description,
			&itr.Url,
			&itr.PubDate,
			&itr.Source,
			&itr.Category,
			&itr.RevScore,
			&itr.LlmSummary,
			&itr.Latitude,
			&itr.Longitude,
		); err != nil {
			rp.LOG.Debug(err.Error(), zap.String("env", rp.ENV))
			return news, err
		}
		news = append(news, itr)
	}
	return news, nil
}

func (rp *V1Rpo) GetNewsByScore() ([]newsv1.NewsSql, error) {
	rows, err := rp.DB.GetNewsByScore()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	news := []newsv1.NewsSql{}

	for rows.Next() {
		itr := newsv1.NewsSql{}
		if err := rows.Scan(
			&itr.Id,
			&itr.Title,
			&itr.Description,
			&itr.Url,
			&itr.PubDate,
			&itr.Source,
			&itr.Category,
			&itr.RevScore,
			&itr.LlmSummary,
			&itr.Latitude,
			&itr.Longitude,
		); err != nil {
			rp.LOG.Debug(err.Error(), zap.String("env", rp.ENV))
			return news, err
		}
		news = append(news, itr)
	}
	return news, nil
}

func (rp *V1Rpo) GetNewsBySearch() {

}

func (rp *V1Rpo) GetNewsBySource(s int) ([]newsv1.NewsSql, error) {
	rows, err := rp.DB.GetNewsBySource(s)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	news := []newsv1.NewsSql{}

	for rows.Next() {
		itr := newsv1.NewsSql{}
		if err := rows.Scan(
			&itr.Id,
			&itr.Title,
			&itr.Description,
			&itr.Url,
			&itr.PubDate,
			&itr.Source,
			&itr.Category,
			&itr.RevScore,
			&itr.LlmSummary,
			&itr.Latitude,
			&itr.Longitude,
		); err != nil {
			rp.LOG.Debug(err.Error(), zap.String("env", rp.ENV))
			return news, err
		}
		news = append(news, itr)
	}
	return news, nil
}

func (rp *V1Rpo) GetNewsByNearby() {

}

func (rp *V1Rpo) PatchLlmSummary(id, ls string) {
	var t string
	row := rp.DB.PatchLlmSummary(id, ls)
	if err := row.Scan(&t); err != nil {
		rp.LOG.Debug(err.Error(),
			zap.String("env", rp.ENV),
		)
	}
}
