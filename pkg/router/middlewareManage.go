package router

func (r *Router) EnableCORS() {
	r.corsEnabled = true
}

func (r *Router) Use(middlewares ...Middleware) {
	r.middlewares = append(r.middlewares, middlewares...)
}
