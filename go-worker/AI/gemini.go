package ai

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var (
	ctx          = context.Background()
	GeminiClient *genai.Client
	Model        *genai.GenerativeModel
	err          error

	ErrNoGeminiResponse = errors.New("no response from gemini")
)

func InitGeminiClient(api_key string) error {
	// Create new gemini client
	GeminiClient, err = genai.NewClient(ctx, option.WithAPIKey(api_key))

	// Ensure no error create client
	if err != nil {
		return err
	}

	Model = GeminiClient.GenerativeModel("gemini-1.5-pro")

	return nil
}

func GeminiGenerate(promt string) (string, error) {
	res, err := Model.GenerateContent(ctx, genai.Text(promt))

	if err != nil {
		return "", err
	}

	if res != nil {
		candidates := res.Candidates

		if candidates != nil {
			content := candidates[0].Content

			if content != nil {
				return fmt.Sprintf("%v", content.Parts), nil
			}
		}
	}

	return "", ErrNoGeminiResponse
}
