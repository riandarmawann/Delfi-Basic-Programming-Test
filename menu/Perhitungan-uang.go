package menu

import (
	"basic/usecase"
	"fmt"
)

func PerhitunganUangMenu() {
	var hargaBarang float64
	var pembayaran float64
	fmt.Println("Masukkan harga barang: ")
	fmt.Scanln(&hargaBarang)

	fmt.Println("Masukkan harga bayar: ")
	fmt.Scanln(&pembayaran)

	uangKembalian := usecase.PerhitunganUang(hargaBarang, pembayaran)
	fmt.Printf("Uang kembalian yang diterima user adalah Rp%.3f\n", uangKembalian.UangKembali)
}
