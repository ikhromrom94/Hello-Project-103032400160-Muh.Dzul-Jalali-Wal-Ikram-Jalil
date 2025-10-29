package main

import "fmt"

type Pengguna struct {
	Nama       string
	Umur       int
	DaerahAsal string
	NoTelepon  string
	Email      string
	Pekerjaan  string
}

type Kamar struct {
	ID                   int
	Ukuran               float64
	JenisKelamin         string
	Fasilitas            int
	KonsepDesain         string
	Harga                float64
	HariMenujuPembayaran int
}

type AplikasiKostPintar struct {
	Pengguna []Pengguna
	Kamars   []Kamar
}

func main() {
	app := AplikasiKostPintar{
		Kamars: []Kamar{
			{ID: 1, Ukuran: 12, JenisKelamin: "pria", Fasilitas: 3, KonsepDesain: "minimalis", Harga: 1500000, HariMenujuPembayaran: 5},
			{ID: 2, Ukuran: 15, JenisKelamin: "wanita", Fasilitas: 4, KonsepDesain: "modern", Harga: 2000000, HariMenujuPembayaran: 10},
			{ID: 3, Ukuran: 10, JenisKelamin: "campur", Fasilitas: 2, KonsepDesain: "klasik", Harga: 1200000, HariMenujuPembayaran: 2},
		},
	}

	// Menu utama
	for {
		fmt.Println("\n=== Aplikasi Info Kost Pintar ===")
		fmt.Println("1. Tambah Data Pengguna")
		fmt.Println("2. Cari Kamar Kost")
		fmt.Println("3. Periksa Peringatan Pembayaran")
		fmt.Println("4. Urutkan Kamar")
		fmt.Println("5. Tampilkan Laporan Harga")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			app.tambahPengguna()
		case 2:
			app.cariKamar()
		case 3:
			app.peringatanPembayaran()
		case 4:
			app.urutkanKamar()
		case 5:
			app.tampilkanLaporanHarga()
		case 6:
			fmt.Println("Terima kasih telah menggunakan aplikasi.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func (a *AplikasiKostPintar) tambahPengguna() {
	var p Pengguna

	fmt.Println("\n--- Tambah Data Pengguna ---")
	fmt.Print("Nama: ")
	fmt.Scanln(&p.Nama)
	fmt.Print("Umur: ")
	fmt.Scanln(&p.Umur)
	fmt.Print("Daerah Asal: ")
	fmt.Scanln(&p.DaerahAsal)
	fmt.Print("Nomor Telepon: ")
	fmt.Scanln(&p.NoTelepon)
	fmt.Print("Email: ")
	fmt.Scanln(&p.Email)
	fmt.Print("Pekerjaan: ")
	fmt.Scanln(&p.Pekerjaan)

	a.Pengguna = append(a.Pengguna, p)
	fmt.Println("Data pengguna berhasil ditambahkan!")
}

func (a *AplikasiKostPintar) cariKamar() {
	if len(a.Kamars) == 0 {
		fmt.Println("\nBelum ada data kamar.")
		return
	}

	var (
		minUkuran    float64
		jenisKelamin string
		minFasilitas int
		konsepDesain string
	)

	fmt.Println("\n--- Pencarian Kamar Kost ---")
	fmt.Print("Ukuran minimal kamar (m²): ")
	fmt.Scanln(&minUkuran)
	fmt.Print("Jenis kelamin penghuni (pria/wanita/campur): ")
	fmt.Scanln(&jenisKelamin)
	fmt.Print("Tingkat fasilitas minimal (1-5): ")
	fmt.Scanln(&minFasilitas)
	fmt.Print("Konsep desain (kosongkan jika tidak penting): ")
	fmt.Scanln(&konsepDesain)

	fmt.Println("\nHasil Pencarian:")
	found := false

	for _, kamar := range a.Kamars {
		if kamar.Ukuran >= minUkuran &&
			(jenisKelamin == "" || kamar.JenisKelamin == jenisKelamin) &&
			kamar.Fasilitas >= minFasilitas &&
			(konsepDesain == "" || kamar.KonsepDesain == konsepDesain) {

			fmt.Printf("ID: %d, Ukuran: %.1f m², Jenis: %s, Fasilitas: %d/5, Desain: %s, Harga: Rp%.0f\n",
				kamar.ID, kamar.Ukuran, kamar.JenisKelamin, kamar.Fasilitas, kamar.KonsepDesain, kamar.Harga)
			found = true
		}
	}

	if !found {
		fmt.Println("Tidak ditemukan kamar yang sesuai dengan kriteria.")
	}
}

func (a *AplikasiKostPintar) peringatanPembayaran() {
	if len(a.Kamars) == 0 {
		fmt.Println("\nBelum ada data kamar.")
		return
	}

	fmt.Println("\n--- Peringatan Pembayaran ---")
	warningShown := false

	for _, kamar := range a.Kamars {
		if kamar.HariMenujuPembayaran <= 7 && kamar.HariMenujuPembayaran >= 0 {
			fmt.Printf("Kamar ID %d: Pembayaran jatuh tempo dalam %d hari\n",
				kamar.ID, kamar.HariMenujuPembayaran)
			warningShown = true
		}
	}

	if !warningShown {
		fmt.Println("Tidak ada pembayaran yang mendekati jatuh tempo.")
	}
}

func (a *AplikasiKostPintar) urutkanKamar() {
	if len(a.Kamars) == 0 {
		fmt.Println("\nBelum ada data kamar.")
		return
	}

	fmt.Println("\n--- Pengurutan Kamar Kost ---")
	fmt.Println("1. Berdasarkan Harga (Termurah)")
	fmt.Println("2. Berdasarkan Harga (Termahal)")
	fmt.Println("3. Berdasarkan Ukuran (Terbesar)")
	fmt.Println("4. Berdasarkan Fasilitas (Terbaik)")
	fmt.Print("Pilih kriteria pengurutan: ")

	var pilihan int
	fmt.Scanln(&pilihan)

	sortedKamars := make([]Kamar, len(a.Kamars))
	copy(sortedKamars, a.Kamars)

	for i := 0; i < len(sortedKamars)-1; i++ {
		for j := i + 1; j < len(sortedKamars); j++ {
			shouldSwap := false

			switch pilihan {
			case 1: // Harga termurah
				shouldSwap = sortedKamars[i].Harga > sortedKamars[j].Harga
			case 2: // Harga termahal
				shouldSwap = sortedKamars[i].Harga < sortedKamars[j].Harga
			case 3: // Ukuran terbesar
				shouldSwap = sortedKamars[i].Ukuran < sortedKamars[j].Ukuran
			case 4: // Fasilitas terbaik
				shouldSwap = sortedKamars[i].Fasilitas < sortedKamars[j].Fasilitas
			default:
				fmt.Println("Pilihan tidak valid.")
				return
			}

			if shouldSwap {
				sortedKamars[i], sortedKamars[j] = sortedKamars[j], sortedKamars[i]
			}
		}
	}

	fmt.Println("\nHasil Pengurutan:")
	for _, kamar := range sortedKamars {
		fmt.Printf("ID: %d, Ukuran: %.1f m², Jenis: %s, Fasilitas: %d/5, Harga: Rp%.0f\n",
			kamar.ID, kamar.Ukuran, kamar.JenisKelamin, kamar.Fasilitas, kamar.Harga)
	}
}

func (a *AplikasiKostPintar) tampilkanLaporanHarga() {
	if len(a.Kamars) == 0 {
		fmt.Println("\nBelum ada data kamar.")
		return
	}

	var total, rataRata, min, max float64
	min = a.Kamars[0].Harga
	max = a.Kamars[0].Harga

	for _, kamar := range a.Kamars {
		total += kamar.Harga
		if kamar.Harga < min {
			min = kamar.Harga
		}
		if kamar.Harga > max {
			max = kamar.Harga
		}
	}

	rataRata = total / float64(len(a.Kamars))

	fmt.Println("\n--- Laporan Harga Kamar ---")
	fmt.Printf("Jumlah Kamar Tersedia: %d\n", len(a.Kamars))
	fmt.Printf("Harga Terendah: Rp%.0f\n", min)
	fmt.Printf("Harga Tertinggi: Rp%.0f\n", max)
	fmt.Printf("Harga Rata-rata: Rp%.0f\n", rataRata)
	fmt.Println("\nDaftar Harga Kamar:")
	for _, kamar := range a.Kamars {
		fmt.Printf("ID: %d - Rp%.0f\n", kamar.ID, kamar.Harga)
	}
}
