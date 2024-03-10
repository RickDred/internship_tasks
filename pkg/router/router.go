package router

import (
	"net/http"
)

func NewRouter() *Router {
	return &Router{
		NotFoundHandler:  http.NotFound,
		MethodNotAllowed: DefaultMethodNotAllowedHandler,
		corsEnabled:      false,
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if r.corsEnabled {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, authorization")

		if req.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	handler := r.findHandler(req)

	for _, middleware := range r.middlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, req)
}
