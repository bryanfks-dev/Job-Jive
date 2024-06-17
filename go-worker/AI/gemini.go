package ai

import (
	"context"
	"log"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var (
	Ctx    context.Context
	Client *genai.Client
	Model  *genai.GenerativeModel
	err    error
)

func CreateClient(api_key string) {
	Ctx := context.Background()

	// Create new gemini client
	Client, err = genai.NewClient(Ctx, option.WithAPIKey(api_key))

	// Ensure no error create client
	if err != nil {
		panic(err.Error())
	}

	Model = Client.GenerativeModel("gemini-pro")
}

func Generate(promt string) /* (string, error) */ {
	res, err := Model.GenerateContent(Ctx, genai.Text(promt))

	if err != nil {
		/* return "", err */
	}

	if res != nil {
		candidates := res.Candidates

		if candidates != nil {
			for iter, candidate := range candidates {
				if iter > 0 {
					content := candidate.Content

					if content != nil {
						text := content.Parts[0]

						log.Println(text)
					}
				}
			}
		}
	}
}
