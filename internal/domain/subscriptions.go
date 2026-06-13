package domain

import "github.com/google/uuid"

type Subscription struct {
	ServiceName string
	Price       int
	UserId      uuid.UUID
	StartDate   string
	EndDate     *string
}
