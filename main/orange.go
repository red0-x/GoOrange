package main

import (
	"fmt"
	listutils "listutils/m/v2"
	"os"
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
[1] Customize Questions [WIP Not Functional]
[2] Generate PDF [WIP Not Functional]
	  			`)
	color.Red("\nChoose an option.")
	var option int
	fmt.Scan(&option)
	switch option {
	case 1:
		listgen()
	case 2:
		pdfgen()
	default:
		color.Red("[!] Choose an Option from 1-2!")
		time.Sleep(2 * time.Second)
		menu()
	}
}

func wordlistSelect(listdir string) string {
	dirs, err := os.ReadDir(listdir)
	if err != nil {
		fmt.Println("[!] Err reading wordlist /lists folder, any contents?:", err)
		return "lists/Easy.txt"
	}
	for i, dir := range dirs {
		fmt.Printf("[%d] %s\n", i+1, dir.Name())
	}
	var listChoice int
	fmt.Scan(&listChoice)
	var selectedList string
	selectedList = dirs[listChoice-1].Name()
	wlPath := "./lists/" + selectedList
	return wlPath
}

func listgen() {
	color.Yellow("\n[?] Choose a wordlist to customize:")

	entries, err := listutils.ReadList(wordlistSelect("./lists/"))
	fmt.Print("\n")
	if err != nil {
		fmt.Println("\n[!] Error reading list:", err)
		return
	}
	color.Green("\n[i] Parsed:")
	for i, e := range entries {
		fmt.Printf("[%d] ", i+1)
		for k, v := range e {
			fmt.Printf("%s=%s ", k, v)
		}
		fmt.Println()
	}
	color.Green("\n[!] Feature WIP.")
}

func pdfgen() {
	color.Yellow("\n[?] Choose a WordList:\n")
	entries, err := listutils.ReadList(wordlistSelect("./lists/"))
	if err != nil {
		fmt.Println("\n[!] Error reading list:", err)
		return
	}
	color.Yellow("[i] Preview of entries:")
	for i, e := range entries {
		if i >= 5 {
			break
		}
		fmt.Printf("[%d] ", i+1)
		for k, v := range e {
			fmt.Printf("%s=%s ", k, v)
		}
		fmt.Println()
	}
	_ = entries

	var fname string = "orange.pdf"
	var contents string = "Go Orange!"
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, contents)

	if err := pdf.OutputFileAndClose(fname); err != nil {
		fmt.Println("[!] Error Creating PDF")
	}
	color.Green("\n[!] Feature WIP, so far this doesn't create a full PDF.")
}
