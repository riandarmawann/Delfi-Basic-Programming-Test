package usecase

import (
	"basic/models"
	"fmt"
)

func BilanganPrima(awal, akhir int) models.Prima {
	for i := awal; i <= akhir; i++ {
		if IsPrima(i) {
			fmt.Println(i)
		}
	}
	return models.Prima{}
}

func IsPrima(angka int) bool {
	if angka <= 1 {
		return false
	}

	for i := 2; i <= angka/2; i++ {
		if angka%i == 0 {
			return false
		}
	}

	return true
}
