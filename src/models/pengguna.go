package models

import (
	"backend-ujian-gofiber/src/database"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RolePengguna string

const (
	Admin RolePengguna = "admin"
	Guru  RolePengguna = "guru"
	Siswa RolePengguna = "siswa"
)

type Pengguna struct {
	gorm.Model
	ID            uuid.UUID    `gorm:"primaryKey;type:uuid;column:id;not null"`
	IdSiswa       *string      `gorm:"uniqueIndex;column:id_siswa"`
	NamaPengguna  string       `gorm:"uniqueIndex;column:nama_pengguna;not null"`
	EmailPengguna *string      `gorm:"uniqueIndex;column:email_pengguna"`
	Password      string       `gorm:"column:password;not null"`
	PasswordPlain *string      `gorm:"column:password_plain;"`
	RolePengguna  RolePengguna `gorm:"type:enum('admin', 'guru', 'siswa');column:role_pengguna;not null"`
	KelasID       *uuid.UUID   `gorm:"column:kelas_id;type:uuid"`
	Kelas         *Kelas       `gorm:"foreignKey:KelasID;references:ID"`
}

func (Pengguna) TableName() string {
	return "pengguna"
}

func FindPenggunaByEmail(email string) (Pengguna, error) {
	var pengguna Pengguna
	err := database.DB.Where("email_pengguna = ?", email).First(&pengguna).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Pengguna{}, errors.New("Admin Tidak Ditemukan!")
		}
		return Pengguna{}, err
	}
	return pengguna, nil
}

func FindPenggunaByIDSiswa(idSiswa string) (Pengguna, error) {
	var pengguna Pengguna
	err := database.DB.Where("id_siswa = ?", idSiswa).First(&pengguna).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Pengguna{}, errors.New("Siswa Tidak Ditemukan!")
		}
		return Pengguna{}, err
	}
	return pengguna, nil
}
