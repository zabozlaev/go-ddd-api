package service

import (
	"go-ddd-api/internal/domain"
	"go-ddd-api/internal/infra/store"
	"go-ddd-api/internal/infra/store/repository"
	"go-ddd-api/pkg/httperr"
	"net/http"

	"github.com/dchest/uniuri"
	"github.com/sirupsen/logrus"
)

type urlService struct {
	store  store.Store
	logger *logrus.Logger
}

func newURLService(logger *logrus.Logger, store store.Store) domain.URLService {
	return &urlService{store, logger}
}

func (us *urlService) Create(d *domain.CreateURL) (*domain.URL, error) {
	s := uniuri.NewLen(7)

	d.Short = s

	return us.store.URLRepo().Create(d)
}

func (us *urlService) FindOrigin(s string) (string, error) {
	found, err := us.store.URLRepo().FindByShort(s)

	if err != nil {
		switch err {
		case repository.ErrURLNotFound:
			return "", httperr.NewHttpError(err, http.StatusNotFound)
		}
		return "", err
	}

	err = us.store.URLRepo().Hit(found.ID)

	if err != nil {
		return "", err
	}

	return found.Origin, nil
}
