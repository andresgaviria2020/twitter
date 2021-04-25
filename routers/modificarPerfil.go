package routers

import (
	"encoding/json"
	"net/http"
	"twitter/bd"
	"twitter/models"
)

func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), 400)
		return
	}
	var status bool
	status, err = bd.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "Reitente nuevamente"+err.Error(), 400)
		return
	}

	if status != true {
		http.Error(w, "Reitente nuevamente"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
