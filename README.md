# Proyecto Integrador: Sistema de Chat y Autenticación

**Integrantes:** Juan Pablo Andino, Matías Reyes, Marco Pino

## Descripción
Este es nuestro proyecto integrador para el tercer semestre. El sistema consiste en un chat en tiempo real dividido en dos servicios independientes para separar la lógica de negocio y la carga del servidor.

Decidimos estructurar el backend usando dos lenguajes distintos aprovechando las ventajas de cada uno:

* **Servidor principal (Go):** Maneja toda la lógica del chat y la concurrencia. Utilizamos la arquitectura nativa de Go (goroutines y channels) junto con WebSockets. Esto nos permite mantener las conexiones de los usuarios vivas al mismo tiempo sin saturar la memoria.
* **Servidor de autenticación (Python):** Desarrollado con FastAPI y SQLModel. Se encarga exclusivamente de registrar usuarios, validar el login y guardar los datos en SQLite. Para la seguridad de las contraseñas estamos usando bcrypt combinado con un "pepper" mediante variables de entorno para evitar ataques directos a la base de datos.

## Requisitos previos
Para evitar problemas de compatibilidad o versiones de lenguajes entre los equipos, todo el proyecto está configurado con contenedores. El único requisito es tener instalado Docker Desktop.

## Instrucciones de ejecución

1. Clonar el repositorio localmente:
git clone https://github.com/juanpaAndino/Proyecto-Integrador.git

2. Configurar la seguridad:
Dentro de la carpeta `PythonAuthentication`, crea un archivo llamado `.env`. Adentro solo debes colocar la clave secreta para las contraseñas, de esta forma:
PEPPER_SECRET="tu_clave_secreta_aqui"

3. Levantar los servidores:
Abre una terminal en la raíz del proyecto y ejecuta:
docker compose up --build

Una vez que las imágenes terminen de compilarse, los servicios quedarán expuestos en los siguientes puertos locales:
- Chat en Go: http://localhost:8080
- API de Login y Swagger: http://localhost:8000/docs