package database

import (
	"auth/internal/config"
	"context"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDatabase - функция для создания нового ПОДКЛЮЧЕНИИЯ к БД.
// Принимает: контекст, DSN (строку подключения)
// Возращает: указатель на gorm.DB, ошибку
func NewDatabase(cfg *config.Config, models ...any) (*gorm.DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.Timeout)*time.Second)
	defer cancel()

	// На всякий сделаем задержку, тк иногда,
	// например, если БД запускается в
	// Docker-контейнере, то для запуска и готвности
	// к соединение может потребоваться времяю
	time.Sleep(1 * time.Second)

	db, err := gorm.Open(postgres.Open(cfg.DBDSN), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("%s: %v", "ошибка подключения к базе данных", err)
	}

	errMigration := runMigrations(db, models...)
	if errMigration != nil {
		return nil, fmt.Errorf("%s: %v", "ошибка миграции БД", errMigration)
	}

	return db.WithContext(ctx), nil
}

// Миграции (автоматические). А что такое миграция?
// По сути это перенос структуры в БД. Была структура User,
// мы ее передали в параметр - она стала сущностью в БД.
func runMigrations(db *gorm.DB, models ...any) error {
	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			return fmt.Errorf("ошибка миграции модели %T: %w", model, err)
		}
	}

	fmt.Println("Все миграции успешно выполнены")
	return nil
}
