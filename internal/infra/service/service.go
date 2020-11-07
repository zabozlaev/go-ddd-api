package service

import (
	"go-ddd-api/internal/domain"
	"go-ddd-api/internal/infra/store"

	"github.com/sirupsen/logrus"
)

type srv struct {
	logger     *logrus.Logger
	store      store.Store
	urlService domain.URLService
}

// Service - manager for all services
type Service interface {
	URLService() domain.URLService
}

// NewService - service factory
func NewService(logger *logrus.Logger, store store.Store) Service {
	return &srv{logger: logger, store: store}
}

func (s *srv) URLService() domain.URLService {
	if s.urlService == nil {
		s.urlService = newURLService(s.logger, s.store)
	}

	return s.urlService
}
