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


func displayAllData(conn *pgx.Conn){
	lists, err := models.GetAllData(conn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// fmt.Println(lists)

	for i, list := range lists {
		fmt.Printf("No: %d\n", i+1)
		fmt.Printf("Name: %s\n", list.Fullname)
		fmt.Printf("Phone: 0%d\n", list.No_Hp)
		fmt.Printf("Email: %s\n", list.Email)
		fmt.Printf("Created At: %s\n", list.Created_At.Format("02/01/2006 15:04:05"))
	}
}

func addData(conn *pgx.Conn){
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
			}else {
				fmt.Println("data yang diinput bukan number")
			}
	
		}
		
	
		list, err := models.AddDataList(models.ListContact{
			Email: email,
			Fullname: name,
			No_Hp: noHp,
		}, conn)
	
		if err != nil {
			fmt.Println("data Gagal Dinput, Coba lagi")
			continue
		}

		fmt.Printf("Data %s sukses dibuat\n",list.Email)
		break
	}
}



func choiceList(choice *string){
	conn := lib.Conn()
	defer conn.Close(context.Background())

	switch *choice {
	case "1" :
		addData(conn)
		lib.PressEnter("Tekan Enter untuk Kembali")		

	case "2" :
	
	case "3" :
		
	case "4" :
		fmt.Println("List Semua Contact List: ")
		fmt.Println("------------------------------------")
		displayAllData(conn)
		fmt.Println("------------------------------------")
		lib.PressEnter("Tekan Enter untuk Kembali")
	case "5" :
		fmt.Println("Thank You")
		os.Exit(1)
	default :
		fmt.Println("Pilih hanya 1 - 4")
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
