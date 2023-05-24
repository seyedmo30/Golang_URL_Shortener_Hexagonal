package domain

import (
	"context"
	"errors"
	"github.com/seyedmo30/Golang_URL_Shortener_Hexagonal/pkg/logger"
	"math/rand"
	"sync"
)

// Domain object
type urlShortner map[string]string

// singleton
var listUrl urlShortner

var once sync.Once

//Services

type ShortnerUrl struct {
}

func init() {

	once.Do(func() {

		listUrl = make(urlShortner)

	})

}

func generate(ctx context.Context) (string, error) {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 10)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b), nil
}

func (s *ShortnerUrl) AddUrl(ctx context.Context, url string) (string, error) {
	var shortGenerated string
	var ok bool
	var err error
	logger.Log().Info(url)
	for {

		shortGenerated, err = generate(ctx)
		logger.Log().Info(shortGenerated)
		if err != nil {
			logger.Log().Error(err.Error())
		}
		_, ok = listUrl[shortGenerated]

		if !ok {
			break
		}

	}
	listUrl[shortGenerated] = url

	return shortGenerated, nil
}

func (s *ShortnerUrl) GetUrl(ctx context.Context, shortner string) (string, error) {
	url, ok := listUrl[shortner]

	if ok {
		return url, nil
	}

	logger.Log().Error("not found")
	return "", errors.New("not found")
}
