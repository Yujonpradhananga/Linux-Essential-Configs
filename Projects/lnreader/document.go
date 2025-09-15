package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gen2brain/go-fitz"
)

type DocumentViewer struct {
	doc         *fitz.Document
	currentPage int
	textPages   []int
	reader      *bufio.Reader
	path        string
	oldState    *termios
	fileType    string // "pdf" or "epub"
}

type termios struct {
	Iflag  uint32
	Oflag  uint32
	Cflag  uint32
	Lflag  uint32
	Cc     [20]uint8
	Ispeed uint32
	Ospeed uint32
}

func NewDocumentViewer(path string) *DocumentViewer {
	ext := strings.ToLower(filepath.Ext(path))
	fileType := strings.TrimPrefix(ext, ".")

	return &DocumentViewer{
		path:     path,
		fileType: fileType,
		reader:   bufio.NewReader(os.Stdin),
	}
}

func (d *DocumentViewer) Open() error {
	doc, err := fitz.New(d.path)
	if err != nil {
		return fmt.Errorf("error opening %s: %v", d.fileType, err)
	}
	d.doc = doc

	d.findTextPages()
	if len(d.textPages) == 0 {
		return fmt.Errorf("no pages with extractable text found")
	}
	return nil
}

func (d *DocumentViewer) findTextPages() {
	d.textPages = []int{}
	for i := 0; i < d.doc.NumPage(); i++ {
		text, err := d.doc.Text(i)
		if err == nil && len(strings.Fields(strings.TrimSpace(text))) >= 3 {
			d.textPages = append(d.textPages, i)
		}
	}
}

func (d *DocumentViewer) Run() {
	defer d.doc.Close()

	fmt.Printf("Press any key to start reading %s, or 'q' to quit\n", strings.ToUpper(d.fileType))
	input, _ := d.reader.ReadString('\n')
	if strings.TrimSpace(input) == "q" {
		return
	}

	oldState, err := d.setRawMode()
	if err != nil {
		fmt.Printf("Error setting raw mode: %v\n", err)
		return
	}
	defer d.restoreTerminal(oldState)

	d.currentPage = 0
	for {
		d.displayCurrentPage()
		char := d.readSingleChar()
		if d.handleInput(char) {
			break
		}
	}

	fmt.Println("\nThanks for reading!")
}

func (d *DocumentViewer) handleInput(c byte) bool {
	switch c {
	case 'q':
		return true
	case 'j', ' ':
		if d.currentPage < len(d.textPages)-1 {
			d.currentPage++
		}
	case 'k':
		if d.currentPage > 0 {
			d.currentPage--
		}
	case 'g':
		d.goToPage()
	case 'h':
		d.showHelp()
	}
	return false
}

func (d *DocumentViewer) goToPage() {
	d.restoreTerminal(d.oldState)
	fmt.Printf("\nGo to page (1-%d): ", len(d.textPages))
	line, _ := d.reader.ReadString('\n')
	var num int
	if _, err := fmt.Sscanf(strings.TrimSpace(line), "%d", &num); err == nil {
		if num >= 1 && num <= len(d.textPages) {
			d.currentPage = num - 1
		}
	}
	d.setRawMode()
}
