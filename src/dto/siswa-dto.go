package dto

type SiswaCreateRequest struct {
	IDSiswa            string `json:"id_siswa" binding:"required"`
	NamaSiswa          string `json:"nama_siswa" binding:"required"`
	Kelas              string `json:"kelas" binding:"required"`
	Password           string `json:"password" binding:"required"`
	KonfirmasiPassword string `json:"konfirmasi_password" binding:"required"`
}

type SiswaCreateResponse struct {
	IDSiswa   string `json:"id_siswa"`
	NamaSiswa string `json:"nama_siswa"`
	Kelas     string `json:"kelas"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
