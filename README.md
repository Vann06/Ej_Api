# 📋 Incidents API

Esta es una API REST desarrollada en **Go** usando el framework **Gin**, que permite crear, consultar, actualizar y eliminar reportes de incidentes en una empresa (como computadoras, impresoras, redes, etc.).

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
- [MySQL](https://www.mysql.com/) como base de datos relacional

---

## 📁 Estructura del proyecto

```
EJ_API/
├── cmd/
│   └── main/
│       └── main.go
├── db/
│   └── db.go
├── pkg/
│   ├── controllers/
│   │   └── incident-controller.go
│   ├── models/
│   │   └── incident.go
│   └── routes/
│       └── incidents-routes.go
├── go.mod
├── go.sum
└── README.md
```

---

## ⚙️ Instalación y ejecución

### 1. Clona el repositorio

```bash
git clone https://github.com/Vann06/Ej_Api.git
cd Ej_Api
```

### 2. Instala dependencias de Go

```bash
go mod tidy
```

### 3. Crea la base de datos MySQL

```sql
CREATE DATABASE incidentes;

USE incidentes;

CREATE TABLE ticket (
  id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
  reporter VARCHAR(100) NOT NULL,
  description VARCHAR(250) NULL,
  status VARCHAR(20) NOT NULL DEFAULT 'pendiente',
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
);
```

### 4. Ejecuta la API

```bash
go run cmd/main/main.go
```

Servidor disponible en:  
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

```json
{
  "reporter": "Zeyda",
  "description": "Mi laptop no enciende"
}
```

---

## 📌 Reglas de negocio

- \`reporter\` es obligatorio  
- \`description\` debe tener al menos 10 caracteres  
- Solo puede actualizarse el campo \`status\` (PUT)  
- Si un incidente no existe, se devuelve error 404  

---

## 🛠️ Cómo instalar Go

### 🔹 Windows

1. Ve a: https://go.dev/dl/
2. Descarga el instalador \`.msi\` (Windows 64-bit recomendado)
3. Instálalo como cualquier programa
4. Verifica en terminal:

```bash
go version
```

✅ Resultado esperado: \`go version go1.24.1 windows/amd64\`

---

### 🔹 macOS (Homebrew)

```bash
brew install go
go version
```

---

### 🔹 Linux (Debian/Ubuntu)

```bash
sudo apt update
sudo apt install golang-go
go version
```

---

### Verifica tu entorno:

```bash
go env
```

---

## 👩‍💻 Autor

Vianka Castro ✨  

