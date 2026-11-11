package recomendacion

import (
	"log"
	"time"

	"github.com/emolina62/trotamundos/pkg/broker"
)

func IniciarMotorAsincrono() {
	go func() {
		log.Println("[WORKER] Motor de recomendaciones esperando eventos...")
		for evento := range broker.FeedbackQueue {
			// Simulamos el recálculo pesado del algoritmo de IA
			log.Printf("[WORKER] Procesando feedback de recomendación %s...", evento.RecomendacionID)
			time.Sleep(3 * time.Second)
			log.Printf("[WORKER] ✅ Perfil del usuario %s actualizado (Acción: %s)", evento.UsuarioID, evento.Accion)
		}
	}()
}
