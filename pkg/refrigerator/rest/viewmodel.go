package rest

import (
	"time"

	"github.com/madnaaaaas/crud/pkg/refrigerator"
)

type postBeerVM struct {
	Title     string `json:"title"`
	ABV       string `json:"abv"`
	ExpiresAt string `json:"expiresAt"`
}

type beerVM struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	ABV       string `json:"abv"`
	ExpiresAt string `json:"expiresAt"`
}

func postBeerVMToDomain(vm *postBeerVM) (*refrigerator.Beer, error) {
	expiresAt, err := time.Parse(time.RFC3339, vm.ExpiresAt)
	if err != nil {
		return nil, err
	}
	return &refrigerator.Beer{
		Title:     vm.Title,
		ABV:       vm.ABV,
		ExpiresAt: expiresAt,
	}, nil
}

func beerDomainToVM(d *refrigerator.Beer) beerVM {
	return beerVM{
		ID:        d.ID,
		Title:     d.Title,
		ABV:       d.ABV,
		ExpiresAt: d.ExpiresAt.Format(time.RFC3339),
	}
}
