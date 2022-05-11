package main

import (
	"awesomeProject/lex"
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"log"
)

const profile = "profile"

func main() {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithSharedConfigProfile(profile),
		//config.WithRegion("ap-northeast-1"),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	c := lex.NewLexClient(cfg)
	text := "ホテルを予約する"
	output, err := c.RecognizeText(ctx, &text)
	for _, m := range output.Messages {
		log.Printf("Message: %s\n", *m.Content)
	}
}
