package service

import (
	"github.com/RickDred/internship_tasks/tree/fifth_task/config"
	"github.com/RickDred/internship_tasks/tree/fifth_task/internal/auth"
	"github.com/RickDred/internship_tasks/tree/fifth_task/internal/models"
	"github.com/RickDred/internship_tasks/tree/fifth_task/pkg/httperrors"
)

type Service struct {
	cfg  *config.Config
	repo auth.Repo
}

func NewAuthService(cfg *config.Config, repo auth.Repo) auth.Service {
	return &Service{
		cfg:  cfg,
		repo: repo,
	}
}

func (s *Service) Register(user *models.User) error {
	user.Standardize()

	_, err := s.repo.GetByEmail(user.Email)
	if err == nil {
		return httperrors.NewBadRequestError("user with this email already exist")
	}

	if err := user.HashPassword(); err != nil {
		return httperrors.NewInternalServerError("hashing error")
	}

	if ok := user.IsValid(); !ok {
		return httperrors.NewBadRequestError("user credentials are not valid")
	}

	if err := s.repo.Insert(user); err != nil {
		return httperrors.NewInternalServerError("inserting user error")
	}

	user.CleanPassword()

	return nil
}

func (s *Service) Login(user *models.User) error {
	user.Standardize()

	exUser, err := s.repo.GetByEmail(user.Email)
	if err != nil {
		return httperrors.NewBadRequestError("wrong email or password")
	}

	if err := exUser.ComparePasswords(user.Password); err != nil {
		return httperrors.NewBadRequestError("wrong email or password")
	}

	user.Name = exUser.Name
	user.Role = exUser.Role
	user.CreatedAt = exUser.CreatedAt

	user.CleanPassword()

	return nil
}

func (s *Service) Logout(u *models.User) error {
	return nil
}
