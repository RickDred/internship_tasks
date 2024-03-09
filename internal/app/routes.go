package app

import "github.com/RickDred/internship_tasks/blob/fourth_task/pkg/router"

func SetRoutes(r *router.Router) {
	r.POST("/api/v1/auth/register")
	r.POST("/api/v1/auth/login")
}
