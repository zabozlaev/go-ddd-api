package web

import (
	"context"
	"go-ddd-api/internal/app"
	"go-ddd-api/internal/config"
	"go-ddd-api/internal/infra/service"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/sirupsen/logrus"
)

type apiAdapter struct {
	config *config.Config
	server *http.Server
}

// NewAPI returns an api adapter
func NewAPI(c *config.Config, logger *logrus.Logger, service service.Service) app.API {

	r := newRouter(chi.NewRouter())
	hs := NewHandlers(logger, service)

	configureRoutes(r, hs)

	a := &apiAdapter{
		server: &http.Server{
			Addr:         c.Port,
			ReadTimeout:  c.ShutdownTimeout,
			WriteTimeout: c.ShutdownTimeout,
			Handler:      r.chiRouter,
		},
	}

	return a
}

func (a *apiAdapter) Shutdown(ctx context.Context) error {
	return a.server.Shutdown(ctx)
}

func (a *apiAdapter) Close() error {
	return a.server.Close()
}

func (a *apiAdapter) Run() error {
	return a.server.ListenAndServe()
}
