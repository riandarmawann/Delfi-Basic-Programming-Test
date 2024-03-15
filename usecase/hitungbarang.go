package usecase

import "basic/models"

func HitungBarang() models.HitungBarang {
	jumlahPiringAwal := 6 * 144    // (total 6 gross = 864)
	piringDipinjamTio := 4 * 12    // Tio meminjam 4 lusin (total 4 lusin = 48)
	piringDipinjamDavid := 2 * 144 // David meminjam 2 gross (total 2 gross = 288)

	jumlahPiringSisa := jumlahPiringAwal - piringDipinjamTio - piringDipinjamDavid

	jumlahGross := jumlahPiringSisa / 144
	sisaPiring := jumlahPiringSisa % 144
	jumlahLusin := sisaPiring / 12

	return models.HitungBarang{
		JumlahPiring: jumlahPiringSisa,
		JumlahGross:  jumlahGross,
		JumlahLusin:  jumlahLusin,
	}
}
