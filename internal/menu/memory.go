package menu

import (
	"sync"
)

type memoryStore struct {
	mu     sync.RWMutex
	dishes []Dish
}

// NewMemoryStore seeds the eight initial Addis Eats dishes
func NewMemoryStore() *memoryStore {
	return &memoryStore{
		dishes: []Dish{
			{ID: "d1", Name: "Shiro Wot", PriceSantim: 12000, IsFasting: true, Note: "Classic chickpea stew"},
			{ID: "d2", Name: "Beef Tibs", PriceSantim: 25000, IsFasting: false, Note: "Sautéed beef with onions and peppers"},
			{ID: "d3", Name: "Beyaynetu", PriceSantim: 18000, IsFasting: true, Note: "Combination vegetarian platter"},
			{ID: "d4", Name: "Special Kitfo", PriceSantim: 32000, IsFasting: false, Note: "Minced beef with mitmita and niter kibbeh"},
			{ID: "d5", Name: "Doro Wot", PriceSantim: 35000, IsFasting: false, Note: "Traditional chicken stew with boiled egg"},
			{ID: "d6", Name: "Asa Tibs", PriceSantim: 22000, IsFasting: true, Note: "Fried fish cutlets"},
			{ID: "d7", Name: "Firfir", PriceSantim: 15000, IsFasting: true, Note: "Shredded injera in spiced sauce"},
			{ID: "d8", Name: "Gomen besiga", PriceSantim: 20000, IsFasting: false, Note: "Collard greens cooked with beef"},
		},
	}
}

func (s *memoryStore) All() ([]Dish, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Return a copy to prevent concurrent mutation
	dst := make([]Dish, len(s.dishes))
	copy(dst, s.dishes)
	return dst, nil
}

func (s *memoryStore) ByID(id string) (Dish, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, d := range s.dishes {
		if d.ID == id {
			return d, nil
		}
	}
	return Dish{}, ErrNotFound
}