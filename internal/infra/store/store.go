package store

import (
	"go-ddd-api/internal/domain"
	"go-ddd-api/internal/infra/store/repository"

	"github.com/jmoiron/sqlx"
)

type storeAdapter struct {
	db      *sqlx.DB
	urlRepo domain.URLRepository
}

// Store - repository registry
type Store interface {
	URLRepo() domain.URLRepository
}

// NewStore - store factory
func NewStore(db *sqlx.DB) Store {
	return &storeAdapter{db: db}
}

func (a *storeAdapter) URLRepo() domain.URLRepository {
	if a.urlRepo == nil {
		a.urlRepo = repository.NewURLRepo(a.db)
	}

	return a.urlRepo
}
