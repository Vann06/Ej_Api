# 📋 Incidents API

Esta es una API REST desarrollada en **Go** usando el framework **Gin**, que permite crear, consultar, actualizar y eliminar reportes de incidentes.

---

## 🚀 Características actuales

✅ Crear un nuevo incidente  
✅ Listar todos los incidentes  
✅ Consultar un incidente por ID  
✅ Actualizar el estado de un incidente  
✅ Eliminar un incidente  

---

## 📦 Tecnologías utilizadas

- [Go](https://golang.org/) 1.24+
- [Gin](https://github.com/gin-gonic/gin) como framework web
- (Próximamente: MySQL como base de datos)

---

## ⚙️ Instalación y ejecución

### 1. Clona el repositorio

\`\`\`bash
git clone https://github.com/Vann06/Ej_Api.git
cd Ej_Api
\`\`\`

### 2. Inicializa el proyecto de Go

\`\`\`bash
go mod tidy
\`\`\`

### 3. Ejecuta la API

\`\`\`bash
go run cmd/main/main.go
\`\`\`

La API se ejecutará en:  
📍 \`http://localhost:8080\`

---

## 📌 Endpoints disponibles

| Método | Endpoint           | Descripción                              |
|--------|--------------------|------------------------------------------|
| GET    | \`/incidents\`       | Lista todos los incidentes               |
| GET    | \`/incidents/:id\`   | Consulta un incidente por su ID          |
| POST   | \`/incidents\`       | Crea un nuevo incidente                  |
| PUT    | \`/incidents/:id\`   | Actualiza el estado de un incidente      |
| DELETE | \`/incidents/:id\`   | Elimina un incidente reportado por error |

---

## 📤 Ejemplo para crear un incidente

### Método: \`POST /incidents\`  
Body JSON:

\`\`\`json
{
  "reporter": "Zeyda",
  "description": "Mi laptop no enciende"
}
\`\`\`

---

## 📌 Notas

- Por ahora, los incidentes se guardan en memoria (slice).
- La conexión con base de datos MySQL se agregará más adelante.
- Código modular con carpetas: \`models/\`, \`controllers/\`, \`routes/\`.

---

## 🛠️ Cómo instalar Go

### 🔹 Windows

1. Ve a la página oficial: https://go.dev/dl/
2. Descarga el instalador para Windows (\`.msi\`) según tu sistema (64-bit usualmente).
3. Ejecuta el instalador y sigue los pasos.
4. Reinicia tu terminal y escribe:

\`\`\`bash
go version
\`\`\`

✅ Deberías ver algo como: \`go version go1.24.1 windows/amd64\`

---

### 🔹 macOS (usando Homebrew)

\`\`\`bash
brew install go
\`\`\`

Verifica con:

\`\`\`bash
go version
\`\`\`

---

### 🔹 Linux (Debian/Ubuntu)

\`\`\`bash
sudo apt update
sudo apt install golang-go
\`\`\`

Verifica con:

\`\`\`bash
go version
\`\`\`

---


Puedes verificar con:

\`\`\`bash
go env
\`\`\`
EOF
