package domain

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name        string    `gorm:"not null"`
	Description string
	Price       float64   `gorm:"not null;type:numeric"`
	Stock       int       `gorm:"not null;default:0"`
	CreatedAt   time.Time `gorm:"default:current_timestamp"`
}

func NewProduct(name, description string, price float64, stock int) (*Product, error) {
	if name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}
	if price < 0 {
		return nil, fmt.Errorf("price cannot be negative")
	}

	return &Product{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
		CreatedAt:   time.Now(),
	}, nil
}
