package services

import (
	"receipt_processor/internal/models"
	"receipt_processor/internal/storage"
	"receipt_processor/pkg/utils"
)

type ReceiptService struct {
	store storage.InMemoryStore
}

func NewReceiptService() *ReceiptService {
	return &ReceiptService{store: storage.NewInMemoryStore()}
}

func (s *ReceiptService) ProcessReceipt(receipt models.Receipt) string {
	id := utils.GenerateUUID()
	s.store.SaveReceipt(id, receipt)
	return id
}

func (s *ReceiptService) GetPoints(id string) (int, bool) {
	receipt, exists := s.store.GetReceipt(id)
	if !exists {
		return 0, false
	}
	return utils.CalculatePoints(receipt), true
}
