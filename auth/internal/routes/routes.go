package routes

import (
	"auth/internal/handler"

	"github.com/gin-gonic/gin"
)

// ВАЖНО:
// "Маршрутизатор не должен заниматься ничем,
// кроме обработки маршрутов"

// SetupRouter - настраивает все маршруты приложения.
// Принимает обработчик, который содержит логику для работы с пользователями
// Возращает натсроенный роутер
func SetupRouter(h *handler.Handler) *gin.Engine {
	// У router есть свой предустановленный middleware
	// для Логирования и Восстановления полсе сбоев
	router := gin.Default()

	// Group - создает группу Маршрутов. Как?
	// Она берет эндпоинты с ОБЩИМ префиксом (/auth),
	// далее приминяет middleware ко всей группе

	// Группа маршрутов для аутенфикации
	auth := router.Group("/auth")

	{
		// Публичные endpoints
		auth.POST("/register", h.RegisterUser)
		auth.POST("/login", h.LoginUser)

		// Защищенные endpoints
		auth.GET("/user", h.GetInfoUser)
		auth.PUT("/user", h.UpdateUser)
		auth.DELETE("/user", h.DeleteUser)

		// Методы у auth привязывают handlers к
		// конкретным HTTP-методам (POST, GET...),
		// а также URL в группе
	}

	// router - структура gin.Engine, она отвечает ЗА:
	//  1. Маршрутизация
	//	2. Обработка запросов
	//	3. Интеграция middleware

	return router
}
