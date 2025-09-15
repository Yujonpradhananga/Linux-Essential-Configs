package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the path to your PDF or EPUB file: ")
	filePath, _ := reader.ReadString('\n')
	filePath = strings.TrimSpace(filePath)

	if filePath == "" {
		fmt.Println("No file path provided. Exiting.")
		return
	}

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Printf("File not found at: %s\n", filePath)
		return
	}

	// Check file extension
	ext := strings.ToLower(filepath.Ext(filePath))
	if ext != ".pdf" && ext != ".epub" {
		fmt.Printf("Unsupported file format: %s\nSupported formats: .pdf, .epub\n", ext)
		return
	}

	// Start the document viewer
	viewer := NewDocumentViewer(filePath)
	if err := viewer.Open(); err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}

	viewer.Run()
}
