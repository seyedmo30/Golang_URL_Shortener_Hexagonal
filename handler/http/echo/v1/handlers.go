package v1

import (
	"context"
	"github.com/seyedmo30/Golang_URL_Shortener_Hexagonal/core/domain"
	"github.com/seyedmo30/Golang_URL_Shortener_Hexagonal/pkg/logger"
)

type concritPorts struct {
}

func (c *concritPorts) Add(ctx context.Context, url string) (string, error) {
	shortnerUrl := new(domain.ShortnerUrl)
	res, err := shortnerUrl.AddUrl(ctx, url)
	if err != nil {
		logger.Log().Error("cont add to list")
	}
	return res, nil
}
func (c *concritPorts) Get(ctx context.Context, shortner string) (string, error) {
	shortnerUrl := new(domain.ShortnerUrl)
	res, err := shortnerUrl.GetUrl(ctx, shortner)
	if err != nil {
		logger.Log().Error("not found")
		return "", err
	}
	return res, nil

}
