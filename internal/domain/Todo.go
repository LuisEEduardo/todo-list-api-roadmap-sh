package domain

import "github.com/google/uuid"

type Todo struct {
	ID          uuid.UUID
	Title       string
	Description string
	UserID      uuid.UUID
}
