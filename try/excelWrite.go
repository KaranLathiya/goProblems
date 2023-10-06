package main

import (
	"log"

	"github.com/xuri/excelize/v2"
)

func main() {
	file := excelize.NewFile() //creates new excel file with one sheet
	file.SetCellValue("Sheet1", "A1", "Name")
	file.SetCellValue("Sheet1", "A2", "Techsolvo")
	file.SetCellValue("Sheet1", "A3", "codersdaily")

	if err := file.SaveAs("companies.xlsx"); err != nil { //checking for errors in creating the file
		log.Fatal(err)
	}
}
