package recomendacion

import (
	"log"
	"time"

	"github.com/emolina62/trotamundos/pkg/broker"
)

func IniciarMotorAsincrono(numWorkers int) {
	for i := 1; i <= numWorkers; i++ {
		go func(workerID int) {
			log.Printf("[WORKER-%d] Motor de recomendaciones esperando eventos...", workerID)
			for evento := range broker.FeedbackQueue {
				log.Printf("[WORKER-%d] Procesando feedback de recomendación %s...", workerID, evento.RecomendacionID)
				time.Sleep(3 * time.Second)
				log.Printf("[WORKER-%d] Perfil del usuario %s actualizado (Acción: %s)", workerID, evento.UsuarioID, evento.Accion)
			}
		}(i)
	}
}
