# 🧠 Addis-Hiwot Backend

A simple Go backend service using **GORM**, **PostgreSQL**, and **Docker**. Follows clean architecture principles with modular layers for models, usecases, repositories, and handlers.

## 🚀 Quick Start

````bash
# Clone the repo
git clone https://github.com/Addis-Hiwot-Team/Addis-Hiwot.git
cd Backend

## ⚙️ Environment Setup

Create a `.env` file in the root of the project with the following variables:

```env
# Database Configuration
DB_HOST=your_db_host
DB_PORT=your_db_port
DB_USER=your_db_username
DB_PASSWORD=your_db_password
DB_NAME=your_db_name

# Application Configuration
APP_PORT=your_app_port
````

Replace the placeholder values with your actual configuration.
These variables will be used by both the application and Docker.

# 📦 Tech Stack

- Go
- GORM
- PostgreSQL
- Docker

# 📂 Structure

- `models/` – GORM models
- `usecases/` – Business logic
- `repository/` – Database layer
- `handlers/` – HTTP endpoints

🛠 Sample API

| Method | Endpoint | Description     |
| ------ | -------- | --------------- |
| POST   | /users   | Create new user |
| GET    | /users   | Get all users   |

# 🧪 Sample Request

### POST /users

```json
{
  "email": "user@example.com",
  "username": "nebiyu",
  "password": "securePassword123",
  "role": "admin"
}
```

# Sample Response

```json
{
  "user_id": 1,
  "email": "user@example.com",
  "username": "nebiyu",
  "role": "admin",
  "is_active": true,
  "created_at": "2025-06-15T10:00:00Z"
}
```

# Run with Docker

docker-compose up --build
