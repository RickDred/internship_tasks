package app

import (
	"github.com/RickDred/internship_tasks/tree/sixth_task/config"
	authRepository "github.com/RickDred/internship_tasks/tree/sixth_task/internal/auth/repository"
	authService "github.com/RickDred/internship_tasks/tree/sixth_task/internal/auth/service"
	authHttp "github.com/RickDred/internship_tasks/tree/sixth_task/internal/auth/transport/http"
	"github.com/RickDred/internship_tasks/tree/sixth_task/internal/web"
	"github.com/RickDred/internship_tasks/tree/sixth_task/pkg/router"
)

func SetRoutes(r *router.Router, cfg *config.Config) {
	r.EnableCORS()

	// api
	aRepo := authRepository.NewAuthService(cfg)
	aServ := authService.NewAuthService(cfg, aRepo)
	aHandlers := authHttp.NewAuthHandlers(cfg, aServ)

	r.POST("/api/v1/register", aHandlers.Register)
	r.POST("/api/v1/login", aHandlers.Login)

	// web routes
	r.GET("/", web.ProfileHandler)
	r.GET("/login", web.LoginHandler)
	r.GET("/register", web.RegisterHandler)
	r.GET("/static/*", web.StaticFilesHandler)
}
