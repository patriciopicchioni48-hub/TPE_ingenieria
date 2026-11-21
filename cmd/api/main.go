package main

import (
	"log"
	"net/http"

	"github.com/emolina62/trotamundos/internal/recomendacion"
	"github.com/emolina62/trotamundos/internal/viajes"
)

func main() {
	recomendacion.IniciarMotorAsincrono(3)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/viajes", viajes.HandleGetViajes)
	mux.HandleFunc("POST /api/v1/recomendaciones/{id}/feedback", recomendacion.HandleFeedback)
	mux.HandleFunc("GET /api/v1/recomendaciones", recomendacion.HandleGetRecomendaciones)

	puerto := ":8080"
	log.Printf("Monolito Trotamundos escuchando en el puerto %s", puerto)
	if err := http.ListenAndServe(puerto, mux); err != nil {
		log.Fatalf("Error iniciando servidor: %v", err)
	}
}
