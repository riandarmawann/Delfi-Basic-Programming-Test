package menu

import (
	"basic/usecase"
	"fmt"
)

func PerhitunganBarangMenu() {
	hasil := usecase.HitungBarang()
	fmt.Printf("Jumlah piring yang kini ada di tangan Vikrie: %d atau sekitar %d gross, %d lusin\n", hasil.JumlahPiring, hasil.JumlahGross, hasil.JumlahLusin)
}
