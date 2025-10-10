package handler

import (
	"auth/internal/config"

	"github.com/gin-gonic/gin"
)

// Handler содержит все обработчики, методы нужные
// для работы с пользователями (HTTP-запросы)
type Handler struct {
	cfg *config.Config
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{
		cfg: cfg,
	}
}

func (h Handler) RegisterUser(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "Пользователь успешно зарегистрирован",
	})
}

func (h Handler) LoginUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":       "Пользователь успешно авторизован",
		"access_token":  "example_access_token",
		"refresh_token": "example_refresh_token",
	})
}

func (h Handler) GetInfoUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Информация о пользователе получена",
	})
}

func (h Handler) UpdateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Данные пользователья обновлены",
	})
}

func (h Handler) DeleteUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Пользователь успешно удален",
	})
}
