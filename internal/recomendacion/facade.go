package recomendacion

import (
	"github.com/emolina62/trotamundos/pkg/broker"
)

// RegistrarFeedback actúa como fachada. No procesa IA, solo encola y devuelve el control.
func RegistrarFeedback(usuarioID, recomendacionID, accion string) {
	evento := broker.FeedbackEvent{
		UsuarioID:       usuarioID,
		RecomendacionID: recomendacionID,
		Accion:          accion,
	}
	// Enviamos el mensaje al broker de forma asíncrona
	broker.PublicarFeedback(evento)
}
