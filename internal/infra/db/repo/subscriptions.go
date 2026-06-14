package repo

import (
	"context"
	"errors"

	domain "aggregator-go-test/internal/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SubRepoAdapter struct {
	pool *pgxpool.Pool
}

var ErrSubscriptionNotFound = errors.New("subscription not found")

func (repo *SubRepoAdapter) GetById(ctx context.Context, id uuid.UUID) (*domain.Subscription, error) {
	var res domain.Subscription
	err := repo.pool.QueryRow(ctx, `
		SELECT id, service_name, price, user_id, start_date, end_date
		FROM subscriptions
		WHERE id = $1`, id,
	).Scan(&res.Id, &res.ServiceName, &res.Price, &res.UserId, &res.StartDate, &res.EndDate)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrSubscriptionNotFound
	}
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (repo *SubRepoAdapter) Create(ctx context.Context, subscription *domain.Subscription) (*domain.Subscription, error) {
	err := repo.pool.QueryRow(
		ctx,
		`INSERT INTO subscriptions (service_name, price, user_id, start_date, end_date)
			VALUES ($1, $2, $3, $4, $5)
			RETURNING id
`, subscription.ServiceName, subscription.Price, subscription.UserId, subscription.StartDate, subscription.EndDate,
	).Scan(&subscription.Id)
	if err != nil {
		return nil, err
	}
	return subscription, nil
}

func NewSubRepoAdapter(pool *pgxpool.Pool) *SubRepoAdapter {
	return &SubRepoAdapter{pool: pool}
}
