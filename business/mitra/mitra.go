package mitra

type Mitra struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Nama_toko string `json:"nama_toko"`
	Password  string `json:"password"`
	Alamat    string `json:"alamat"`
}

type AuthMitra struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type MitraRegister struct {
	Email     string `json:"email"`
	Nama_toko string `json:"nama_toko"`
	Password  string `json:"password"`
	Alamat    string `json:"alamat"`
}
