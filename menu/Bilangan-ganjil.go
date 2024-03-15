package menu

import (
	"basic/usecase"
	"fmt"
)

func BilanganGanjilMenu() {
	ganjilAwal := 1
	ganjilAkhir := 100
	fmt.Println("Deret Bilangan Ganjil dari", ganjilAwal, "hingga", ganjilAkhir)
	usecase.BilanganGanjil(ganjilAwal, ganjilAkhir)
}
