package repository

import (
	"sync"

	"siem-sistem/internal/model"
)

type Repository struct {
	mu    sync.Mutex
	items map[int]model.Item
}

// Конструктор без аргументов
func NewRepository() *Repository {
	return &Repository{
		items: make(map[int]model.Item),
	}
}

// Метод Load (если он нужен)
func (r *Repository) Load() error {
	// Заглушка, если этот метод должен загружать данные из файлов
	return nil
}

func (r *Repository) GetAllItems() []model.Item {
	r.mu.Lock()
	defer r.mu.Unlock()

	var result []model.Item
	for _, item := range r.items {
		result = append(result, item)
	}
	return result
}

func (r *Repository) GetItemByID(id int) (model.Item, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	item, found := r.items[id]
	return item, found
}

func (r *Repository) AddItem(item model.Item) int {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := len(r.items) + 1
	item.ID = id
	r.items[id] = item
	return id
}

func (r *Repository) UpdateItem(id int, updatedItem model.Item) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, found := r.items[id]; !found {
		return false
	}

	updatedItem.ID = id
	r.items[id] = updatedItem
	return true
}

func (r *Repository) DeleteItem(id int) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, found := r.items[id]; !found {
		return false
	}

	delete(r.items, id)
	return true
}
