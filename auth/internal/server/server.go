package server

import (
	"auth/internal/config"
	"auth/internal/handler"
	"auth/internal/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

// У сервера должны быть:
// 	1. Конфигурации
// 	2. Обработчики
//  3. Запуск

// Структура сервера
type Server struct {
	// Конфиг (Порт, хост и тд)
	cfg *config.Config

	// Маршрутизатор ("/auth/register" вызывает RegisterUser)
	router *gin.Engine
}

// Конструктор сервера
func NewServer(cfg *config.Config) (*Server, error) {
	// Проверим, что конфиг не пустой
	if cfg == nil {
		return nil, fmt.Errorf("конфигурация сервера не может быть nil")
	}

	// Новый экземпляр обработчика
	handler := handler.NewHandler(cfg)
	if handler == nil {
		return nil, fmt.Errorf("не удалось создать обработчик сервера")
	}
	fmt.Println("Обработчика сервера успешно создан")

	// Создаем экземпляр маршрутизатора
	router := routes.SetupRouter(handler)

	// Раз всё окей, то создадим сервер
	return &Server{
		router: router,
		cfg:    cfg,
	}, nil
}

func (s *Server) Serve() error {
	// Запускаем Сервер
	address := fmt.Sprintf("%s:%s", s.cfg.Host, s.cfg.Port)
	fmt.Printf("Сервер запущен на http://%s\n", address)

	return s.router.Run(address)
}

// Stop - остановка сервера
func (s *Server) Stop() error {
	fmt.Println("Сервер остановлен")
	return nil
}
