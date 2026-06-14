package http_server

import "github.com/google/uuid"

type CreateSubscriptionRequest struct {
	ServiceName string    `json:"service_name" validate:"required"                    example:"Netflix"`
	Price       int       `json:"price"        validate:"required,gt=0"               example:"499"`
	UserId      uuid.UUID `json:"user_id"      validate:"required"                    example:"550e8400-e29b-41d4-a716-446655440000"`
	StartDate   string    `json:"start_date"   validate:"required,datetime=01-2006"   example:"01-2026"`
	EndDate     *string   `json:"end_date,omitempty" validate:"omitempty,datetime=01-2006" example:"12-2026"`
}
