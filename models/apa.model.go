package models

type Keluarga struct {
	Hubungan     string `json:"hubungan"`
	Nama         string `json:"nama"`
	TanggalLahir string `json:"tanggal_lahir"`
}

type Output struct {
	Keluarga        interface{} `json:"keluarga"`
	Nama            string      `json:"nama"`
	TanggalLahir    string      `json:"Tanggal_lahir"`
	Telpon          string      `json:"telpon"`
	Email           string      `json:"email"`
	Kewarganegaraan string      `json:"kewarganegaraan"`
}
