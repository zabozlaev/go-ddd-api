package web

import (
	"net/http"
)

func configureRoutes(r *router, h Handlers) {
	r.Handle(http.MethodPost, "/urls", h.URLHandlers().CreateURL())
	r.Handle(http.MethodGet, "/{short}", h.URLHandlers().GetRedirectURL())
}
