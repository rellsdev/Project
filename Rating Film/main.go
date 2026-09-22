package main

import "fmt"

// ================= KONSTANTA & STRUCT =================

const MAX_FILM = 100

type Film struct {
	Judul    string
	Genre    string
	Rating   float64
	Favorite bool
}

var daftarFilm [MAX_FILM]Film
var jumlahFilm int = 0

// ================= HELPER TEXT =================

// Mengubah huruf besar menjadi huruf kecil
// Pengganti strings.ToLower()
func toLower(teks string) string {
	hasil := ""

	for _, karakter := range teks {
		if karakter >= 'A' && karakter <= 'Z' {
			karakter = karakter + ('a' - 'A')
		}

		hasil += string(karakter)
	}

	return hasil
}

// Membaca input teks.
// Bisa menerima beberapa kata, misalnya:
// Avengers: Endgame
// The Lord of the Rings
func bacaString(prompt string) string {
	fmt.Print(prompt)

	var kata1, kata2, kata3, kata4, kata5 string
	var kata6, kata7, kata8, kata9, kata10 string

	jumlah, _ := fmt.Scanln(
		&kata1,
		&kata2,
		&kata3,
		&kata4,
		&kata5,
		&kata6,
		&kata7,
		&kata8,
		&kata9,
		&kata10,
	)

	hasil := ""

	if jumlah >= 1 {
		hasil += kata1
	}
	if jumlah >= 2 {
		hasil += " " + kata2
	}
	if jumlah >= 3 {
		hasil += " " + kata3
	}
	if jumlah >= 4 {
		hasil += " " + kata4
	}
	if jumlah >= 5 {
		hasil += " " + kata5
	}
	if jumlah >= 6 {
		hasil += " " + kata6
	}
	if jumlah >= 7 {
		hasil += " " + kata7
	}
	if jumlah >= 8 {
		hasil += " " + kata8
	}
	if jumlah >= 9 {
		hasil += " " + kata9
	}
	if jumlah >= 10 {
		hasil += " " + kata10
	}

	return hasil
}

func bacaInt(prompt string) int {
	for {
		fmt.Print(prompt)

		var teks string
		fmt.Scanln(&teks)

		var nilai int
		_, err := fmt.Sscanf(teks, "%d", &nilai)

		if err == nil {
			return nilai
		}

		fmt.Println("Input tidak valid! Masukkan angka.")
	}
}

func bacaFloat(prompt string) float64 {
	for {
		fmt.Print(prompt)

		var teks string
		fmt.Scanln(&teks)

		var nilai float64
		_, err := fmt.Sscanf(teks, "%f", &nilai)

		if err == nil && nilai >= 0 && nilai <= 10 {
			return nilai
		}

		fmt.Println("Input tidak valid! Masukkan angka 0-10.")
	}
}

// ================= PROCEDURE: INPUT / TAMPIL =================
func tambahFilm(judul string, genre string, rating float64) {

	if jumlahFilm >= MAX_FILM {
		fmt.Println("Gagal! Daftar film sudah penuh.")
		return
	}

	daftarFilm[jumlahFilm] = Film{
		Judul:    judul,
		Genre:    genre,
		Rating:   rating,
		Favorite: false,
	}

	jumlahFilm++

	fmt.Println("Film berhasil ditambahkan!")
}

func tampilkanSemuaFilm() bool {

	if jumlahFilm == 0 {
		fmt.Println("Belum ada data film.")
		return false
	}

	fmt.Println("\n=============================== DAFTAR FILM ===============================")
	fmt.Printf("%-4s %-25s %-15s %-8s %-8s\n",
		"No", "Judul", "Genre", "Rating", "Favorite")

	fmt.Println("---------------------------------------------------------------------------------")

	for i := 0; i < jumlahFilm; i++ {

		f := daftarFilm[i]

		tandaFav := "-"

		if f.Favorite {
			tandaFav = "Ya"
		}

		fmt.Printf(
			"%-4d %-25s %-15s %-8.1f %-8s\n",
			i+1,
			f.Judul,
			f.Genre,
			f.Rating,
			tandaFav,
		)
	}

	fmt.Println("---------------------------------------------------------------------------------")

	return true
}

// ================= EDIT / HAPUS =================
func hapusFilm(index int) bool {

	if index < 0 || index >= jumlahFilm {
		return false
	}

	for i := index; i < jumlahFilm-1; i++ {
		daftarFilm[i] = daftarFilm[i+1]
	}

	jumlahFilm--

	return true
}

func editRatingFilm(index int, ratingBaru float64) bool {

	if index < 0 || index >= jumlahFilm {
		return false
	}

	daftarFilm[index].Rating = ratingBaru

	return true
}

func editGenreFilm(index int, genreBaru string) bool {

	if index < 0 || index >= jumlahFilm {
		return false
	}

	daftarFilm[index].Genre = genreBaru

	return true
}

func toggleFavoriteFilm(index int) bool {

	if index < 0 || index >= jumlahFilm {
		return false
	}

	daftarFilm[index].Favorite = !daftarFilm[index].Favorite

	return true
}

func cariFilmLinear(judul string) int {

	judul = toLower(judul)

	for i := 0; i < jumlahFilm; i++ {

		if toLower(daftarFilm[i].Judul) == judul {
			return i
		}
	}

	return -1
}

func cariFilmBinary(judul string) int {

	kiri := 0
	kanan := jumlahFilm - 1

	judul = toLower(judul)

	for kiri <= kanan {

		tengah := (kiri + kanan) / 2

		judulTengah := toLower(daftarFilm[tengah].Judul)

		if judulTengah == judul {
			return tengah

		} else if judulTengah < judul {
			kiri = tengah + 1

		} else {
			kanan = tengah - 1
		}
	}

	return -1
}

func bubbleSortRatingDesc() {

	for i := 0; i < jumlahFilm-1; i++ {

		for j := 0; j < jumlahFilm-1-i; j++ {

			if daftarFilm[j].Rating < daftarFilm[j+1].Rating {

				daftarFilm[j], daftarFilm[j+1] =
					daftarFilm[j+1], daftarFilm[j]
			}
		}
	}
}

func selectionSortJudulAsc() {

	for i := 0; i < jumlahFilm-1; i++ {

		idxMin := i

		for j := i + 1; j < jumlahFilm; j++ {

			if toLower(daftarFilm[j].Judul) <
				toLower(daftarFilm[idxMin].Judul) {

				idxMin = j
			}
		}

		if idxMin != i {
			daftarFilm[i], daftarFilm[idxMin] =
				daftarFilm[idxMin], daftarFilm[i]
		}
	}
}

// ================= DATA AWAL =================

func isiDataAwal() {

	tambahFilm("Interstellar", "Sci-Fi", 8.7)
	tambahFilm("The Godfather", "Drama", 9.2)
	tambahFilm("Parasite", "Thriller", 8.6)
	tambahFilm("Toy Story", "Animation", 8.3)
	tambahFilm("Avengers: Endgame", "Action", 8.4)
}

// ================= MENU KELOLA FILM =================

func menuKelolaFilm() {

	ada := tampilkanSemuaFilm()

	if !ada {
		return
	}

	for {

		fmt.Println("\n--- KELOLA FILM ---")
		fmt.Println("1. Edit Rating")
		fmt.Println("2. Edit Genre")
		fmt.Println("3. Hapus Film")
		fmt.Println("4. Tandai / Batal Favorite")
		fmt.Println("0. Kembali ke Menu Utama")

		pilihan := bacaInt("Pilih aksi: ")

		switch pilihan {

		case 1:

			no := bacaInt("Masukkan nomor film yang ingin diedit ratingnya: ")
			idx := no - 1

			ratingBaru := bacaFloat("Rating baru (0-10): ")

			if editRatingFilm(idx, ratingBaru) {
				fmt.Println("Rating berhasil diperbarui!")
			} else {
				fmt.Println("Nomor tidak valid.")
			}

			tampilkanSemuaFilm()

		case 2:

			no := bacaInt("Masukkan nomor film yang ingin diedit genrenya: ")
			idx := no - 1

			genreBaru := bacaString("Genre baru: ")

			if editGenreFilm(idx, genreBaru) {
				fmt.Println("Genre berhasil diperbarui!")
			} else {
				fmt.Println("Nomor tidak valid.")
			}

			tampilkanSemuaFilm()

		case 3:

			no := bacaInt("Masukkan nomor film yang ingin dihapus: ")
			idx := no - 1

			if hapusFilm(idx) {
				fmt.Println("Film berhasil dihapus.")
			} else {
				fmt.Println("Nomor tidak valid.")
			}

			if !tampilkanSemuaFilm() {
				return
			}

		case 4:

			no := bacaInt(
				"Masukkan nomor film yang ingin ditandai/dibatalkan favorite: ",
			)

			idx := no - 1

			if toggleFavoriteFilm(idx) {

				if daftarFilm[idx].Favorite {
					fmt.Println("Film ditandai sebagai favorite!")
				} else {
					fmt.Println("Status favorite dibatalkan.")
				}

			} else {
				fmt.Println("Nomor tidak valid.")
			}

			tampilkanSemuaFilm()

		case 0:
			return

		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}

// ================= MENU UTAMA =================

func tampilkanMenu() {

	fmt.Println("\n========== MOVIE RATING SYSTEM ==========")
	fmt.Println("1. Tambah Film")
	fmt.Println("2. Tampilkan Semua Film (Edit / Hapus / Favorite)")
	fmt.Println("3. Cari Film (Linear Search)")
	fmt.Println("4. Cari Film (Binary Search - otomatis urutkan judul dulu)")
	fmt.Println("5. Urutkan Film Berdasarkan Rating (Bubble Sort)")
	fmt.Println("6. Urutkan Film Berdasarkan Judul (Selection Sort)")
	fmt.Println("0. Keluar")
	fmt.Println("==========================================")
}

// ================= MAIN =================

func main() {

	isiDataAwal()

	for {

		tampilkanMenu()

		pilihan := bacaInt("Pilih menu: ")

		switch pilihan {

		case 1:

			judul := bacaString("Judul film: ")
			genre := bacaString("Genre: ")
			rating := bacaFloat("Rating (0-10): ")

			tambahFilm(judul, genre, rating)

		case 2:

			menuKelolaFilm()

		case 3:

			judul := bacaString("Masukkan judul film yang dicari: ")

			idx := cariFilmLinear(judul)

			if idx == -1 {

				fmt.Println("Film tidak ditemukan.")

			} else {

				f := daftarFilm[idx]

				fmt.Printf(
					"Ditemukan! Judul: %s | Genre: %s | Rating: %.1f\n",
					f.Judul,
					f.Genre,
					f.Rating,
				)
			}

		case 4:

			selectionSortJudulAsc()

			fmt.Println("(Data diurutkan berdasarkan judul terlebih dahulu)")

			judul := bacaString("Masukkan judul film yang dicari: ")

			idx := cariFilmBinary(judul)

			if idx == -1 {

				fmt.Println("Film tidak ditemukan.")

			} else {

				f := daftarFilm[idx]

				fmt.Printf(
					"Ditemukan! Judul: %s | Genre: %s | Rating: %.1f\n",
					f.Judul,
					f.Genre,
					f.Rating,
				)
			}

		case 5:

			bubbleSortRatingDesc()

			fmt.Println(
				"Film berhasil diurutkan berdasarkan rating (tertinggi ke terendah).",
			)

			tampilkanSemuaFilm()

		case 6:

			selectionSortJudulAsc()

			fmt.Println(
				"Film berhasil diurutkan berdasarkan judul (A-Z).",
			)

			tampilkanSemuaFilm()

		case 0:

			fmt.Println(
				"Terima kasih sudah menggunakan Movie Rating System!",
			)

			return

		default:

			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}