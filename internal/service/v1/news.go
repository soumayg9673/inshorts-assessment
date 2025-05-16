package v1svc

import (
	"github.com/soumayg9673/inshorts-assessment/internal/llm/gemini"
	newsv1 "github.com/soumayg9673/inshorts-assessment/internal/models/v1/news"
)

func (sc *V1Svc) GetNewsByCategory(q []string) ([]newsv1.NewsApi, error) {
	//TODO: check cache if exists with category id and timestamp
	news, err := sc.RPO.GetNewsByCategory(q)
	if err != nil {
		return nil, err
	}

	data := []newsv1.NewsApi{}
	for _, n := range news {
		llmSummary := n.LlmSummary.String
		if !n.LlmSummary.Valid {
			llmSummary = sc.Llm.GeminiAi.QueryText(gemini.GeminiMdl.Gemini_2_0_Flash, n.Url)
			go sc.RPO.PatchLlmSummary(n.Id, llmSummary)
		}
		data = append(data, newsv1.NewsApi{
			Title:       n.Title,
			Description: n.Description,
			Url:         n.Url,
			PubDate:     n.PubDate,
			Source:      n.Source,
			Category:    n.Category,
			RevScore:    n.RevScore,
			LlmSummary:  llmSummary,
			Latitude:    n.Latitude,
			Longitude:   n.Longitude,
		})
	}
	return data, nil
}

func (sc *V1Svc) GetNewsByScore() ([]newsv1.NewsApi, error) {
	//TODO: check cache if exists with category id and timestamp
	news, err := sc.RPO.GetNewsByScore()
	if err != nil {
		return nil, err
	}

	data := []newsv1.NewsApi{}
	for _, n := range news {
		llmSummary := n.LlmSummary.String
		if !n.LlmSummary.Valid {
			llmSummary = sc.Llm.GeminiAi.QueryText(gemini.GeminiMdl.Gemini_2_0_Flash, n.Url)
			go sc.RPO.PatchLlmSummary(n.Id, llmSummary)
		}
		data = append(data, newsv1.NewsApi{
			Title:       n.Title,
			Description: n.Description,
			Url:         n.Url,
			PubDate:     n.PubDate,
			Source:      n.Source,
			Category:    n.Category,
			RevScore:    n.RevScore,
			LlmSummary:  llmSummary,
			Latitude:    n.Latitude,
			Longitude:   n.Longitude,
		})
	}
	return data, nil
}

func (sc *V1Svc) GetNewsBySearch() {

}

func (sc *V1Svc) GetNewsBySource(s int) ([]newsv1.NewsApi, error) {
	//TODO: check cache if exists with category id and timestamp
	news, err := sc.RPO.GetNewsBySource(s)
	if err != nil {
		return nil, err
	}

	data := []newsv1.NewsApi{}
	for _, n := range news {
		llmSummary := n.LlmSummary.String
		if !n.LlmSummary.Valid {
			llmSummary = sc.Llm.GeminiAi.QueryText(gemini.GeminiMdl.Gemini_2_0_Flash, n.Url)
			go sc.RPO.PatchLlmSummary(n.Id, llmSummary)
		}
		data = append(data, newsv1.NewsApi{
			Title:       n.Title,
			Description: n.Description,
			Url:         n.Url,
			PubDate:     n.PubDate,
			Source:      n.Source,
			Category:    n.Category,
			RevScore:    n.RevScore,
			LlmSummary:  llmSummary,
			Latitude:    n.Latitude,
			Longitude:   n.Longitude,
		})
	}
	return data, nil
}

func (sc *V1Svc) GetNewsByNearby() {

}
