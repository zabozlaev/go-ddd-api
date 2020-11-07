package web

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

// Handler - request handler
type Handler func(w http.ResponseWriter, r *http.Request) error

// Middleware - executed before
type Middleware func(Handler) Handler

type router struct {
	chiRouter chi.Router
}

func newRouter(r chi.Router) *router {
	return &router{
		r,
	}
}

func wrapMiddleware(handler Handler, mw []Middleware) Handler {

	for i := len(mw) - 1; i > 0; i-- {
		if mw[i] != nil {
			handler = mw[i](handler)
		}
	}

	return handler
}

func (r *router) Handle(method, pattern string, handler Handler, mw ...Middleware) {
	handler = wrapMiddleware(handler, mw)

	fn := func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err != nil {
			err = RespondError(w, r, err)
			if err != nil {
				fmt.Printf("err during response handling %s", err.Error())
			}
		}
	}

	r.chiRouter.MethodFunc(method, pattern, fn)
}
