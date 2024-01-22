package main

import (
	"fmt"
)


		/*
		*		ВСПОМОГАТЕЛЬНЫЕ ФУНКЦИИ
		*/

func clearScreen() {
	for i := 0; i < 100; i++ {
		fmt.Println("")
	}
}

func drawLine() {
	fmt.Printf("+")
	for i := 0; i < 49; i++ {
		fmt.Printf("-")
	}
	fmt.Println("+")
}

func updateRecordField(record *Record) *Record {

	var fieldSelect uint8 = 0
	var newValue string = "NONE"

	drawLine()
	fmt.Println(" 1) Username\n 2) Password")
	for {
		fmt.Printf(" Select the record field to edit (1 or 2): "); fmt.Scan(&fieldSelect)

		if fieldSelect == 1 {
			drawLine()
			fmt.Printf(" Enter new username: "); fmt.Scan(&newValue)
			record.username = newValue
			break
		} else if fieldSelect == 2 {
			drawLine()
			fmt.Printf(" Enter new password: "); fmt.Scan(&newValue)
			record.password = newValue
			break
		} else {
			drawLine()
			fmt.Println(" ~~ Incorrect input ~~")
			drawLine()
		}
	}

	return &Record{record.id, record.username, record.password}
}

func checkEmptySlice(RecordsSlice []*Record) bool {
	if RecordsSlice == nil {
		return true
	} else {
		return false
	}
}


		/*
		*		ФУНКЦИИ СТИЛИЗАЦИИ
		*/

func drawMainMenu() {
	clearScreen()

	drawLine()
	fmt.Printf("|%35s%14s|\n", "FoxyPasswordManager", "")
	drawLine()
	fmt.Printf("|%24s%25s|\n", "1 - Create a new record", "")
	fmt.Printf("|%21s%28s|\n", "2 - Read all records", "")
	fmt.Printf("|%20s%29s|\n", "3 - Update a record", "")
	fmt.Printf("|%20s%29s|\n", "4 - Delete a record", "")
	fmt.Printf("|%50s", "|")
	fmt.Printf("\n|%9s%40s|\n", "0 - Exit", "")
	drawLine()
}

func drawCreateHeader() {
	clearScreen()

	drawLine()
	fmt.Printf("|%35s%14s|\n", "Creating a new record", "")
	drawLine()
}

func drawReadHeader() {
	clearScreen()

	drawLine()
	fmt.Printf("|%34s%15s|\n", "Reading all records", "")
	drawLine()
}

func drawUpdateHeader() {
	clearScreen()

	drawLine()
	fmt.Printf("|%33s%16s|\n", "Updating a record", "")
	drawLine()
}

func drawDeleteHeader() {
	clearScreen()

	drawLine()
	fmt.Printf("|%33s%16s|\n", "Deleting a record", "")
	drawLine()
}


		/*
		*		ФУНКЦИИ УПРАВЛЕНИЯ ЗАПИСЯМИ			
		*/

func createRecord(id uint64, username string, password string) *Record {

	fmt.Printf(" Username: "); fmt.Scan(&username)
	fmt.Printf(" Password: "); fmt.Scan(&password)

	clearScreen()
	fmt.Println(" The record has been successfully created.")

	return &Record{id, username, password}
}

func readRecords(RecordsSlice []*Record) {
	for _, record := range RecordsSlice {
		fmt.Printf("%11s%d\n", "ID: ", record.id)
		fmt.Printf("%11s%s\n", "Username: ", record.username)
		fmt.Printf("%11s%s\n", "Password: ", record.password)
		drawLine()
	}
}

func updateRecord(searchingId uint64, RecordsSlice []*Record) []*Record {

	var newRecordsSlice []*Record = nil
	var flag uint8 = 0

	for _, record := range RecordsSlice {

		if record.id != searchingId {
			newRecordsSlice = append(newRecordsSlice, record)
		} else {
			newRecordsSlice = append(newRecordsSlice, updateRecordField(record))
			flag = 1
		}

	}

	clearScreen()

	if flag == 0 {
		fmt.Println(" ~~ There is no record with this ID ~~")
	} else {
		fmt.Println(" The record has been successfully updated. ")
	}

	return newRecordsSlice
}

func deleteRecord(searchingId uint64, RecordsSlice []*Record) []*Record {

	var newRecordsSlice []*Record = nil
	var flag uint8 = 0

	for _, record := range RecordsSlice {

		if record.id != searchingId {
			newRecordsSlice = append(newRecordsSlice, record)
		} else {
			flag = 1
		}

	}

	clearScreen()

	if flag == 0 {
		fmt.Println(" ~~ There is no record with this ID ~~")
	} else {
		fmt.Println(" The record has been successfully deleted.  ")
	}

	return newRecordsSlice
}


type Record struct {		// Шаблон типа структуры для записи
	id uint64
	username string
	password string
}


func main() {

	var (
		Select uint8 = 5				// Выбор пользователя
		RecordsSlice []*Record = nil	// Слайс указателей на начало структуры (записи)
		Counter uint64 = 0				// Для хранения последнего id пользователя
		Exit string = "NONE"			// Для выхода из меню
	)

	for Select != 0 {

		var tempRecord Record = Record{0, "NONE", "NONE"}		// Временная запись для выполнения промежуточных операций (обнуляем в конце каждой итерации)

		drawMainMenu()
		fmt.Printf(" Foxy -> ")
		fmt.Scan(&Select)

		switch Select {

		case 1:		/*    МЕНЮ СОЗДАНИЯ ЗАПИСИ    */

			drawCreateHeader()

			RecordsSlice = append(RecordsSlice, createRecord(Counter, tempRecord.username, tempRecord.password))
			Counter += 1

			fmt.Printf("\n Type 'x' to return to main menu: "); fmt.Scan(&Exit)

		case 2:		/*    МЕНЮ ЧТЕНИЯ ВСЕХ ЗАПИСЕЙ    */

			drawReadHeader()

			if checkEmptySlice(RecordsSlice) == false {

				readRecords(RecordsSlice)

			} else {
				fmt.Println("\n There are no records.")
			}

			fmt.Printf("\n Type 'x' to return to main menu: "); fmt.Scan(&Exit)

		case 3:		/*    МЕНЮ ИЗМЕНЕНИЯ ЗАПИСИ    */

			drawUpdateHeader()

			if checkEmptySlice(RecordsSlice) == false {

				fmt.Printf(" Enter the ID of the record you want to delete: "); fmt.Scan(&tempRecord.id)
				RecordsSlice = updateRecord(tempRecord.id, RecordsSlice)

			} else {
				fmt.Println("\n There are no records.")
			}

			fmt.Printf("\n Type 'x' to return to main menu: "); fmt.Scan(&Exit)

		case 4:		/*    МЕНЮ УДАЛЕНИЯ ЗАПИСИ    */

			drawDeleteHeader()

			if checkEmptySlice(RecordsSlice) == false {

				fmt.Printf(" Enter the ID of the entry you want to delete: "); fmt.Scan(&tempRecord.id)
				RecordsSlice = deleteRecord(tempRecord.id, RecordsSlice)

			} else {
				fmt.Println("\n There are no records.")
			}

			fmt.Printf("\n Type 'x' to return to main menu: "); fmt.Scan(&Exit)

		case 0:     /*    ВЫХОД ИЗ ПРОГРАММЫ    */

			break

		default:

			clearScreen()
			fmt.Println(" ~~ Incorrect selection in the main menu ~~")

			fmt.Printf("\n Type 'x' to return to main menu: "); fmt.Scan(&Exit)

		}
	}
}