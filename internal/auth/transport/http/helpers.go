package http

import (
	"net/http"

	"github.com/RickDred/internship_tasks/tree/sixth_task/internal/models"
	"github.com/golang-jwt/jwt"
)

func (h *Handlers) createJwtCookie(w http.ResponseWriter, u *models.User) error {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid": u.ID,
		"name":   u.Name,
		"email":  u.Email,
		"role":   u.Role,
	})

	tokenString, err := token.SignedString([]byte(h.cfg.Server.JwtSecretKey))
	if err != nil {
		return err
	}

	cookie := http.Cookie{
		Value:    tokenString,
		Path:     "/",
		Name:     h.cfg.Cookie.Name,
		Secure:   h.cfg.Cookie.Secure,
		HttpOnly: h.cfg.Cookie.HTTPOnly,
		MaxAge:   h.cfg.Cookie.MaxAge,
	}

	http.SetCookie(w, &cookie)
	return nil
}
