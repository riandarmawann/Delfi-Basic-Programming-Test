package menu

import (
	"basic/usecase"
	"fmt"
)

func BilanganPrimaMenu() {
	primaAwal := 1
	primaAkhir := 100
	fmt.Println("Deret Bilangan Prima dari", primaAwal, "hingga", primaAkhir)
	usecase.BilanganPrima(primaAwal, primaAkhir)
}
