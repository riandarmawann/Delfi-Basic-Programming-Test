package usecase

import (
	"basic/models"
)

func Simulasibunga(tabungan float64) models.BungaBank {
	bunga := tabungan * 5 / 100
	return models.BungaBank{
		Bunga: bunga,
	}
}
