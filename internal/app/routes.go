package app

import (
	authHttp "github.com/RickDred/internship_tasks/blob/fifth_task/internal/auth/transport/http"
	"github.com/RickDred/internship_tasks/blob/fifth_task/pkg/router"
)

func SetRoutes(r *router.Router) {
	r.POST("/api/v1/auth/register", authHttp.Register)
}
