# -- Etapa 1: Construcción (Builder) --
FROM golang:1.26-alpine AS builder

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los gestores de dependencias y descargarlas
# (Asegúrate de haber ejecutado 'go mod tidy' en tu terminal antes)
COPY go.mod go.sum* ./
RUN go mod download

# Copiar todo el código fuente del proyecto
COPY . .

# Compilar el binario. CGO_ENABLED=0 asegura un binario estático compatible con Alpine
RUN CGO_ENABLED=0 GOOS=linux go build -o chat-server ./cmd/server/main.go

# -- Etapa 2: Producción --
FROM alpine:latest

# Añadir certificados de seguridad por si haces peticiones HTTPS en el futuro
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiar SOLO el binario compilado de la etapa anterior
COPY --from=builder /app/chat-server .

# Exponer el puerto que definimos en main.go
EXPOSE 8080

# Ejecutar el binario al arrancar el contenedor
CMD ["./chat-server"]