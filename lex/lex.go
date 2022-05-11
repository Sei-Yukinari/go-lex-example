package lex

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lexruntimev2"
)

const (
	botAliasId = "botAliasID"
	botId      = "botID"
	localeId   = "ja_JP"
	sessionId  = "123456789"
)

type ILex interface {
	RecognizeText(ctx context.Context, text *string) (*lexruntimev2.RecognizeTextOutput, error)
}

type Lex struct {
	client *lexruntimev2.Client
}

func NewLexClient(cfg aws.Config) ILex {
	c := lexruntimev2.NewFromConfig(cfg)
	return &Lex{client: c}
}

func (l *Lex) RecognizeText(ctx context.Context, text *string) (*lexruntimev2.RecognizeTextOutput, error) {
	return l.client.RecognizeText(ctx, l.builderRecognizeTextInput(text))
}

func (l *Lex) builderRecognizeTextInput(text *string) *lexruntimev2.RecognizeTextInput {

	return &lexruntimev2.RecognizeTextInput{
		BotAliasId:        aws.String(botAliasId),
		BotId:             aws.String(botId),
		LocaleId:          aws.String(localeId),
		SessionId:         aws.String(sessionId),
		Text:              text,
		RequestAttributes: nil,
		SessionState:      nil,
	}
}
