package main

import (
	"conectDB/lib"
	"conectDB/models"
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5"
)

func displayAllData(conn *pgx.Conn) {
	lists, err := models.GetAllData(conn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// fmt.Println(lists)
	fmt.Println("List Semua Contact List: ")
	fmt.Println("------------------------------------")
	for i, list := range lists {
		fmt.Printf("No: %d\n", i+1)
		fmt.Printf("Name: %s\n", list.Fullname)
		fmt.Printf("Phone: 0%d\n", list.No_Hp)
		fmt.Printf("Email: %s\n", list.Email)
		fmt.Printf("Created At: %s\n", list.Created_At.Format("02/01/2006 15:04:05"))
		fmt.Printf("Last Update: %s\n", list.Updated_At.Format("02/01/2006 15:04:05"))
		fmt.Println("------------------------------------")
	}
}

func addData(conn *pgx.Conn) {
	for {
		fmt.Println("Input Contact List")
		fmt.Print("Nama Lengkap: ")
		name := lib.Input()
		fmt.Print("Email: ")
		email := lib.Input()
		noHp := 0
		for {
			fmt.Print("No Hp: ")
			input, err := strconv.Atoi(lib.Input())
			if err == nil {
				noHp = input
				break
			} else {
				fmt.Println("data yang diinput bukan number")
			}

		}

		list, err := models.AddDataList(models.ListContact{
			Email:    email,
			Fullname: name,
			No_Hp:    noHp,
		}, conn)

		if err != nil {
			fmt.Println("data Gagal Dinput, Coba lagi")
			continue
		}

		fmt.Printf("Data %s sukses dibuat\n", list.Email)
		break
	}
}

func editData(conn *pgx.Conn) {
	for {
		displayAllData(conn)
		fmt.Print("Input Email yang ingin update data: ")
		email := lib.Input()

		list, err := models.GetDataByEmail(email, conn)

		if err != nil {
			fmt.Println(err)
			continue
		}
		lib.CallClear()
		fmt.Printf("Email %s Ditemukan\n", list.Email)
		fmt.Println("------------------------------")
		fmt.Println("Update Data")
		for {
			fmt.Printf("Email lama: %s\n", list.Email)
			fmt.Print("Email Baru: ")
			email := lib.Input()
			fmt.Println("---------------------------")
			fmt.Printf("Nama lama: %s\n", list.Fullname)
			fmt.Print("Nama Baru: ")
			name := lib.Input()
			fmt.Println("---------------------------")
			fmt.Printf("No Hp lama: 0%d\n", list.No_Hp)
			noHp := 0
			for {
				fmt.Print("No Hp Baru: ")
				input, err := strconv.Atoi(lib.Input())
				if err == nil {
					noHp = input
					break
				} else {
					fmt.Println("data yang diinput bukan number, Coba lagi")
				}
			}
			update, err := models.UpdateDataList(list.Id, models.ListContact{
				Fullname: name,
				Email:    email,
				No_Hp:    noHp,
			}, conn)

			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("Berhasil Update data", update.Email)
			break
		}
		break
	}
}

func deletData(conn *pgx.Conn) {
	for {
		displayAllData(conn)
		fmt.Print("Input Email yang ingin Dihapus: ")
		email := lib.Input()

		list, err := models.GetDataByEmail(email, conn)

		if err != nil {
			fmt.Println(err)
			continue
		}
		lib.CallClear()
		fmt.Printf("Email %s Ditemukan\n", list.Email)
		fmt.Print("Konfirmasi Hapus Y / N: ")
		choice := lib.Input()
		if choice == "N" || choice == "n" {
			fmt.Println("Delete data dibatalkan")
			break
		} else if choice == "Y" || choice == "y" {
			err := models.DeleteDataList(list.Id, conn)

			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Berhasil Hapus Data")
			break
		}

		lib.PressEnter("Pilihan Salah Enter untuk Ulangi Input email")
		lib.CallClear()

	}

}

func choiceList(choice *string) {
	conn := lib.Conn()
	defer conn.Close(context.Background())

	switch *choice {
	case "1":
		lib.CallClear()
		addData(conn)
		lib.PressEnter("Tekan Enter untuk Kembali")
	case "2":
		lib.CallClear()
		editData(conn)
		lib.PressEnter("Tekan Enter untuk Kembali")
	case "3":
		lib.CallClear()
		deletData(conn)
		lib.PressEnter("Tekan Enter untuk Kembali")
	case "4":
		lib.CallClear()
		displayAllData(conn)
		lib.PressEnter("Tekan Enter untuk Kembali")
	case "5":
		fmt.Println("Thank You")
		os.Exit(1)
	default:
		fmt.Println("Pilih hanya 1 - 4")
		lib.PressEnter("Tekan Enter untuk Kembali")
	}

}

func main() {
	// fmt.Print("Masukkan Email: ")
	// email := lib.Input()

	for {
		lib.CallClear()
		fmt.Println(`------------List Contact--------------------
		1. Add List
		2. Edit Data
		3. Delete Data
		4. Display All Data
		5. Exit`)
		fmt.Println("---------------------------------------")
		fmt.Print("Input Pilihan: ")
		choice := lib.Input()
		choiceList(&choice)
	}

}
