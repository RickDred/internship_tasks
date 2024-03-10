package router

import "net/http"

type Route struct {
	Method      string
	Pattern     string
	Handler     http.HandlerFunc
	Middlewares []Middleware
}

type Router struct {
	routes           []Route
	middlewares      []Middleware
	NotFoundHandler  http.HandlerFunc
	MethodNotAllowed http.HandlerFunc
	corsEnabled      bool
}

type Middleware func(http.HandlerFunc) http.HandlerFunc
