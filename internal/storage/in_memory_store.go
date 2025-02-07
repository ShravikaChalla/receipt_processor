package storage

import (
	"receipt_processor/internal/models"
	"sync"
)

type InMemoryStore struct {
	data map[string]models.Receipt
	mu   sync.Mutex
}

func NewInMemoryStore() InMemoryStore {
	return InMemoryStore{data: make(map[string]models.Receipt)}
}

func (s *InMemoryStore) SaveReceipt(id string, receipt models.Receipt) {
	s.mu.Lock()
	s.data[id] = receipt
	s.mu.Unlock()
}

func (s *InMemoryStore) GetReceipt(id string) (models.Receipt, bool) {
	s.mu.Lock()
	receipt, exists := s.data[id]
	s.mu.Unlock()
	return receipt, exists
}
