package postgre

import (
	"github.com/google/uuid"
	"github.com/guregu/null"
	"time"
)

type BaseModel struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
	DeletedAt null.Time `json:"deletedAt" db:"deleted_at"`
}

func NewBaseModel() *BaseModel {
	now := time.Now()

	return &BaseModel{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
	}
}

var BaseModelFields = []string{"id", "created_at", "updated_at", "deleted_at"}
