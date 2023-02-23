package repo

import (
	"time"

	"github.com/madnaaaaas/crud/pkg/refrigerator"
)

type beer struct {
	ID        int64     `sql:"id"`
	Title     string    `sql:"title"`
	ABV       string    `sql:"abv"`
	ExpiresAt time.Time `sql:"expires_at"`
}

func beerToDomain(p beer) *refrigerator.Beer {
	return &refrigerator.Beer{
		ID:        p.ID,
		Title:     p.Title,
		ABV:       p.ABV,
		ExpiresAt: p.ExpiresAt,
	}
}
