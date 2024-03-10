package auth

import "net/http"

type Handlers interface {
	Register(http.ResponseWriter, *http.Request)
	Login(http.ResponseWriter, *http.Request)
	Logout(http.ResponseWriter, *http.Request)
	Profile(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
}
