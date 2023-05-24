package ports

import "context"

type IHandlers interface {
	Add(ctx context.Context, url string) (string, error)
	Get(ctx context.Context, shortner string) (string, error)
}

type IshortnerUrl interface {
	AddUrl(ctx context.Context, url string) (string, error)
	GetUrl(ctx context.Context, shortner string) (string, error)
}
