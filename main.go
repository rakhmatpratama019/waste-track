package main

import "fmt"

const NMAX int = 100

type Warga struct {
	ID           int
	Nama         string
	Alamat       string
	JenisSampah  string
	BeratSampah  float64
	TanggalSetor string
}

var dataWarga [NMAX]Warga
var jumlahData int

func tampilkanHeader() {

	fmt.Println("===================================================")
	fmt.Println("         WASTE-TRACK MANAGEMENT SYSTEM")
	fmt.Println("===================================================")
	fmt.Println("Kelompok:")
	fmt.Println("1. Rakhmat Pratama")
	fmt.Println("2. Ahmad Luthfi Habibie")
	fmt.Println("3. Muhammad Haidar Az Zacky")
	fmt.Println("===================================================")
	fmt.Println()
}

func tampilkanMenu() {

	fmt.Println()
	fmt.Println("============== MENU UTAMA ==============")
	fmt.Println("1. Tambah Data Warga")
	fmt.Println("2. Ubah Data Warga")
	fmt.Println("3. Hapus Data Warga")
	fmt.Println("4. Tampilkan Data Warga")
	fmt.Println("5. Tambah Transaksi Sampah")
	fmt.Println("6. Searching Data")
	fmt.Println("7. Sorting Data")
	fmt.Println("8. Statistik Sampah")
	fmt.Println("9. Keluar")
	fmt.Println("========================================")
}

func sequentialSearchID(id int) int {

	var i int
	var found bool

	i = 0
	found = false

	for i < jumlahData && !found {

		if dataWarga[i].ID == id {
			found = true
		} else {
			i++
		}
	}

	if found {
		return i
	}

	return -1
}

func tambahWarga() {

	if jumlahData >= NMAX {

		fmt.Println("⚠️ Data sudah penuh!")
		return
	}

	var wargaBaru Warga

	fmt.Println()
	fmt.Println("========== TAMBAH DATA WARGA ==========")

	// ID otomatis
	wargaBaru.ID = jumlahData + 1

	fmt.Println("ID Warga Otomatis      :", wargaBaru.ID)

	fmt.Print("Masukkan Nama Warga    : ")
	fmt.Scan(&wargaBaru.Nama)

	if wargaBaru.Nama == "" {

		fmt.Println("⚠️ Nama tidak boleh kosong!")
		return
	}

	fmt.Print("Masukkan Alamat        : ")
	fmt.Scan(&wargaBaru.Alamat)

	if wargaBaru.Alamat == "" {

		fmt.Println("⚠️ Alamat tidak boleh kosong!")
		return
	}

	dataWarga[jumlahData] = wargaBaru
	jumlahData++

	fmt.Println("✅ Data warga berhasil ditambahkan!")
}

func tampilkanWarga() {

	if jumlahData == 0 {

		fmt.Println("⚠️ Belum ada data warga!")
		return
	}

	fmt.Println()
	fmt.Println("=============== DATA WARGA ===============")

	for i := 0; i < jumlahData; i++ {

		fmt.Println("------------------------------------------")
		fmt.Println("ID              :", dataWarga[i].ID)
		fmt.Println("Nama            :", dataWarga[i].Nama)
		fmt.Println("Alamat          :", dataWarga[i].Alamat)

		if dataWarga[i].JenisSampah == "" {

			fmt.Println("Jenis Sampah    : Belum ada transaksi")
			fmt.Println("Berat Sampah    : 0 KG")
			fmt.Println("Tanggal Setor   : -")

		} else {

			fmt.Println("Jenis Sampah    :", dataWarga[i].JenisSampah)
			fmt.Println("Berat Sampah    :", dataWarga[i].BeratSampah, "KG")
			fmt.Println("Tanggal Setor   :", dataWarga[i].TanggalSetor)
		}
	}

	fmt.Println("------------------------------------------")
}

func ubahWarga() {

	var id int

	fmt.Print("Masukkan ID warga yang ingin diubah : ")
	fmt.Scan(&id)

	index := sequentialSearchID(id)

	if index == -1 {

		fmt.Println("⚠️ Data warga tidak ditemukan!")
		return
	}

	fmt.Print("Masukkan Nama Baru     : ")
	fmt.Scan(&dataWarga[index].Nama)

	fmt.Print("Masukkan Alamat Baru   : ")
	fmt.Scan(&dataWarga[index].Alamat)

	fmt.Println("✅ Data warga berhasil diubah!")
}

func hapusWarga() {

	var id int

	fmt.Print("Masukkan ID warga yang ingin dihapus : ")
	fmt.Scan(&id)

	index := sequentialSearchID(id)

	if index == -1 {

		fmt.Println("⚠️ Data warga tidak ditemukan!")
		return
	}
	for i := index; i < jumlahData-1; i++ {

		dataWarga[i] = dataWarga[i+1]
	}

	jumlahData--
	for i := 0; i < jumlahData; i++ {

		dataWarga[i].ID = i + 1
	}

	fmt.Println("✅ Data warga berhasil dihapus!")
	fmt.Println("✅ ID berhasil dirapikan ulang!")
}

func tambahTransaksi() {

	var id int

	fmt.Println()
	fmt.Println("========== TRANSAKSI SAMPAH ==========")

	fmt.Print("Masukkan ID Warga      : ")
	fmt.Scan(&id)

	index := sequentialSearchID(id)

	if index == -1 {

		fmt.Println("⚠️ Warga tidak ditemukan!")
		return
	}

	fmt.Print("Masukkan Jenis Sampah  : ")
	fmt.Scan(&dataWarga[index].JenisSampah)

	if dataWarga[index].JenisSampah == "" {

		fmt.Println("⚠️ Jenis sampah tidak boleh kosong!")
		return
	}

	fmt.Print("Masukkan Berat Sampah  : ")
	fmt.Scan(&dataWarga[index].BeratSampah)

	if dataWarga[index].BeratSampah <= 0 {

		fmt.Println("⚠️ Berat sampah harus lebih dari 0 KG!")
		return
	}

	var tanggal int
	var bulan string
	var tahun int

	fmt.Print("Masukkan Tanggal       : ")
	fmt.Scan(&tanggal)

	if tanggal < 1 || tanggal > 31 {

		fmt.Println("⚠️ Tanggal tidak valid!")
		return
	}

	fmt.Print("Masukkan Bulan         : ")
	fmt.Scan(&bulan)

	if bulan != "Januari" &&
		bulan != "Februari" &&
		bulan != "Maret" &&
		bulan != "April" &&
		bulan != "Mei" &&
		bulan != "Juni" &&
		bulan != "Juli" &&
		bulan != "Agustus" &&
		bulan != "September" &&
		bulan != "Oktober" &&
		bulan != "November" &&
		bulan != "Desember" {

		fmt.Println("⚠️ Nama bulan tidak valid!")
		fmt.Println("Gunakan format contoh: Januari / Mei / Desember")
		return
	}

	fmt.Print("Masukkan Tahun         : ")
	fmt.Scan(&tahun)

	if tahun < 2000 || tahun > 2050 {

		fmt.Println("⚠️ Tahun tidak valid!")
		return
	}

	dataWarga[index].TanggalSetor = fmt.Sprintf("%d %s %d", tanggal, bulan, tahun)

	fmt.Println("✅ Transaksi sampah berhasil ditambahkan!")
}

func selectionSortNama() {

	var i, j, idxMin int
	var temp Warga

	for i = 0; i < jumlahData-1; i++ {

		idxMin = i

		for j = i + 1; j < jumlahData; j++ {

			if dataWarga[j].Nama < dataWarga[idxMin].Nama {

				idxMin = j
			}
		}

		temp = dataWarga[i]
		dataWarga[i] = dataWarga[idxMin]
		dataWarga[idxMin] = temp
	}
}

func insertionSortBerat() {

	var i, j int
	var temp Warga

	for i = 1; i < jumlahData; i++ {

		temp = dataWarga[i]
		j = i - 1

		for j >= 0 && dataWarga[j].BeratSampah < temp.BeratSampah {

			dataWarga[j+1] = dataWarga[j]
			j--
		}

		dataWarga[j+1] = temp
	}
}

func binarySearchNama(nama string) int {

	selectionSortNama()

	var low, high, mid int

	low = 0
	high = jumlahData - 1

	for low <= high {

		mid = (low + high) / 2

		if dataWarga[mid].Nama == nama {

			return mid

		} else if nama < dataWarga[mid].Nama {

			high = mid - 1

		} else {

			low = mid + 1
		}
	}

	return -1
}

func menuSearching() {

	var pilihan int

	fmt.Println()
	fmt.Println("=========== MENU SEARCHING ===========")
	fmt.Println("1. Sequential Search ID")
	fmt.Println("2. Binary Search Nama")
	fmt.Println("======================================")

	fmt.Print("Pilih menu searching : ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {

		var id int

		fmt.Print("Masukkan ID : ")
		fmt.Scan(&id)

		index := sequentialSearchID(id)

		if index == -1 {

			fmt.Println("⚠️ Data tidak ditemukan!")

		} else {

			fmt.Println("✅ Data ditemukan!")
			fmt.Println("Nama :", dataWarga[index].Nama)
		}

	} else if pilihan == 2 {

		var nama string

		fmt.Print("Masukkan Nama : ")
		fmt.Scan(&nama)

		index := binarySearchNama(nama)

		if index == -1 {

			fmt.Println("⚠️ Data tidak ditemukan!")

		} else {

			fmt.Println("✅ Data ditemukan!")
			fmt.Println("Nama :", dataWarga[index].Nama)
		}

	} else {

		fmt.Println("⚠️ Pilihan searching tidak valid!")
	}
}

func menuSorting() {

	var pilihan int

	fmt.Println()
	fmt.Println("============ MENU SORTING ============")
	fmt.Println("1. Selection Sort Nama")
	fmt.Println("2. Insertion Sort Berat Sampah")
	fmt.Println("======================================")

	fmt.Print("Pilih menu sorting : ")
	fmt.Scan(&pilihan)

	switch pilihan {

	case 1:

		selectionSortNama()
		fmt.Println("✅ Data berhasil diurutkan berdasarkan nama!")

	case 2:

		insertionSortBerat()
		fmt.Println("✅ Data berhasil diurutkan berdasarkan berat sampah!")

	default:

		fmt.Println("⚠️ Pilihan sorting tidak valid!")
	}
}

func hitungTotalRecursive(n int) float64 {

	if n == 0 {

		return 0
	}

	return dataWarga[n-1].BeratSampah + hitungTotalRecursive(n-1)
}

func tampilkanStatistik() {

	if jumlahData == 0 {

		fmt.Println("⚠️ Belum ada data!")
		return
	}

	var total float64

	total = hitungTotalRecursive(jumlahData)

	fmt.Println()
	fmt.Println("============ STATISTIK SAMPAH ============")
	fmt.Println("Jumlah Warga            :", jumlahData)
	fmt.Println("Total Sampah Terkumpul  :", total, "KG")

	if jumlahData > 0 {

		fmt.Println("Rata-rata Sampah        :", total/float64(jumlahData), "KG")
	}

	fmt.Println("==========================================")
}

func main() {

	tampilkanHeader()

	var pilihan int

	for pilihan != 9 {

		tampilkanMenu()

		fmt.Print("Masukkan pilihan menu : ")
		fmt.Scan(&pilihan)

		switch pilihan {

		case 1:
			tambahWarga()

		case 2:
			ubahWarga()

		case 3:
			hapusWarga()

		case 4:
			tampilkanWarga()

		case 5:
			tambahTransaksi()

		case 6:
			menuSearching()

		case 7:
			menuSorting()

		case 8:
			tampilkanStatistik()

		case 9:

			fmt.Println()
			fmt.Println("========================================")
			fmt.Println(" Terima kasih telah menggunakan")
			fmt.Println("      WASTE-TRACK SYSTEM")
			fmt.Println("========================================")

		default:

			fmt.Println("⚠️ Menu tidak tersedia!")
		}
	}
}