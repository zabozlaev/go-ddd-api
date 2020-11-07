package web

import (
	"go-ddd-api/internal/infra/service"

	"github.com/sirupsen/logrus"
)

// Handlers is a manager interface
type Handlers interface {
	URLHandlers() URLHandler
}

// Handlers - collection of handlers
type handlersManager struct {
	URLHandler URLHandler
}

// NewHandlers is a factory for handlers manager
func NewHandlers(logger *logrus.Logger, service service.Service) Handlers {
	return &handlersManager{
		NewURLHandler(logger, service),
	}
}

func (h *handlersManager) URLHandlers() URLHandler {
	return h.URLHandler
}
