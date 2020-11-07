package web

import (
	"go-ddd-api/internal/domain"
	"go-ddd-api/internal/infra/service"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/sirupsen/logrus"
)

// URLHandler - /api/urls
type URLHandler interface {
	CreateURL() Handler
	GetRedirectURL() Handler
}

type urlHandler struct {
	logger  *logrus.Logger
	service service.Service
}

// NewURLHandler - handler factory
func NewURLHandler(logger *logrus.Logger, service service.Service) URLHandler {
	return &urlHandler{logger, service}
}

func (uh *urlHandler) CreateURL() Handler {
	return Handler(func(w http.ResponseWriter, r *http.Request) error {
		var data domain.CreateURL

		err := DecodeJSONBody(r, &data)

		if err != nil {
			return err
		}

		data.IP = r.RemoteAddr
		data.UserAgent = r.Header.Get("User-Agent")
		data.Hits = 0

		created, err := uh.service.URLService().Create(&data)

		if err != nil {
			return err
		}

		return RespondJSON(w, r, created, http.StatusCreated)
	})
}

func (uh *urlHandler) GetRedirectURL() Handler {
	return Handler(func(w http.ResponseWriter, r *http.Request) error {
		short := chi.URLParam(r, "short")

		url, err := uh.service.URLService().FindOrigin(short)

		if err != nil {
			return err
		}

		http.Redirect(w, r, url, http.StatusSeeOther)
		return nil
	})
}
