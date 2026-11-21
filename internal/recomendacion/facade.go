package recomendacion

import (
	"github.com/emolina62/trotamundos/pkg/broker"
)

// RegistrarFeedback actúa como fachada. No procesa IA, solo encola y devuelve el control.
func RegistrarFeedback(usuarioID, recomendacionID, accion string) bool {
	evento := broker.FeedbackEvent{
		UsuarioID:       usuarioID,
		RecomendacionID: recomendacionID,
		Accion:          accion,
	}
	return broker.PublicarFeedback(evento)
}
