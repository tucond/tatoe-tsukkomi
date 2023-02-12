package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatalln("Missing API KEY")
	}

	ctx := context.Background()
	client := gpt3.NewClient(apiKey)
	defaultPrompt :=
		`
暑い=砂漠か！
暗い=田舎道か！
ケチ=給料日前か！
汚い=公園の端に落ちてるやつか！
長い=万里の長城か！
少食=歯医者の帰りか！
`

	appendPrompt := "かわいい"
	wholePrompt := defaultPrompt + appendPrompt + "="

	resp, err := client.CompletionWithEngine(ctx, gpt3.TextDavinci003Engine, gpt3.CompletionRequest{
		// Prompt:    []string{"The first thing you should know about javascript is"},
		Prompt:    []string{wholePrompt},
		MaxTokens: gpt3.IntPtr(30),
		Stop:      []string{"."},
		Echo:      true,
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp.Choices[0].Text)
}
