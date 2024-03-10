package router

import (
	"net/http"
)

func (r *Router) SetNotFoundHandler(f http.HandlerFunc) {
	r.NotFoundHandler = f
}

func (r *Router) SetMethodNotAllowedHandler(f http.HandlerFunc) {
	r.MethodNotAllowed = f
}
