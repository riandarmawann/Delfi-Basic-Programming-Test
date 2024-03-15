package menu

import (
	"basic/usecase"
	"fmt"
	"log"
)

func SimulasiBungaBankMenu() {
	var tabungan float64
	fmt.Println("Masukkan jumlah tabungan: ")
	_, err := fmt.Scanf("%f\n", &tabungan)
	if err != nil {
		log.Println("Error membaca inputan:", err)
		return
	}

	if tabungan <= 0 {
		fmt.Println("Jumlah tabungan tidak valid!")
		return
	}

	bunga := usecase.Simulasibunga(tabungan)
	fmt.Printf("Bunga bank selama 1 tahun: Rp%.3f\n", bunga.Bunga)
}
