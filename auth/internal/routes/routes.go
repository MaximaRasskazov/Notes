package routes

import (
	"auth/internal/handler"

	"github.com/gin-gonic/gin"
)

// ВАЖНО:
// Маршрутизатор не должен заниматься ничем, кроме
// обработки маршрутов.

// SetupRouter - настраивает все маршруты приложения.
// Принимает обработчик, который содержит логику для работы с пользователями
// Возращает натсроенные роутер
func SetupRouter(h *handler.Handler) *gin.Engine {
	return nil
}
