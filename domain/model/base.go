package model

import (
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Base struct {
	ID        string    `json:"id" valid:"uuid" gorm:"column:id; type:uuid; not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at; type:date;" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at: type:date;" valid:"-"`
}
