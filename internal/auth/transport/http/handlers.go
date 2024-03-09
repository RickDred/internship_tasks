package http

import (
	"log"
	"net/http"

	"github.com/RickDred/internship_tasks/blob/fourth_task/internal/models"
	"github.com/RickDred/internship_tasks/blob/fourth_task/pkg/httpjson"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := httpjson.ReadJSON(r, &user); err != nil {
		log.Printf("error while reading json: %v\n", err)

		httpjson.WriteError(w, httperrors.)
	}
}
func Login(w http.ResponseWriter, r *http.Request) {}
