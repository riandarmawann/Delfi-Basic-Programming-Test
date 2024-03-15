package usecase

import "basic/models"

func PerhitunganUang(hargaBarang, pembayaran float64) models.HitungUang {
	diskon := 0.15 // Diskon 15%
	hargaDiskon := hargaBarang * diskon
	hargaBayar := hargaBarang - hargaDiskon
	uangKembali := pembayaran - hargaBayar

	return models.HitungUang{
		UangKembali: uangKembali,
	}
}
