package auth

import (
	"github.com/RickDred/internship_tasks/tree/sixth_task/internal/models"
	"github.com/google/uuid"
)

type Service interface {
	Register(*models.User) error
	Login(*models.User) error
	Profile(uuid.UUID) (*models.User, error)
	Update(*models.User) error
}
