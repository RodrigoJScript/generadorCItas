#  Generador de Citas / Quote Generator

<div align="center">

![Go Version](https://img.shields.io/badge/Go-1.22.2-blue.svg)
![License](https://img.shields.io/badge/License-MIT-green.svg)
![Status](https://img.shields.io/badge/Status-Active-success.svg)

**Una aplicación web simple y elegante para generar citas inspiracionales aleatorias**

*A simple and elegant web application to generate random inspirational quotes*

</div>

---

##  Descripción

**Generador de Citas** es un proyecto educativo desarrollado como parte de mi aprendizaje en Go. Esta aplicación web permite a los usuarios generar citas inspiracionales aleatorias y ver un historial de todas las citas generadas durante la sesión.

El proyecto demuestra conceptos fundamentales de Go como:
- Manejo de servidores HTTP
- Servicio de archivos estáticos
- Creación de APIs REST simples
- Manejo de JSON
- Concurrencia básica

##  Características

-  **Generación aleatoria de citas**: Obtén una cita inspiracional aleatoria con un solo clic
-  **Historial de citas**: Visualiza todas las citas generadas durante la sesión
-  **Interfaz moderna**: Diseño limpio y responsivo usando **Pico CSS**, un framework CSS minimalista y elegante
- **Rendimiento rápido**: Servidor HTTP ligero y eficiente en Go
- **API RESTful**: Endpoints simples y bien estructurados

##  Tecnologías

### Backend
- **Go 1.22.2** - Lenguaje de programación principal
- **net/http** - Servidor HTTP estándar de Go

### Frontend
- **HTML5** - Estructura de la aplicación
- **CSS3** - Estilos personalizados
- **Pico CSS** - Framework CSS minimalista utilizado para el diseño de la interfaz
- **JavaScript** - Interactividad del cliente

##  Requisitos

- Go 1.22.2 o superior
- Navegador web moderno (Chrome, Firefox, Safari, Edge)

##  Instalación

1. **Clona el repositorio**

2. **Verifica la instalación de Go**
   ```bash
   go version
   ```

3. **Instala las dependencias**
   ```bash
   go mod download
   ```

4. **Ejecuta la aplicación**
   ```bash
   go run main.go
   ```

5. **Abre tu navegador**
   ```
   http://localhost:8080
   ```

##  Uso

1. **Generar una cita**: Haz clic en el botón "Generar Cita" para obtener una cita inspiracional aleatoria
2. **Ver historial**: Haz clic en "Ver Citas Generadas" para ver todas las citas que has generado en esta sesión
3. **Disfrutar**: ¡Disfruta de las citas inspiracionales y compártelas!


##  API Endpoints

### `GET /api/quote`

Genera una cita aleatoria y la agrega al historial.

**Respuesta:**
```json
{
  "quote": "La vida es bella",
  "total": 1
}
```

### `GET /api/quotes`

Obtiene todas las citas generadas durante la sesión.

**Respuesta:**
```json
{
  "quotes": [
    "La vida es bella",
    "El camino es el destino",
    "El obstáculo es el camino"
  ],
  "total": 3
}
```

##  Funcionalidades Técnicas

- **Servidor HTTP**: Implementación usando el paquete estándar `net/http`
- **Servicio de archivos estáticos**: Archivos CSS, JS servidos eficientemente
- **API REST**: Endpoints JSON para comunicación cliente-servidor
- **Almacenamiento en memoria**: Las citas se almacenan en un slice durante la sesión
- **Manejo de errores**: Validación de métodos HTTP y manejo de errores apropiado
- **Framework CSS**: Utiliza **Pico CSS** para un diseño minimalista y responsivo sin necesidad de configuración compleja

##  Pruebas

Para probar los endpoints de la API, puedes usar `curl`:

```bash
# Generar una cita
curl http://localhost:8080/api/quote

# Obtener todas las citas
curl http://localhost:8080/api/quotes
```

##  Notas

Este proyecto fue creado como parte de mi proceso de aprendizaje de Go. Es una aplicación simple pero funcional que demuestra los conceptos básicos de desarrollo web con Go.

**Nota importante**: Las citas generadas se almacenan solo en memoria. Al reiniciar el servidor, el historial se perderá.

**Sobre Pico CSS**: Este proyecto utiliza Pico CSS, un framework CSS minimalista que proporciona estilos elegantes y responsivos sin necesidad de clases CSS personalizadas. Pico CSS se integra perfectamente con HTML semántico y ofrece un diseño moderno con mínimo esfuerzo.

##  Contribuciones

Las contribuciones son bienvenidas. Siéntete libre de abrir un issue o enviar un pull request.

##  Licencia

Este proyecto está bajo la Licencia MIT. Ver el archivo `LICENSE` para más detalles.

##  Autor

**RodrigoJScript**

- GitHub: [@RodrigoJScript](https://github.com/RodrigoJScript)
- Email: rodrigo777.rjp@gmail.com
