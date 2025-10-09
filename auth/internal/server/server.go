package server

import (
	"auth/internal/config"
	"auth/internal/handler"
	"fmt"
)

// У серва должны быть
// Конфигурации
// Обработчики

type Server struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config) (*Server, error) {
	// Проверим, что конфиг не пустой
	if cfg == nil {
		return nil, fmt.Errorf("Конфигурация сервера не может быть nil")
	}

	// Новый экземпляр обработчика
	handler := handler.NewHandler(cfg)
	if handler == nil {
		return nil, fmt.Errorf("Не удалось создать обработчик сервера")
	}
	fmt.Println("Обработчика сервера успешно создан")

	// Раз все окей, то создадим сервер
	return &Server{
		cfg: cfg,
	}, nil
}

func (s *Server) Serve() error {
	// Запускаем Сервер
	address := fmt.Sprintf("%s:%s", s.cfg.Host, s.cfg.Port)
	fmt.Printf("Сервер запущен на %s\n", address)
	return nil
}
