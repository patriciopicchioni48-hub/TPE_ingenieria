# Trotamundos - Prototipo Arquitectónico

Prototipo funcional del sistema de gestión de viajes **Trotamundos**, desarrollado como prueba de concepto para validar las decisiones arquitectónicas tomadas durante el proceso de Diseño Dirigido por Atributos (ADD).

Este prototipo implementa el **Módulo Avanzado de Recomendaciones ** dentro de una arquitectura de **Monolito Modular** escrita en Go.

## 🏗️ Decisiones Arquitectónicas Implementadas

El código fuente materializa los siguientes Registros de Decisiones Arquitectónicas (ADRs):

*   **ADR-01 (Monolito Modular):** Código segregado por dominios de negocio en la carpeta `/internal`, compilado en un único ejecutable.
*   **Spec-Driven Development:** Contrato de la API definido estáticamente en `/docs/api/openapi.yaml`.
*   **ADR-09 y ADR-11 (Procesamiento Asíncrono):** Uso de un *Message Broker* en memoria y *Workers* en background (goroutines) para procesar el *feedback* del algoritmo de IA sin bloquear los hilos HTTP. Garantiza el cumplimiento de **QA-01 (Latencia < 800ms)** y **QA-03 (Escalabilidad)**.

## ⚙️ Requisitos Previos

*   [Docker](https://www.docker.com/) (v24 o superior)
*   Git

## 🚀 Cómo ejecutar el proyecto

El proyecto está completamente contenerizado (Multi-stage build) para garantizar su ejecución en cualquier entorno.

1. Clonar el repositorio:
   ```bash
   git clone [https://github.com/patriciopicchioni48-hub/TPE_ingenieria.git](https://github.com/patriciopicchioni48-hub/TPE_ingenieria.git)
   cd TPE_ingenieria
   ```

2. Levantar la infraestructura con Docker Compose:
   ```bash
   docker compose up -d --build
   ```

3. Verificar que el servidor y el Worker estén corriendo:
   ```bash
   docker compose logs -f api
   ```

## 🧪 Pruebas de Concepto (Demostración de Drivers)

Con el contenedor corriendo, abrí una terminal nueva y ejecutá las siguientes pruebas para validar el cumplimiento de los Atributos de Calidad:

### Prueba 1: Lectura Ultrarrápida (QA-01)
Demuestra que la consulta de recomendaciones pre-calculadas responde de forma inmediata, cumpliendo con el umbral de 800ms para aplicaciones web.

```bash
curl -X GET http://localhost:8080/api/v1/recomendaciones
```
*Resultado esperado:* Retorna inmediatamente un array JSON con las sugerencias.

### Prueba 2: Feedback Asíncrono (Human-in-the-loop)
Demuestra que al interactuar con una sugerencia, el servidor responde al instante (HTTP 202) y encola el trabajo pesado.

```bash
curl -X POST http://localhost:8080/api/v1/recomendaciones/rec-123/feedback \
  -H "Content-Type: application/json" \
  -d '{"accion": "rechazar"}'
```
*Resultado esperado:* El cliente recibe respuesta en microsegundos. En los logs de Docker (`docker compose logs -f api`), se observa cómo el Broker encola el evento y el Worker lo procesa tomando 3 segundos.

### Prueba 3: Resiliencia ante Picos de Tráfico (QA-03)
Simula la llegada de 10 peticiones concurrentes para ajustar parámetros de IA.

```bash
for i in {1..10}; do
  curl -s -X POST http://localhost:8080/api/v1/recomendaciones/rec-$i/feedback \
  -H "Content-Type: application/json" \
  -d '{"accion": "aceptar"}' &
done; wait; echo -e "\n¡Todas las peticiones fueron recibidas!"
```
*Resultado esperado:* La terminal del cliente finaliza al instante. En los logs del servidor se observa cómo el sistema absorbe el pico de tráfico encolando los 10 eventos simultáneos, y el Worker los procesa de forma secuencial y aislada, evitando el agotamiento de recursos (*thread starvation*).