package model

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Base struct {
	ID        string    `json:"id" valid:"uuid" gorm:"column:id; type:uuid; not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at; type:date; not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at: type:date;"`
}
