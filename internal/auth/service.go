package auth

import (
	"github.com/RickDred/internship_tasks/tree/fifth_task/internal/models"
)

type Service interface {
	Register(*models.User) error
	Login(*models.User) error
	Logout(*models.User) error
}
