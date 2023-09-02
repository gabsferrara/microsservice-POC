package gateway

import "github.com/gabsferrara/microsservice-poc/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}
