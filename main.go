package main

import (
	"fmt"
	"os"
)

// qr_code_generator - Generate QR codes
func qr_code_generator(path string) {
	fmt.Println("========================================")
	fmt.Println("  QR-Code-Generator")
	fmt.Println("  Generate QR codes")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	qr_code_generator(path)
}
