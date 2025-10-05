package server

import (
	"auth/internal/config"
	"fmt"
)

type Server struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config) (*Server, error) {
	// Проверим, что конфиг не пустой
	if cfg == nil {
		return nil, fmt.Errorf("конфигурация сервера не может быть nil")
	}

	// Раз все окей, то создадим сервер
	return &Server{
		cfg: cfg,
	}, nil
}

func (s *Server) Serve() error {
	// Запускаем Сервер
	address := fmt.Sprintf("%s:%s", s.cfg.Host, s.cfg.Port)
	fmt.Printf("Сервер запущен на %s...\n", address)
	return nil
}
