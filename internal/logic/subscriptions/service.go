package logic

import (
	"aggregator-go-test/internal/domain"
	"context"

	"github.com/google/uuid"
)

type SubscriptionService struct {
	Repo SubscriptionRepoPort
}

func (src *SubscriptionService) GetById(ctx context.Context, id uuid.UUID) (*domain.Subscription, error) {
	return src.Repo.GetById(ctx, id)
}

func (src *SubscriptionService) Create(ctx context.Context, subscription *domain.Subscription) (*domain.Subscription, error) {
	return src.Repo.Create(ctx, subscription)
}

func NewService(repo SubscriptionRepoPort) *SubscriptionService {
	return &SubscriptionService{Repo: repo}
}
