package router

import (
	"context"
	"net/http"
	"strings"
)

func (r *Router) routesMatch(path, pattern string) (bool, map[string]string) {
	patternParts := strings.Split(pattern, "/")
	pathParts := strings.Split(path, "/")

	if len(patternParts) != len(pathParts) {
		return false, nil
	}

	params := make(map[string]string)

	for i := 0; i < len(patternParts); i++ {
		patternPart := patternParts[i]
		pathPart := pathParts[i]

		if strings.HasPrefix(patternPart, ":") {
			params[patternPart[1:]] = pathPart
		} else if patternPart != pathPart {
			return false, nil
		}
	}

	return true, params
}

func (r *Router) findHandler(req *http.Request) http.HandlerFunc {
	method := req.Method
	path := strings.TrimRight(req.URL.Path, "/")

	handler := r.NotFoundHandler
	found := false
	methodNotAllowed := false

	for _, route := range r.routes {
		pattern := strings.TrimRight(route.Pattern, "/")

		ok, params := r.routesMatch(path, pattern)
		if !ok {
			continue
		}
		found = true

		if route.Method != method {
			methodNotAllowed = true
			continue
		}

		handler = route.Handler

		for _, middleware := range route.Middlewares {
			handler = middleware(handler)
		}
		ctx := req.Context()

		for key, value := range params {
			ctx = context.WithValue(ctx, key, value)
		}

		req = req.WithContext(ctx)

		methodNotAllowed = false
		break
	}

	if !found {
		return r.NotFoundHandler
	}

	if methodNotAllowed {
		return r.MethodNotAllowed
	}
	return handler
}

func DefaultMethodNotAllowedHandler(w http.ResponseWriter, req *http.Request) {
	http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
}
