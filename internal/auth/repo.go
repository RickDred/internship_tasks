package auth

import (
	"github.com/RickDred/internship_tasks/tree/fifth_task/internal/models"
	"github.com/google/uuid"
)

type Repo interface {
	Insert(*models.User) error
	GetByEmail(string) (*models.User, error)
	GetByID(uuid.UUID) (*models.User, error)
	GetAll() ([]models.User, error)
}
