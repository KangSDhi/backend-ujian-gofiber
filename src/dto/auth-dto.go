package dto

type LoginRequest struct {
	EmailOrIDSiswa string `json:"email_or_id_siswa" validate:"required"`
	Password       string `json:"password" validate:"required"`
}
