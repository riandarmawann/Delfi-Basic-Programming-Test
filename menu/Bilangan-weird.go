package menu

import (
	"basic/usecase"
	"fmt"
	"log"
)

func BilanganWeirdMenu() {
	var N int
	fmt.Println("Masukkan angka (1 - 100): ")
	_, err := fmt.Scanf("%d\n", &N)
	if err != nil {
		log.Println("Error membaca inputan:", err)
		return
	}

	if N < 1 || N > 100 {
		fmt.Println("Angka tidak valid!")
		return
	}

	usecase.Bilanganweird(N)
}
