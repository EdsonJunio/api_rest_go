package repository

import (
	"api_rest_go/internal/model"
	"fmt"

	"sync"
)

type InMemoryUserRepository struct {
	mu     sync.RWMutex
	data   map[int]model.User
	nextID int
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		data:   make(map[int]model.User),
		nextID: 1,
	}
}

func (r *InMemoryUserRepository) GetAll() ([]model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	users := make([]model.User, 0, len(r.data))
	for _, u := range r.data {
		users = append(users, u)
	}
	return users, nil
}

func (r *InMemoryUserRepository) GetByID(id int) (model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	u, exists := r.data[id]
	if !exists {
		return model.User{}, fmt.Errorf("user with id %d not found", id)
	}
	return u, nil
}

func (r *InMemoryUserRepository) Create(user model.User) (model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user.ID = r.nextID
	r.data[user.ID] = user
	r.nextID++
	return user, nil
}

func (r *InMemoryUserRepository) Update(user model.User) (model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.data[user.ID]; !exists {
		return model.User{}, fmt.Errorf("user with id %d not found", user.ID)
	}
	r.data[user.ID] = user
	return user, nil
}

func (r *InMemoryUserRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.data[id]; !exists {
		return fmt.Errorf("user with id %d not found", id)
	}
	delete(r.data, id)
	return nil
}
