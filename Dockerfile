# Etapa 1: Compilación (Builder)
FROM golang:1.22-alpine AS builder

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos de dependencias y descargarlas
COPY go.mod ./
# COPY go.sum ./ (Descomentar si luego agregan dependencias externas)
RUN go mod download

# Copiar todo el código fuente al contenedor
COPY . .

# Compilar el ejecutable estático para Linux
# Se apunta específicamente al punto de entrada del monolito
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o trotamundos cmd/api/main.go

# Etapa 2: Imagen de Producción (Final)
FROM alpine:latest

# Agregar certificados de seguridad para conexiones HTTPS externas (Terceros)
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Copiar únicamente el binario compilado desde la etapa anterior
COPY --from=builder /app/trotamundos .

# Exponer el puerto definido en el código
EXPOSE 8080

# Comando por defecto al iniciar el contenedor
CMD ["./trotamundos"]