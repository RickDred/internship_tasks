package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/RickDred/internship_tasks/tree/sixth_task/internal/models"
	"github.com/RickDred/internship_tasks/tree/sixth_task/pkg/httperrors"
	"github.com/RickDred/internship_tasks/tree/sixth_task/pkg/httpjson"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func (h *Handlers) setError(w http.ResponseWriter, err error, defaultStatus int) {

	switch e := err.(type) {
	case httperrors.BadRequestError:
		httpjson.WriteError(w, http.StatusBadRequest, e)
	case httperrors.UnauthorizedError:
		httpjson.WriteError(w, http.StatusUnauthorized, e)
	case httperrors.NotFoundError:
		httpjson.WriteError(w, http.StatusNotFound, e)
	case httperrors.MethodNotAllowedError:
		httpjson.WriteError(w, http.StatusMethodNotAllowed, e)
	case httperrors.InternalServerError:
		httpjson.WriteError(w, http.StatusInternalServerError, e)
	default:
		httpjson.WriteError(w, defaultStatus, e)
	}

}

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

func (h *Handlers) cleanCookie(w http.ResponseWriter) {
	cookie := http.Cookie{
		Value:   "",
		Path:    "/",
		Name:    h.cfg.Cookie.Name,
		Expires: time.Now().AddDate(-1, 0, 0),
	}

	http.SetCookie(w, &cookie)
}

func (h *Handlers) getUserFromCookie(r *http.Request) (*models.User, error) {
	token, err := h.getCookieValue(r)
	if err != nil {
		return nil, err
	}

	claims, err := h.extractClaims(token)
	if err != nil {
		return nil, err
	}

	user := &models.User{}

	// Extracting user information from claims
	if userIDStr, ok := claims["userid"].(string); ok {
		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			return user, fmt.Errorf("failed to parse user ID: %v", err)
		}
		user.ID = userID
	} else {
		return user, fmt.Errorf("userid not found or not a string")
	}

	if name, ok := claims["name"].(string); ok {
		user.Name = name
	} else {
		return user, fmt.Errorf("name not found or not a string")
	}

	if email, ok := claims["email"].(string); ok {
		user.Email = email
	} else {
		return user, fmt.Errorf("email not found or not a string")
	}

	if role, ok := claims["role"].(string); ok {
		user.Role = role
	} else {
		return user, fmt.Errorf("role not found or not a string")
	}

	return user, nil
}

func (h *Handlers) getCookieValue(r *http.Request) (string, error) {
	cookie, err := r.Cookie(h.cfg.Cookie.Name)
	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}

func (h *Handlers) extractClaims(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return h.cfg.Server.JwtSecretKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
