package tokenmngr

import (
	"context"

	"github.com/shellhub-io/shellhub/api/pkg/store"
	"github.com/shellhub-io/shellhub/pkg/models"
)

type Service interface {
	CreateToken(ctx context.Context, token models.Token) (*models.Token, error)
	GetToken(ctx context.Context, id string) (*models.Token, error)
}

type service struct {
	store store.Store
}

func NewService(store store.Store) Service {
	return &service{store}
}

func (s *service) CreateToken(ctx context.Context, token models.Token) (*models.Token, error) {
	return s.store.CreateToken(ctx, token)

}

func (s *service) GetToken(ctx context.Context, id string) (*models.Token, error) {
	return s.store.GetToken(ctx, id)
}
