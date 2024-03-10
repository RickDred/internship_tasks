package router

import "net/http"

func (r *Router) GET(path string, handler http.HandlerFunc, middlewares ...Middleware) {
	r.AddRoute("GET", path, handler, middlewares...)
}

func (r *Router) POST(path string, handler http.HandlerFunc, middlewares ...Middleware) {
	r.AddRoute("POST", path, handler, middlewares...)
}

func (r *Router) PUT(path string, handler http.HandlerFunc, middlewares ...Middleware) {
	r.AddRoute("PUT", path, handler, middlewares...)
}

func (r *Router) DELETE(path string, handler http.HandlerFunc, middlewares ...Middleware) {
	r.AddRoute("DELETE", path, handler, middlewares...)
}

func (r *Router) AddRoute(method, path string, handler http.HandlerFunc, middlewares ...Middleware) {
	r.routes = append(r.routes, Route{Method: method, Pattern: path, Handler: handler, Middlewares: middlewares})
}
