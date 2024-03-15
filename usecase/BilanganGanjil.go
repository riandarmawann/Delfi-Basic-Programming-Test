package usecase

import (
	"basic/models"
	"fmt"
)

func BilanganGanjil(ganjilAwal, ganjilAkhir int) models.Ganjil {
	for i := ganjilAwal; i <= ganjilAkhir; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
	return models.Ganjil{}
}
