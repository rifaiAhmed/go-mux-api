package gomuxService

import (
	"github.com/rifaiAhmed/go-rest-api-mux/models"
	"github.com/rifaiAhmed/go-rest-api-mux/repositories/gomuxRepository"
)

func GetData() models.Output {
	var res models.Output

	data := gomuxRepository.GetData()

	res.Keluarga = data
	res.Kewarganegaraan = "Indonesia (ID)"
	res.Nama = "Test1"
	res.TanggalLahir = "02-10-1997"
	res.Telpon = "085709471660"
	res.Email = "email@gmail.com"

	return res
}
