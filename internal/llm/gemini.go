package llm

import (
	"context"
	"fmt"
	"os"

	"go.uber.org/zap"
	"google.golang.org/genai"
)

type GeminiAi struct {
	Ctx    context.Context
	Client *genai.Client
}

func newGeminiAi(log *zap.Logger) GeminiAi {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Debug(err.Error())
	}

	return GeminiAi{
		Ctx:    ctx,
		Client: client,
	}
}

func (g GeminiAi) QueryText(gm string, url string) {
	config := &genai.GenerateContentConfig{
		SystemInstruction: genai.NewContentFromText("You are a user interested in news article summary without clicking on the link", genai.RoleUser),
	}

	result, _ := g.Client.Models.GenerateContent(
		g.Ctx,
		gm,
		genai.Text(url),
		config,
	)

	fmt.Println(result.Text())
}
