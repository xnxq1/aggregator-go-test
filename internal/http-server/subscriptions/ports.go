package http_server

import (
	"aggregator-go-test/internal/domain"
	"context"

	"github.com/google/uuid"
)

type SubscriptionServicePort interface {
	GetById(ctx context.Context, id uuid.UUID) (*domain.Subscription, error)
	Create(ctx context.Context, subscription *domain.Subscription) (*domain.Subscription, error)
}
