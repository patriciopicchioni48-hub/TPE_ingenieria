package viajes

import (
	"encoding/json"
	"net/http"
)

type Viaje struct {
	ID     string `json:"id"`
	Nombre string `json:"nombre"`
	Estado string `json:"estado"`
}

func HandleGetViajes(w http.ResponseWriter, r *http.Request) {
	lista := []Viaje{
		{
			ID:     "viaje-001",
			Nombre: "Vacaciones en Bariloche",
			Estado: "activo",
		},
		{
			ID:     "viaje-002",
			Nombre: "Escapada a Mendoza",
			Estado: "archivado",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(lista)
}
