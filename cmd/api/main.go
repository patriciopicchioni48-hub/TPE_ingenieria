package main

import (
	"log"
	"net/http"

	"github.com/emolina62/trotamundos/internal/recomendacion"
)

func main() {
	// 1. Inicializar dependencias asíncronas
	recomendacion.IniciarMotorAsincrono()

	// 2. Configurar rutas nativas (Go 1.22+)
	mux := http.NewServeMux()

	// Registramos las rutas siguiendo el contrato Spec-Driven
	mux.HandleFunc("POST /api/v1/recomendaciones/{id}/feedback", recomendacion.HandleFeedback)

	// Nueva ruta GET agregada:
	mux.HandleFunc("GET /api/v1/recomendaciones", recomendacion.HandleGetRecomendaciones)

	// 3. Levantar el servidor
	puerto := ":8080"
	log.Printf("🚀 Monolito Trotamundos escuchando en el puerto %s", puerto)
	if err := http.ListenAndServe(puerto, mux); err != nil {
		log.Fatalf("Error iniciando servidor: %v", err)
	}
}
