package gomuxRepository

import "github.com/rifaiAhmed/go-rest-api-mux/models"

func GetData() []models.Keluarga {
	var obj = []models.Keluarga{
		models.Keluarga{Hubungan: "Istri1", Nama: "Nama_Istri", TanggalLahir: "10-01-1998"},
		models.Keluarga{Hubungan: "Anak_ke1", Nama: "Nama_Anak", TanggalLahir: "04-16-2025"},
	}

	return obj
}
