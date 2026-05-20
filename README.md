# Sistema de Chat en Go - Análisis y Arquitectura (Etapa 2)

## 1. Evaluación de Arquitectura y Estándares (Evidencia)
Tras analizar las arquitecturas tradicionales orientadas a objetos, determinamos que aplicar un modelo fuertemente acoplado sería un error para un sistema de chat en tiempo real. En su lugar, cuestionamos ese estándar y adoptamos un enfoque de **Domain-Driven Design (DDD) simplificado**. Go nos permite aislar el dominio organizando el código en paquetes estratégicos (`api`, `models`, `websocket`), lo que facilita la mantenibilidad y evita las dependencias circulares comunes en otros frameworks.

## 2. Posición Específica y Decisiones Técnicas
Nuestra postura técnica se basa en priorizar la seguridad y el rendimiento concurrente, reconociendo las complejidades de un chat en vivo:
* **Seguridad (Hashing):** Rechazamos el uso de librerías criptográficas obsoletas o genéricas. Nuestra posición es implementar **únicamente `bcrypt`** para el hashing de contraseñas, ya que incorpora un factor de costo ("work factor") que mitiga ataques de fuerza bruta, adaptándose a la capacidad computacional actual.
* **Concurrencia (WebSockets):** En lugar de saturar el servidor HTTP con peticiones continuas (Long Polling), diseñamos una estructura `Hub` y `Client` en el paquete `websocket`. Esta decisión aprovecha las *goroutines* y *channels* nativos de Go, permitiendo manejar miles de conexiones simultáneas con un consumo mínimo de recursos.

## 3. Conclusiones y Logros Relacionados
La estructura implementada no es solo organizativa, sino estratégica. Al separar los controladores HTTP (endpoints `/register`, `/login`) del gestor de WebSockets, logramos un sistema donde el tráfico pesado de mensajes no bloquea los procesos de autenticación. Esta priorización de la evidencia técnica sobre las convenciones básicas garantiza que la base del proyecto es altamente escalable y tolerante a alta demanda, cumpliendo con los requisitos críticos de la fase de diseño.