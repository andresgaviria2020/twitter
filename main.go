package main

import (
	"log"

	"twitter/bd"
	"twitter/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la bd")
		return
	}
	handlers.Manejadores()
}
