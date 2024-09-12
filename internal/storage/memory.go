package storage

import (
	"receipt-processor-challenge/internal/model"
	"sync"

	"github.com/google/uuid"
)

// InMemoryStorage is an in-memory storage for receipts
type InMemoryStorage struct {
	mu       sync.Mutex
	receipts map[uuid.UUID]model.Receipt
	mu2      sync.Mutex
	users    map[string]int
}

// New creates a new instance of InMemoryStorage
func New() *InMemoryStorage {
	return &InMemoryStorage{
		receipts: make(map[uuid.UUID]model.Receipt),
		users:    make(map[string]int),
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

// SaveUserHistory saves a receipt to the in-memory storage
func (s *InMemoryStorage) SaveUserHistory(user string) error {
	s.mu2.Lock()
	defer s.mu2.Unlock()
	if _, exists := s.users[user]; exists {
		s.users[user]++
	} else {
		s.users[user] = 0
	}
	return nil
}

// GetUserReceiptCount retrieves a users receipt count by user ID from the in-memory storage
func (s *InMemoryStorage) GetUserReceiptCount(user string) (int, bool) {
	s.mu2.Lock()
	defer s.mu2.Unlock()
	userIdReceiptCount, found := s.users[user]
	return userIdReceiptCount, found
}
