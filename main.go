package main
import (
	"fmt"
	"os"
	"strings"
	"bufio"
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
	KodeBukuBaru := bufio.NewReader(os.Stdin)
	JudulBukuBaru := bufio.NewReader(os.Stdin)
	PengarangBukuBaru := bufio.NewReader(os.Stdin)
	PenerbitBukuBaru := bufio.NewReader(os.Stdin)
	JumlahHalamanBukuBaru := 0
	TahunTerbitBukuBaru := 0 

	fmt.Println("==============================")
	fmt.Println("Menambahkan Buku Baru")
	fmt.Println("==============================\n")

	// kode buku
	fmt.Print("Kode Buku Baru : ")
	KodeBukuTambah, err := KodeBukuBaru.ReadString('\n')
	if err != nil {
		fmt.Println("Ups, Terjadi error pada Kode Buku!", err)
		return
	}
	KodeBukuTambah = strings.TrimSpace(KodeBukuTambah)

	// judul buku
	fmt.Print("Judul Buku Baru : ")
	JudulBukuTambah, err := JudulBukuBaru.ReadString('\n')
	if err != nil {
		fmt.Println("Ups, Terjadi error pada Judul Buku!", err)
		return
	}
	JudulBukuTambah = strings.TrimSpace(JudulBukuTambah)

	// pengarang buku
	fmt.Print("Pengarang Buku Baru : ")
	PengarangBukuTambah, err := PengarangBukuBaru.ReadString('\n')
	if err != nil {
		fmt.Println("Ups, Terjadi error pada Judul Buku!", err)
		return
	}
	PengarangBukuTambah = strings.TrimSpace(PengarangBukuTambah)

	// penerbit buku
	fmt.Print("Penerbit Buku Baru : ")
	PenerbitBukuTambah, err := PenerbitBukuBaru.ReadString('\n')
	if err != nil {
		fmt.Println("Ups, Terjadi error pada Judul Buku!",err)
		return
	}
	PenerbitBukuTambah = strings.TrimSpace(PenerbitBukuTambah)

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

	// memasukan data inputan kedalam struct 
	ListBukuPerpustakaan = append(ListBukuPerpustakaan, DataBukuPerpustakaan {
		KodeBuku : KodeBukuTambah,
		JudulBuku : JudulBukuTambah,
		Pengarang : PengarangBukuTambah,
		Penerbit : PenerbitBukuTambah,
		JumlahHalaman : JumlahHalamanBukuBaru,
		TahunTerbit : TahunTerbitBukuBaru,
	}) 
	fmt.Println("berhasil menambahkan buku baru kedalam perpustakaan!")
	fmt.Println("\n======================================")
	fmt.Println("Tekan 'Enter' untuk melanjutkan...")		
	bufio.NewReader(os.Stdin).ReadBytes('\n')

}

func TampilkanListBuku(){
	cls.CLS()
	fmt.Println("======================================")
	fmt.Println("List Buku yang ada di Perpustakaan ini")
	fmt.Println("======================================\n")

	for banyakBuku, DataBukuPerpustakaan := range ListBukuPerpustakaan{
		fmt.Printf(" | %d. | Kode Buku : %s | Judul Buku : %s | Pengarang Buku : %s | Penerbit Buku :%s | Jumlah Halaman Buku : %d | Tahun Terbit Buku : %d |\n",
			banyakBuku + 1,
			DataBukuPerpustakaan.KodeBuku,
			DataBukuPerpustakaan.JudulBuku,
			DataBukuPerpustakaan.Pengarang,
			DataBukuPerpustakaan.Penerbit,
			DataBukuPerpustakaan.JumlahHalaman,
			DataBukuPerpustakaan.TahunTerbit,
		)
	}

	fmt.Println("\n======================================")
	fmt.Println("Tekan 'Enter' untuk melanjutkan...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

}

func main(){
	cls.CLS()
	var PilihanAksi int

	fmt.Println("===========================================")
	fmt.Println("Aplikasi Manajemen Daftar Buku Perpustakaan")
	fmt.Println("===========================================")
	fmt.Println("silahkan pilih menu : ")
	fmt.Println("1. Menambahkan Buku Baru Perpustakaan")
	fmt.Println("2. Menampilkan Buku Perpustakaan")
	fmt.Println("3. Hapus Buku Perpustakaan")
	fmt.Println("4. Edit Buku Perpustakaan")
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
		case 3 :
			HapusDataBukuPerpustakaan()
		case 4 :
			UpdateDataBukuPerpustakaan()
	}
	main()

}

func HapusDataBukuPerpustakaan(){
	cls.CLS()
	var KodeBukuPerpustakaan string
	fmt.Println("======================================")
	fmt.Println("Hapus Buku yang ada diPerpustakaan")
	fmt.Println("======================================\n")
	TampilkanListBuku()
	fmt.Println("======================================\n")

	fmt.Print("Masukan Kode Buku yang ingin dihapus : ")
	fmt.Scanln(&KodeBukuPerpustakaan)

	for i,j := range ListBukuPerpustakaan {
		if j.KodeBuku == KodeBukuPerpustakaan{
			ListBukuPerpustakaan = append(
				ListBukuPerpustakaan[:i],
				ListBukuPerpustakaan[i+1:]...
			)
			fmt.Println("Buku Perpustakaan Berhasil Dihapus!")
		} else {
			fmt.Println("Terjadi error dalam pemilihan Kode Buku atau kode buku tidak ada!")
		}
	}

	fmt.Println("\n======================================")
	fmt.Println("Tekan 'Enter' untuk melanjutkan...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

}

func UpdateDataBukuPerpustakaan() {
	cls.CLS()
	var UpdateDenganKode string

	fmt.Println("======================================")
	fmt.Println("Update Data Buku yang ada diPerpustakaan")
	fmt.Println("======================================\n")
	TampilkanListBuku()
	fmt.Println("======================================\n")

	fmt.Print("Masukan Kode Buku yang ingin diupdate : ")
	fmt.Scanln(&UpdateDenganKode)

	for i, dataBuku := range ListBukuPerpustakaan{
		if dataBuku.KodeBuku == UpdateDenganKode{
			dataBaruUpdate := bufio.NewReader(os.Stdin)

			fmt.Println("=====================================================")
			fmt.Println("Data yang akan diupdate dari kode : ", dataBuku.KodeBuku)
			fmt.Println("=====================================================\n")

			fmt.Printf("Judul Baru : ")
			ListBukuPerpustakaan[i].JudulBuku, _ = dataBaruUpdate.ReadString('\n')
			ListBukuPerpustakaan[i].JudulBuku = strings.TrimSpace(ListBukuPerpustakaan[i].JudulBuku)
			fmt.Printf("Pengarang Baru : ")
			ListBukuPerpustakaan[i].Pengarang, _ = dataBaruUpdate.ReadString('\n')
			ListBukuPerpustakaan[i].Pengarang = strings.TrimSpace(ListBukuPerpustakaan[i].Pengarang)
			fmt.Printf("Penerbit Baru : ")
			ListBukuPerpustakaan[i].Penerbit, _ = dataBaruUpdate.ReadString('\n')
			ListBukuPerpustakaan[i].Penerbit = strings.TrimSpace(ListBukuPerpustakaan[i].Penerbit)
			fmt.Printf("Jumlah Halaman Baru : ")
			fmt.Scanln(&ListBukuPerpustakaan[i].JumlahHalaman)
			fmt.Printf("Tahun Terbit Baru : ")
			fmt.Scanln(&ListBukuPerpustakaan[i].TahunTerbit)

			fmt.Println("Data berhasil diupdate!")
			fmt.Println("\n======================================")
			fmt.Println("Tekan 'Enter' untuk melanjutkan...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			return
		} else {
			fmt.Println("Terjadi error dalam pemilihan Kode Buku atau kode buku tidak ada!")
			fmt.Println("\n======================================")
			fmt.Println("Tekan 'Enter' untuk melanjutkan...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
		}
	}

}