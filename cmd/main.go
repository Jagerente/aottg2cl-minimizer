package main

import (
	"aottg2cl-minimizer/internal/minimizer"
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Drag and drop script file to this executable.")
		fmt.Println("Or pass the file path as an argument.")
		time.Sleep(3 * time.Second)
		return
	}

	filePath := os.Args[1]

	m := minimizer.New()

	if err := m.MinimizeFile(filePath); err != nil {
		fmt.Printf("Failed to minimize file: %v\n", err)
		time.Sleep(5 * time.Second)
	}
}
