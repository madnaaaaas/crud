package refrigerator

import "context"

type Service interface {
	GetBeerByTitle(ctx context.Context, title string) (*Beer, error)
	InsertBeer(ctx context.Context, beer *Beer) (int64, error)
}
