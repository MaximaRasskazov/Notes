# Notes Service

Веб-сервер для управления заметками на Go с использованием фреймворка Gin.

## 🚀 Технологии

- **Backend**: Go + Gin
- **Базы данных**: 
  - PostgreSQL (основные данные)
  - MongoDB (дополнительные данные)
- **Контейнеризация**: Docker + Docker Compose
- **Веб-сервер**: Nginx

## 📋 Функциональность (в разработке)

### Планируется реализовать:
- [ ] REST API с использованием Gin
- [ ] Система аутентификации и авторизации
- [ ] CRUD операции для заметок
- [ ] Подключение PostgreSQL
- [ ] Подключение MongoDB  
- [ ] Docker-контейнеризация
- [ ] Настройка Nginx как reverse proxy

## 🏗️ Структура проекта (предварительная)

```
notes_go/
├── cmd/
│ └── server/
├── internal/
│ ├── handlers/
│ ├── models/
│ ├── repository/
│ └── middleware/
├── pkg/
├── docker/
├── docker-compose.yml
└── README.md
```