package handlers

import (
	"log"
	"net/http"
	"os"
	"twitter/middlew"
	"twitter/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/ver-perfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerPefil))).Methods("GET")
	router.HandleFunc("/modificar-perfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerPefil))).Methods("PUT")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
