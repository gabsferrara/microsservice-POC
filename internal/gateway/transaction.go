package gateway

import "github.com/gabsferrara/microsservice-poc/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
