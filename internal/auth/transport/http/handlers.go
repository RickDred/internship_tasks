package http

import (
	"log"
	"net/http"
	"time"

	"github.com/RickDred/internship_tasks/blob/fifth_task/internal/models"
	"github.com/RickDred/internship_tasks/blob/fifth_task/pkg/httperrors"
	"github.com/RickDred/internship_tasks/blob/fifth_task/pkg/httpjson"
)

type Handlers struct{}

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := httpjson.ReadJSON(r, &user); err != nil {
		log.Printf("error while reading json: %v\n", err)

		httpjson.WriteError(w, httperrors.NewBadRequestError("error"))
		return
	}

	respUser := struct {
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"createdAt"`
	}{user.Name, user.Email, time.Now()}

	httpjson.WriteJSON(w, &respUser)
}

// not implemented yet
func Login(w http.ResponseWriter, r *http.Request) {}

func Logout(w http.ResponseWriter, r *http.Request) {}
