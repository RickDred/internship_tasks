package app

import (
	"github.com/RickDred/internship_tasks/tree/fifth_task/config"
	authRepository "github.com/RickDred/internship_tasks/tree/fifth_task/internal/auth/repository"
	authService "github.com/RickDred/internship_tasks/tree/fifth_task/internal/auth/service"
	authHttp "github.com/RickDred/internship_tasks/tree/fifth_task/internal/auth/transport/http"
	"github.com/RickDred/internship_tasks/tree/fifth_task/pkg/router"
)

func SetRoutes(r *router.Router, cfg *config.Config) {
	aRepo := authRepository.NewAuthService(cfg)
	aServ := authService.NewAuthService(cfg, aRepo)
	aHandlers := authHttp.NewAuthHandlers(cfg, aServ)

	r.POST("/api/v1/register", aHandlers.Register)
	r.POST("/api/v1/login", aHandlers.Login)
}
