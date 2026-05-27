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
	Minggu       int
}

var dataWarga [NMAX]Warga
var jumlahData int

//Tampilan Awal
func tampilkanHeader() {

	fmt.Println("===================================================")
	fmt.Println("         WASTE-TRACK MANAGEMENT SYSTEM")
	fmt.Println("===================================================")
	fmt.Println("Kelompok 10:")
	fmt.Println("1. Rakhmat Pratama - 109082530037")
	fmt.Println("2. Ahmad Luthfi Habibie - 109082500190")
	fmt.Println("3. Muhammad Haidar Az Zacky - 109082530035")
	fmt.Println("===================================================")
	fmt.Println()
}

//Tampilan Menu Sesuai Spesifikasi Program
func tampilkanMenu() {

	fmt.Println()
	fmt.Println("============== MENU UTAMA ==============")
	fmt.Println("1. Tambah Data Warga")
	fmt.Println("2. Ubah Data Warga")
	fmt.Println("3. Hapus Data Warga")
	fmt.Println("4. Tampilkan Data Warga")
	fmt.Println("5. Tambah Transaksi Setoran Sampah")
	fmt.Println("6. Searching Data Warga")
	fmt.Println("7. Sorting Data Warga")
	fmt.Println("8. Statistik Sampah Mingguan")
	fmt.Println("9. Selesai")
	fmt.Println("========================================")
}

//Untuk Menambahkan Data Warga
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

//Untuk Menampilkan Data Warga
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
			fmt.Println("Minggu Ke-      :", dataWarga[i].Minggu)
		}
	}

	fmt.Println("------------------------------------------")
}

//Untuk Mengubah Data Warga yang sudah diisi
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

//Untuk Menghapus Data Warga
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

//Untuk Menambahkan Transaksi Setoran sampah oleh warga yang sudah ada di data
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

	fmt.Print("Masukkan Minggu Ke-    : ")
	fmt.Scan(&dataWarga[index].Minggu)

	if dataWarga[index].Minggu < 1 || dataWarga[index].Minggu > 4 {

	fmt.Println("⚠️ Minggu hanya 1 - 4!")
	return
}

	fmt.Println("✅ Transaksi sampah berhasil ditambahkan!")
}

//Untuk Mencari Data Warga menggunakan ID dengan Sequential Search
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

//Untuk Mencari Data Warga menggunakan Nama dengan Sequential Search
func sequentialSearchNama(nama string) int {

	var i int
	var found bool

	i = 0
	found = false

	for i < jumlahData && !found {

		if dataWarga[i].Nama == nama {

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

//Untuk Mencari Data Warga menggunakan ID dengan Binary Search
func binarySearchID(id int) int {

	selectionSortID()

	var low, high, mid int

	low = 0
	high = jumlahData - 1

	for low <= high {

		mid = (low + high) / 2

		if dataWarga[mid].ID == id {

			return mid

		} else if id < dataWarga[mid].ID {

			high = mid - 1

		} else {

			low = mid + 1
		}
	}

	return -1
}

//Nyambung keatas yang BinarySearch Pakai ID
func selectionSortID() {

	var i, j, idxMin int
	var temp Warga

	for i = 0; i < jumlahData-1; i++ {

		idxMin = i

		for j = i + 1; j < jumlahData; j++ {

			if dataWarga[j].ID < dataWarga[idxMin].ID {

				idxMin = j
			}
		}

		temp = dataWarga[i]
		dataWarga[i] = dataWarga[idxMin]
		dataWarga[idxMin] = temp
	}
}

//Untuk Mencari Data Warga menggunakan Nama dengan Binary Search
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

//Nyambung keatas yang BinarySearch Pakai Nama
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

//Untuk Mengurutkan Data Warga menggunakan Berat dengan Selection Sort
func selectionSortBerat() {

	var i, j, idxMax int
	var temp Warga

	for i = 0; i < jumlahData-1; i++ {

		idxMax = i

		for j = i + 1; j < jumlahData; j++ {

			if dataWarga[j].BeratSampah > dataWarga[idxMax].BeratSampah {

				idxMax = j
			}
		}

		temp = dataWarga[i]
		dataWarga[i] = dataWarga[idxMax]
		dataWarga[idxMax] = temp
	}
}

//Untuk Mengurutkan Data Warga menggunakan Berat dengan Insertion Sort
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

//Menu Utama Saat kita akan mencari Data Warga Menggunakan Sequential Search dan Binary Search
func menuSearching() {

	var pilihan int

	fmt.Println()
	fmt.Println("=========== MENU SEARCHING ===========")
	fmt.Println("1. Sequential Search ID")
	fmt.Println("2. Sequential Search Nama")
	fmt.Println("3. Binary Search ID")
	fmt.Println("4. Binary Search Nama")
	fmt.Println("======================================")

	fmt.Print("Pilih menu searching : ")
	fmt.Scan(&pilihan)

	switch pilihan {

	case 1:

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

	case 2:

		var nama string

		fmt.Print("Masukkan Nama : ")
		fmt.Scan(&nama)

		index := sequentialSearchNama(nama)

		if index == -1 {

			fmt.Println("⚠️ Data tidak ditemukan!")

		} else {

			fmt.Println("✅ Data ditemukan!")
			fmt.Println("Nama :", dataWarga[index].Nama)
		}

	case 3:

		var id int

		fmt.Print("Masukkan ID : ")
		fmt.Scan(&id)

		index := binarySearchID(id)

		if index == -1 {

			fmt.Println("⚠️ Data tidak ditemukan!")

		} else {

			fmt.Println("✅ Data ditemukan!")
			fmt.Println("Nama :", dataWarga[index].Nama)
		}

	case 4:

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

	default:

		fmt.Println("⚠️ Pilihan searching tidak valid!")
	}
}

//Menu Utama Saat kita akan mengurutkan Data Warga Menggunakan Selection Sort dan Insertion Sort dengan Jumlah Sampah Terbanyak
func menuSorting() {

	var pilihan int

	fmt.Println()
	fmt.Println("============ MENU SORTING ============")
	fmt.Println("1. Selection Sort Berat Sampah")
	fmt.Println("2. Insertion Sort Berat Sampah")
	fmt.Println("======================================")

	fmt.Print("Pilih menu sorting : ")
	fmt.Scan(&pilihan)

	switch pilihan {

	case 1:

		selectionSortBerat()

		fmt.Println("✅ Data berhasil diurutkan menggunakan Selection Sort!")

	case 2:

		insertionSortBerat()

		fmt.Println("✅ Data berhasil diurutkan menggunakan Insertion Sort!")

	default:

		fmt.Println("⚠️ Pilihan sorting tidak valid!")
	}
}

//Menu Utama Saat kita akan menampilkan Data Statistik Total Akumulasi Sampah dalam satu Minggu
func tampilkanStatistik() {

	if jumlahData == 0 {

		fmt.Println("⚠️ Belum ada data!")
		return
	}

	var minggu int
	var total float64
	var jumlah int

	fmt.Println()
	fmt.Println("========== STATISTIK SAMPAH ==========")

	fmt.Print("Masukkan Minggu Ke- (1-4) : ")
	fmt.Scan(&minggu)

	if minggu < 1 || minggu > 4 {

		fmt.Println("⚠️ Minggu hanya 1 - 4!")
		return
	}

	for i := 0; i < jumlahData; i++ {

		if dataWarga[i].Minggu == minggu {

			total += dataWarga[i].BeratSampah
			jumlah++
		}
	}

	if jumlah == 0 {

		fmt.Println("⚠️ Tidak ada transaksi pada minggu tersebut!")
		return
	}

	fmt.Println()
	fmt.Println("=========== STATISTIK TOTAL SAMPAH MINGGUAN ===========")
	fmt.Println("Minggu Ke-             :", minggu)
	fmt.Println("Jumlah Transaksi       :", jumlah)
	fmt.Println("Total Sampah           :", total, "KG")
	fmt.Println("Rata-rata Sampah       :", total/float64(jumlah), "KG")
	fmt.Println("=======================================")
}

//Menu Utama untuk keseluruhan Fitur
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
			fmt.Println("   Terima kasih telah menggunakan  ")
			fmt.Println(" WASTE-TRACK SYSTEM BY KELOMPOK 10 ")
			fmt.Println("========================================")

		default:

			fmt.Println("⚠️ Menu tidak tersedia!")
		}
	}
}
