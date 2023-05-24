package domain

import (
	"context"
	"github.com/seyedmo30/Golang_URL_Shortener_Hexagonal/pkg/logger"
	"testing"
)

func TestGenerator(t *testing.T) {
	ctx := context.Background()
	logs := logger.Log()
	randShortner, err := generate(ctx)
	if err != nil {
		logs.Error(err.Error())
	}
	logs.Info(randShortner)
}
