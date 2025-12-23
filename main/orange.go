package main

import (
	"fmt"
	"time"

	"codeberg.org/go-pdf/fpdf"
	"github.com/fatih/color"
)

func main() {
	menu()
}

func menu() {
	color.Yellow(`
    ⣀⣀⣀⣀⣀⠀⠀⠀⢀⣠⡴⢶⡾⣆⠀⠀⠀⠀⠀⠀
⠀⠐⣾⠛⠋⠉⠛⠋⠙⢻⣦⣠⡟⠁⠀⢸⢠⡿⠀⠀⠀⠀⠀⠀
⠀⠀⠹⣦⡀⠀⠀⠀⠀⠀⢹⣿⣓⣠⣴⣵⠟⠁⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠘⠻⣦⣤⣤⣤⣤⣼⠾⠻⠿⠿⣷⣤⡀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⣠⡾⠛⠁⠀⠀⠀⠀⠀⠀⠀⠀⠉⠪⣽⢷⣄⠀⠀⠀⠀
⠀⢠⡾⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⢣⡙⢷⡄                             
⢠⡟⠁⢰⣰⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢱⡈⢿⡄   
⣾⠃⠀⠀⠀⣀⡢⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠃⠘⣿⠀ 
⣿⠀⠸⠜⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⠀⢿ 
⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⠀⣿
⢿⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡄⢠⡿  
⠈⢿⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡜⢁⣾⠃⠀
⠀⠈⢻⣆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡀⣰⡟⠁⠀⠀
⠀⠀⠀⠙⠷⣤⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣐⣥⠟⠋⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠈⠉⠛⠶⠶⠦⠤⠶⣶⠾⠛⠋⠁`)
	color.Yellow(`
[1] Generate List
[2] Select Question Types 
[3] Generate PDF
	  			`)
	color.Red("\nChoose an option.")
	var option int
	fmt.Scan(&option)
	switch option {
	case 1:
		listgen()
	case 2:
		questionselection()
	case 3:
		pdfgen()
	default:
		color.Red("Choose an Option from 1-3!")
		time.Sleep(2)
		menu()
	}
}

func listgen() {
	fmt.Println("Work in progress!")
}

func questionselection() {
	fmt.Println("Work in progress!")
}

func pdfgen() {
	var fname string = "orange.pdf"
	var contents string = "Go Orange!"
	fmt.Println("Work in progress!")
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, contents)

	err := pdf.OutputFileAndClose(fname)
	if err != nil {
		fmt.Println("Error Creating PDF")
	}
}
