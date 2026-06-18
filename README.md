# Proyecto Integrador: Sistema de Chat y Autenticación

**Integrantes:** Juan Pablo Andino, Matías Reyes, Marco Pino

## Descripción
Proyecto integrador de tercer semestre que combina el desarrollo de un sistema de chat distribuido con una infraestructura de red enterprise. El sistema está diseñado bajo principios de **Domain-Driven Design (DDD)** para garantizar escalabilidad y desacoplamiento.

### Componentes Arquitectónicos
* **Backend Chat (Go):** Orquestador de alta concurrencia mediante Goroutines y WebSockets, actuando como API Gateway.
* **Backend Auth (Python/FastAPI):** Servicio especializado en seguridad, utilizando SQLModel para persistencia y bcrypt para el hashing de contraseñas.
* **Infraestructura (Docker):** Despliegue en microservicios gestionado mediante `docker-compose`, asegurando entornos idénticos entre desarrollo y producción.
* **Red (Cisco Packet Tracer):** Topología enterprise con segmentación VLAN, Router-on-a-Stick, seguridad de capa 2 (Port Security) y redundancia con EtherChannel (LACP/PAgP).

## Instrucciones de Ejecución

1. **Clonar repositorio:**
   `git clone https://github.com/juanpaAndino/Proyecto-Integrador.git`

2. **Seguridad (Pepper):**
   Dentro de `PythonAuthentication/`, crea un archivo `.env` con:
   `PEPPER_SECRET="tu_clave_secreta_aqui"`

3. **Despliegue:**
   Ejecuta `docker compose up --build`.
   - API de Auth: `http://localhost:8000/docs`
   - Chat: `http://localhost:8080/cliente.html.html`

4. **Documentación Técnica:**
   La topología de red detallada y las justificaciones de seguridad se encuentran en el archivo `infraestructura_red/topologia_final.pkt`.