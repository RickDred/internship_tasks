package httpjson

import (
	"encoding/json"
	"log"
	"net/http"
)

func ReadJSON(r *http.Request, dst interface{}) error {
	err := json.NewDecoder(r.Body).Decode(dst)
	return err
}

func WriteJSON(w http.ResponseWriter, data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(http.StatusText(http.StatusInternalServerError))); err != nil {
			log.Printf("Write failed: %v\n", err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if _, err := w.Write(jsonData); err != nil {
		log.Printf("Write failed: %v\n", err)
	}
}

func WriteError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	er := struct {
		Message string `json:"message"`
	}{err.Error()}
	WriteJSON(w, er)
}
