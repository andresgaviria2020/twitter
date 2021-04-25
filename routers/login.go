package routers

import (
	"encoding/json"
	"net/http"
	"time"
	"twitter/bd"
	"twitter/models"
	"twitter/utilidades"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "", 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "", 400)
		return
	}
	documento, existe := bd.IntentoLogin(t.Email, t.Password)
	if existe == false {
		http.Error(w, "", 400)
		return
	}

	jwtKey, err := utilidades.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "", 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Contect-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
