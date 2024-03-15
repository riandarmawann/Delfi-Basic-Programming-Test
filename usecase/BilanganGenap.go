package usecase

import (
	"basic/models"
	"fmt"
)

func BilanganGenap(genapAwal, genapAkhir int) models.Genap {
	for i := genapAwal; i <= genapAkhir; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	return models.Genap{}
}
