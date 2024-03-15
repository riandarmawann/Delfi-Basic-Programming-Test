package menu

import (
	"fmt"
	"log"
)

func Run() {
	showMenu()
}

var pilihan int

var Exit bool

func showMenu() {
	for {
		if Exit {
			break
		}

		fmt.Println("------Delfi Test------")
		fmt.Println("1. Bilangan Prima")
		fmt.Println("2. Bilangan Ganjil")
		fmt.Println("3. Bilangan Genap")
		fmt.Println("4. Bilangan Fibonacci")
		fmt.Println("5. Simulasi Bunga Bank")
		fmt.Println("6. Perhitungan Uang Kembali")
		fmt.Println("7. Perhitungan Jumlah Barang")
		fmt.Println("8. Bilangan Weird")
		fmt.Println("9. Keluar aplikasi")

		fmt.Print("Masukkan pilihan menu: ")
		_, err := fmt.Scanf("%d\n", &pilihan)
		if err != nil {
			log.Println("Error membaca inputan:", err)
			return
		}

		switch pilihan {
		case 1:
			BilanganPrimaMenu()
		case 2:
			BilanganGanjilMenu()
		case 3:
			BilanganGenapMenu()
		case 4:
			BilanganFibonacciMenu()
		case 5:
			SimulasiBungaBankMenu()
		case 6:
			PerhitunganUangMenu()
		case 7:
			PerhitunganBarangMenu()
		case 8:
			BilanganWeirdMenu()
		case 9:
			exitConfirmation()
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func exitConfirmation() {
	fmt.Print("Anda yakin ingin keluar dari aplikasi? Ya/Tidak: ")
	var confirm string
	_, err := fmt.Scanf("%s\n", &confirm)
	if err != nil {
		log.Println("Error reading input:", err)
		return
	}

	if confirm == "Ya" || confirm == "ya" {
		fmt.Println("Anda telah keluar dari aplikasi")
		Exit = true
	} else {
		fmt.Println("Batal keluar dari aplikasi")
	}
}
