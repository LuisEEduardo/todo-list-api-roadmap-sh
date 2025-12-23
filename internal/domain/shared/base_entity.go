package shared

import (
	"time"

	"github.com/google/uuid"
)

type BaseEntity struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time  `gorm:"autoCreateTime;"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime;"`
	DeletedAt *time.Time `gorm:"autoDeleteTime;"`
}
