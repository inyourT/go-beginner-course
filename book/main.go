package main

import "fmt"

type Book struct {
	Judul   string
	Author  string
	Halaman int
}

var books []Book

func addNewBook() {
	var judul, author string
	var halaman int
	fmt.Println("Masukan Judul :")
	fmt.Scan(&judul)
	fmt.Println("Masukan Author :")
	fmt.Scan(&author)
	fmt.Println("Masukan Jumlah Halaman :")
	fmt.Scan(&halaman)

	b := Book{Judul: judul, Author: author, Halaman: halaman}
	books = append(books, b)
	fmt.Println("Buku berhasil ditambah")
}


func getBook() {
	// Cek apakah buku ada
	if len(books) == 0 {
		fmt.Println("Belum ada buku.")
		return
	}
	fmt.Print("\n Daftar buku :")
	for i, b := range books{
		fmt.Printf("%d. %s - %s - %d halaman\n", i+1, b.Judul, b.Author, b.Halaman)
	}
}

func deleteBoo() {
	// cek apakah data buku ada
	var index int
	fmt.Print("Masukan nomer buku yang ingin di hapus :")
	fmt.Scan(&index)

	if index < 1 || index > len(books){
		fmt.Println("Nomer buku tidak ditemukan")
		return
	}

	// Jika buku ada
	books = append(books[:index-1], books[index:]...)
	fmt.Println("Buku berhasil di hapus")

}

func putBook() {
	// cek apakah data buku ada
	var index int
	fmt.Print("Masukan nomer buku yang ingin diubah :")
	fmt.Scan(&index)

	if index < 1 || index > len(books){
		fmt.Println("Nomer buku tidak ditemukan")
		return
	}

	var newJudul, newAuthor string
	var newHalaman int

	fmt.Print("Masukan judul baru:")
	fmt.Scan(&newJudul)
	fmt.Print("Masukan author baru:")
	fmt.Scan(&newAuthor)
	fmt.Print("Masukan halaman baru:")
	fmt.Scan(&newHalaman)

	books[index-1].Judul = newJudul
	books[index-1].Author = newAuthor
	books[index-1].Halaman = newHalaman

	fmt.Print("Data buku berhasil di ubah")
}

func main() {
	for{
		fmt.Println("Tess App")
		fmt.Println("1. Tambah Mahasiswa")
		fmt.Println("2. Lihat Daftar Mahasiswa")
		fmt.Println("3. Ubah Data Mahasiswa")
		fmt.Println("4. Hapus Data Mahasiswa")
		fmt.Println("5. Keluar")
		fmt.Println("Pilih menu: ")
		
		var pilihan int
		fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		addNewBook()
	case 2:
		getBook()
	case 3:
		putBook()
	case 4:
		deleteBoo()
	case 5:
		fmt.Println("Keluar dari program ...")
			return
	default:
		fmt.Println("Pilihan tidak valid.")
		}
	}	
}