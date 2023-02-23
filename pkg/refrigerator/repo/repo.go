package repo

import (
	"context"
	"github.com/go-pg/pg"

	"github.com/madnaaaaas/crud/pkg/refrigerator"
)

type repo struct {
	db *pg.DB
}

func NewRepo(db *pg.DB) *repo {
	return &repo{db}
}

func (r *repo) GetBeerByTitle(ctx context.Context, title string) (*refrigerator.Beer, error) {
	var res beer
	_, err := r.db.QueryContext(ctx, &res,
		`SELECT * FROM beer
			   WHERE title = ?`,
		title)

	if err != nil {
		return nil, err
	}

	return beerToDomain(res), nil
}

func (r *repo) InsertBeer(ctx context.Context, b *refrigerator.Beer) (int64, error) {
	var res int64
	_, err := r.db.QueryContext(ctx, &res,
		`INSERT INTO beer (title, abv, expires_at)
			   VALUES (?, ?, ?)
			   RETURNING id`,
		b.Title, b.ABV, b.ExpiresAt)

	if err != nil {
		return 0, err
	}

	return res, nil
}
