package menu

import (
	"errors"
)

// Sentinel error for missing dishes
var ErrNotFound = errors.New("dish not found")

// Dish represents a menu item with snake_case JSON tags
type Dish struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	PriceSantim int    `json:"price_santim"` // Kept as integer to avoid floating-point errors
	IsFasting   bool   `json:"is_fasting"`
	Note        string `json:"note,omitempty"` // Omitted if empty
}

// DishReader defines read capabilities for menu items
type DishReader interface {
	All() ([]Dish, error)
	ByID(id string) (Dish, error)
}






