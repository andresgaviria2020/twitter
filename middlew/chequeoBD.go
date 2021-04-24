package middlew

import (
	"net/http"
	"twitter/twitter/bd"
)

func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(rw, "Conexión perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(rw, r)
	}
}
