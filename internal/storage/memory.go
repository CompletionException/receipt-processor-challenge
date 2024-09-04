package storage

import (
	"receipt-processor-challenge/internal/model"
	"github.com/google/uuid"
	"sync"
)

// InMemoryStorage is an in-memory storage for receipts
type InMemoryStorage struct {
	mu       sync.Mutex
	receipts map[uuid.UUID]model.Receipt
}

// New creates a new instance of InMemoryStorage
func New() *InMemoryStorage {
	return &InMemoryStorage{
		receipts: make(map[uuid.UUID]model.Receipt),
	}
}

// SaveReceipt saves a receipt to the in-memory storage
func (s *InMemoryStorage) SaveReceipt(receipt model.Receipt) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.receipts[receipt.ID] = receipt
	return nil
}

// GetReceipt retrieves a receipt by its ID from the in-memory storage
func (s *InMemoryStorage) GetReceipt(id uuid.UUID) (model.Receipt, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	receipt, found := s.receipts[id]
	return receipt, found
}
