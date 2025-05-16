package llm

import "go.uber.org/zap"

type Llm struct {
	GeminiAi
}

func NewLlmStore(log *zap.Logger) Llm {
	return Llm{
		GeminiAi: newGeminiAi(log),
	}
}
