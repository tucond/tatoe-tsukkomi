package main

import (
	"context"
	"log"
	"os"
	"unicode/utf8"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
)

// Promptに入力する文字列の最大長
const MaxInputLength = 20

// 文字数超過時に返すメッセージ
const VerifyMessage string = "20文字以内で入力してください"

func getGptResponse(inputString string) *string {

	godotenv.Load()

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatalln("Missing API KEY")
	}

	// 文字数ベリファイ
	verifyMessageForPtr := VerifyMessage //https://onl.tw/ZqBmnZc
	if utf8.RuneCountInString(inputString) > MaxInputLength {
		return &verifyMessageForPtr
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

	Prompt := defaultPrompt + inputString + "="

	resp, err := client.CompletionWithEngine(ctx, gpt3.TextDavinci003Engine, gpt3.CompletionRequest{
		Prompt:    []string{Prompt},
		MaxTokens: gpt3.IntPtr(30),
		Stop:      []string{"."},
		Echo:      false,
	})
	if err != nil {
		log.Fatalln(err)
	}
	return &resp.Choices[0].Text
}
