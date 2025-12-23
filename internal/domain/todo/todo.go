package todo

import (
	"todo-list-api-roadmap-sh/internal/domain/shared"

	"github.com/google/uuid"
)

type Todo struct {
	shared.BaseEntity
	Title       string    `gorm:"column:title,type:text"`
	Description string    `gorm:"column:description,type:text"`
	UserID      uuid.UUID `gorm:"column:user_id"`
}
