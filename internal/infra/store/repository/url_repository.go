package repository

import (
	"database/sql"
	"go-ddd-api/internal/domain"

	"github.com/pkg/errors"

	"github.com/jmoiron/sqlx"
)

type urlRepo struct {
	db *sqlx.DB
}

var (
	// ErrURLNotFound - not found err
	ErrURLNotFound = errors.New("url not found")
)

// NewURLRepo - repo factory
func NewURLRepo(db *sqlx.DB) domain.URLRepository {
	return &urlRepo{db}
}

func (ur *urlRepo) FindByShort(short string) (*domain.URL, error) {
	var found domain.URL

	err := ur.db.Get(&found, "SELECT * FROM urls WHERE short = $1;", short)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrURLNotFound
		}
		return nil, errors.Wrap(err, "selecting err")
	}

	return &found, nil
}

func (ur *urlRepo) Create(d *domain.CreateURL) (*domain.URL, error) {
	r, err := ur.db.NamedQuery("INSERT INTO urls (short, origin, agent, ip, hits) VALUES (:short, :origin, :agent, :ip, :hits) RETURNING *;", d.MapToModel())

	if err != nil {
		return nil, err
	}

	var created domain.URL

	if r.Next() {
		err = r.StructScan(&created)

		if err != nil {
			return nil, err
		}
	}

	return &created, nil
}

func (ur *urlRepo) Hit(id int) error {

	_, err := ur.db.Exec("UPDATE urls SET hits = hits + 1 WHERE id = $1;", id)

	if err != nil {
		return err
	}

	return nil
}
