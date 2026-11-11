package broker

import "log"

// FeedbackEvent representa el mensaje que viaja por la cola
type FeedbackEvent struct {
	UsuarioID       string
	RecomendacionID string
	Accion          string
}

// Cola en memoria con un buffer para absorber picos (QA-03)
var FeedbackQueue = make(chan FeedbackEvent, 1000)

func PublicarFeedback(evento FeedbackEvent) {
	FeedbackQueue <- evento
	log.Printf("[BROKER] Evento encolado: %v", evento)
}
