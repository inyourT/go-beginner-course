package main

import "fmt"

type Mahasiswa struct {
	Nama string
	NIM  string
}

var dataMahasiswa []Mahasiswa

func tambahMahasiswa() {
	var nama, nim string
	fmt.Print("Masukan nama: ")
	fmt.Scan(&nama)
	fmt.Print("Masukan NIM: ")
	fmt.Scan(&nim)

	m := Mahasiswa{Nama: nama, NIM: nim}
	dataMahasiswa = append(dataMahasiswa, m)
	fmt.Println("Mahasiswa berhasil ditambahkan")
}

func lihatMahasiswa() {
	if len(dataMahasiswa) == 0 {
		fmt.Println("Belum ada data mahasiswa.")
		return
	}
	fmt.Println("\nDaftar mahasiswa:")
	for i, m := range dataMahasiswa {
		fmt.Printf("%d. %s - %s\n", i+1, m.Nama, m.NIM)
	}
}

func hapusMahaiswa() {
	var index int
	fmt.Print("Masukan nomer yang akan dihapus: ")
	fmt.Scan(&index)

	if index < 1 || index > len(dataMahasiswa){
		fmt.Println("Indeks tidak valid !")
		return
	}

	dataMahasiswa = append(dataMahasiswa[:index-1], dataMahasiswa[index:]...)
	fmt.Println("Mahasiswa berhasil dihapus")
}

func ubahMahasiswa() {
	var index int
	lihatMahasiswa()
	fmt.Print("Masukan nomer yang ingin diubah: ")
	fmt.Scan(&index)

	if index < 1 || index > len(dataMahasiswa){
		fmt.Println("Index tidak ditemukan")
		return
	}

	var namaBaru, nimBaru string
	fmt.Print("Nama Baru: ")
	fmt.Scan(&namaBaru)
	fmt.Print("Masukan NIM baru: ")
	fmt.Scan(nimBaru)

	dataMahasiswa[index-1].Nama = namaBaru
	dataMahasiswa[index-1].NIM = nimBaru

	fmt.Println("Data mahasiswa berhasil diubah!")
}

func main() {
	for {
		fmt.Println("\n=== MENU CRUD MAHASISWA ===")
		fmt.Println("1. Tambah Mahasiswa")
		fmt.Println("2. Lihat Daftar Mahasiswa")
		fmt.Println("3. Ubah Data Mahasiswa")
		fmt.Println("4. Hapus Data Mahasiswa")
		fmt.Println("5. Keluar")
		fmt.Println("Pilih menu: ")
		
		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan{
		case 1:
			tambahMahasiswa()
		case 2:
			lihatMahasiswa()
		case 3:
			ubahMahasiswa()
		case 4:
			hapusMahaiswa()
		case 5:
			fmt.Println("Keluar dari program ...")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}