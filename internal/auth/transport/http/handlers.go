package http

import (
	"log"
	"net/http"

	"github.com/RickDred/internship_tasks/tree/sixth_task/config"

	"github.com/RickDred/internship_tasks/tree/sixth_task/internal/auth"
	"github.com/RickDred/internship_tasks/tree/sixth_task/internal/models"

	"github.com/RickDred/internship_tasks/tree/sixth_task/pkg/httperrors"
	"github.com/RickDred/internship_tasks/tree/sixth_task/pkg/httpjson"
)

type Handlers struct {
	cfg     *config.Config
	service auth.Service
}

func NewAuthHandlers(cfg *config.Config, service auth.Service) auth.Handlers {
	return &Handlers{
		cfg:     cfg,
		service: service,
	}
}

func (h *Handlers) Register(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := httpjson.ReadJSON(r, &input); err != nil {
		log.Printf("error while reading json: %v\n", err)

		httpjson.WriteError(w, http.StatusBadRequest, httperrors.NewBadRequestError("invalid credentials request"))
		return
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	if err := h.service.Register(&user); err != nil {
		httpjson.WriteError(w, http.StatusBadRequest, httperrors.NewBadRequestError(err.Error()))
		return
	}

	if err := h.createJwtCookie(w, &user); err != nil {
		httpjson.WriteError(w, http.StatusInternalServerError, httperrors.NewInternalServerError(err.Error()))
		return
	}

	httpjson.WriteJSON(w, &user)
}

// not implemented yet
func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := httpjson.ReadJSON(r, &input); err != nil {
		log.Printf("error while reading json: %v\n", err)

		httpjson.WriteError(w, http.StatusBadRequest, httperrors.NewBadRequestError("bad request"))
		return
	}

	user := models.User{
		Email:    input.Email,
		Password: input.Password,
	}

	if err := h.service.Login(&user); err != nil {
		httpjson.WriteError(w, http.StatusInternalServerError, httperrors.NewInternalServerError(err.Error()))
		return
	}

	if err := h.createJwtCookie(w, &user); err != nil {
		httpjson.WriteError(w, http.StatusInternalServerError, httperrors.NewInternalServerError(err.Error()))
		return
	}

	httpjson.WriteJSON(w, &input)
}

func (h *Handlers) Logout(w http.ResponseWriter, r *http.Request) {}
