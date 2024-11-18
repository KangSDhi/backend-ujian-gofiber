package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Kelas struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;column:id;not null"`
	NamaKelas string    `gorm:"column:nama_kelas;not null"`
}
