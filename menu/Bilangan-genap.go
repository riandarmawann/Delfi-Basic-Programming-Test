package menu

import (
	"basic/usecase"
	"fmt"
)

func BilanganGenapMenu() {
	genapAwal := 1
	genapAkhir := 100
	fmt.Println("Deret Bilangan Genap dari", genapAwal, "hingga", genapAkhir)
	usecase.BilanganGenap(genapAwal, genapAkhir)
}
