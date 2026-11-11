package recomendacion

import (
	"encoding/json"
	"net/http"
)

type FeedbackRequest struct {
	Accion string `json:"accion"`
}

func HandleFeedback(w http.ResponseWriter, r *http.Request) {
	recomendacionID := r.PathValue("id")

	var req FeedbackRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Cuerpo inválido", http.StatusBadRequest)
		return
	}

	usuarioMock := "usr-999"

	RegistrarFeedback(usuarioMock, recomendacionID, req.Accion)

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"status": "procesando en background"}`))
}

type Recomendacion struct {
	ID         string  `json:"id"`
	Tipo       string  `json:"tipo"`
	Titulo     string  `json:"titulo"`
	MatchScore float64 `json:"match_score"`
}

func HandleGetRecomendaciones(w http.ResponseWriter, r *http.Request) {
	// Simulamos una lectura ultrarrápida de datos ya pre-calculados por el Worker
	sugerencias := []Recomendacion{
		{
			ID:         "rec-101",
			Tipo:       "alojamiento",
			Titulo:     "Hotel Llao Llao - Oferta exclusiva",
			MatchScore: 0.95,
		},
		{
			ID:         "rec-102",
			Tipo:       "ruta",
			Titulo:     "Circuito Chico con paradas fotográficas",
			MatchScore: 0.88,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Devolvemos el JSON inmediatamente (Cumple QA-01 < 800ms)
	json.NewEncoder(w).Encode(sugerencias)
}
