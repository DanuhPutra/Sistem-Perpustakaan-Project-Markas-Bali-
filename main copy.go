package main
import (
	"fmt"
	// "os"
	// "strings"
	// "bufio"
	// tm "github.com/buger/goterm"
	"github.com/MasterDimmy/go-cls"
)

type DataBukuPerpustakaan struct {
	KodeBuku string 
	JudulBuku string
	Pengarang string
	Penerbit string
	JumlahHalaman int
	TahunTerbit int
}

var ListBukuPerpustakaan []DataBukuPerpustakaan

func TambahBukuBaru(){
	cls.CLS()
	// deklarasi variabel
	// KodeBukuBaru := bufio.NewReader(os.Stdin)
	// JudulBukuBaru := bufio.NewReader(os.Stdin)
	// PengarangBukuBaru := bufio.NewReader(os.Stdin)
	// PenerbitBukuBaru := bufio.NewReader(os.Stdin)
	KodeBukuBaru := ""
	JudulBukuBaru := ""
	PengarangBukuBaru := ""
	PenerbitBukuBaru := ""
	JumlahHalamanBukuBaru := 0
	TahunTerbitBukuBaru := 0 

	fmt.Println("==============================")
	fmt.Println("Menambahkan Buku Baru")
	fmt.Println("==============================\n")

	// kode buku
	// fmt.Print("Kode Buku Baru : ")
	// KodeBukuTambah, err := KodeBukuBaru.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("Ups, Terjadi error pada Kode Buku!", err)
	// 	return
	// }
	// KodeBukuTambah = strings.Replace(
	// 	KodeBukuTambah,
	// 	"\n",
	// 	"",
	// 	1,
	// )

	// judul buku
	// fmt.Print("Judul Buku Baru : ")
	// JudulBukuTambah, err := JudulBukuBaru.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("Ups, Terjadi error pada Judul Buku!", err)
	// 	return
	// }
	// JudulBukuTambah = strings.Replace(
	// 	JudulBukuTambah,
	// 	"\n",
	// 	"",
	// 	1,
	// )

	// pengarang buku
	// fmt.Print("Pengarang Buku Baru : ")
	// PengarangBukuTambah, err := PengarangBukuBaru.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("Ups, Terjadi error pada Judul Buku!", err)
	// 	return
	// }
	// PengarangBukuTambah = strings.Replace(
	// 	PengarangBukuTambah,
	// 	"\n",
	// 	"",
	// 	1,
	// )

	// penerbit buku
	// fmt.Print("Penerbit Buku Baru : ")
	// PenerbitBukuTambah, err := PenerbitBukuBaru.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("Ups, Terjadi error pada Judul Buku!",err)
	// 	return
	// }
	// PenerbitBukuTambah = strings.Replace(
	// 	PenerbitBukuTambah,
	// 	"\n",
	// 	"",
	// 	1,
	// )



	fmt.Print("kode buku :")
	_, err := fmt.Scanln(&KodeBukuBaru)
	if err != nil {
		fmt.Println("Ups, Terjadi error pada Jumlah Halaman Buku! :", err)
		return
	}
	fmt.Print("judul buku :")
	_, err = fmt.Scanln(&JudulBukuBaru)
	if err != nil {
		fmt.Println("Ups, Terjadi error pada Jumlah Halaman Buku! :", err)
		return
	}
	fmt.Print("pengarang :")
	_, err = fmt.Scanln(&PengarangBukuBaru)
	if err != nil {
		fmt.Println("Ups, Terjadi error pada Jumlah Halaman Buku! :", err)
		return
	}
	fmt.Print("penerbit :")
	_, err = fmt.Scanln(&PenerbitBukuBaru)
	if err != nil {
		fmt.Println("Ups, Terjadi error pada Jumlah Halaman Buku! :", err)
		return
	}

	// jumlah halaman buku
	fmt.Print("silahkan masukan Jumlah Halaman pada Buku Baru :")
	_, err = fmt.Scanln(&JumlahHalamanBukuBaru)
	if err != nil {
		fmt.Println("Ups, Terjadi error pada Jumlah Halaman Buku! :", err)
		return
	}

	// tahun terbit buku
	fmt.Print("silahkan masukan Tahun Terbit Buku :")
	_, err = fmt.Scanln(&TahunTerbitBukuBaru)
	if err != nil {
		fmt.Println("Ups, Terjadi error pada Tahun Terbit Buku! :", err)
		return
	}

	ListBukuPerpustakaan = append(ListBukuPerpustakaan, DataBukuPerpustakaan {
		KodeBuku : KodeBukuBaru,
		JudulBuku : JudulBukuBaru,
		Pengarang : PengarangBukuBaru,
		Penerbit : PenerbitBukuBaru,
		JumlahHalaman : JumlahHalamanBukuBaru,
		TahunTerbit : TahunTerbitBukuBaru,
	}) 
	fmt.Println("berhasil menambahkan buku baru kedalam perpustakaan!")


}

func TampilkanListBuku(){
	fmt.Println("======================================")
	fmt.Println("List Buku yang ada di Perpustakaan ini")
	fmt.Println("======================================")

	for _, DataBukuPerpustakaan := range ListBukuPerpustakaan{
		fmt.Printf("%s,%s,%s,%s,%d,%d",
			DataBukuPerpustakaan.KodeBuku,
			DataBukuPerpustakaan.JudulBuku,
			DataBukuPerpustakaan.Pengarang,
			DataBukuPerpustakaan.Penerbit,
			DataBukuPerpustakaan.JumlahHalaman,
			DataBukuPerpustakaan.TahunTerbit,
		)

		fmt.Println("\n======================================")
	}


}

func main(){
	var PilihanAksi int

	fmt.Println("===========================================")
	fmt.Println("Aplikasi Manajemen Daftar Buku Perpustakaan")
	fmt.Println("===========================================")
	fmt.Println("silahkan pilih menu : ")
	fmt.Println("1. Menambahkan Buku Baru")
	fmt.Println("2. Menampilkan List Buku")
	fmt.Println("3. Hapus Buku")
	fmt.Println("4. Edit Buku")
	fmt.Println("===========================================")
	fmt.Print("masukan pilihan : ")
	_, err := fmt.Scanln(&PilihanAksi)
	if err != nil {
		fmt.Println("Ups, Terjadi error pada aksi yang kamu pilih!", err)
	}

	switch PilihanAksi{
		case 1 :
			TambahBukuBaru()
		case 2 :
			TampilkanListBuku()
	}
	main()

}

func HapusDataBukuPerpustakaan(){

}