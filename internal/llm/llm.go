package llm

import (
	"context"

	"go.uber.org/zap"
)

type Llm struct {
	GeminiAi
}

func NewLlmStore(log *zap.Logger, ctx context.Context) Llm {
	return Llm{
		GeminiAi: newGeminiAi(log, ctx),
	}
}
