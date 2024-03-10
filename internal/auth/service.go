package auth

import (
	"github.com/RickDred/internship_tasks/blob/fifth_task/internal/models"
)

type Service interface {
	Register(models.User) 
	Login(models.User) 
	Logout(models.User) 
}
