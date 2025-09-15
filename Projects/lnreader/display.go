package main

import (
	"fmt"
	"strings"
	"unicode"
)

func (d *DocumentViewer) displayCurrentPage() {
	fmt.Print("\033[2J\033[H") // clear screen
	termWidth, termHeight := d.getTerminalSize()
	actualPage := d.textPages[d.currentPage]

	text, err := d.doc.Text(actualPage)
	if err != nil {
		fmt.Printf("Error extracting text: %v\n", err)
		return
	}

	// Reflow text to fit terminal width
	reflowedLines := d.reflowText(text, termWidth)

	reserved := 1 // Only page info
	available := termHeight - reserved

	// Display as many lines as possible within available height
	displayedLines := 0
	for i, line := range reflowedLines {
		if displayedLines >= available {
			break
		}
		fmt.Println(line)
		displayedLines++

		// If this is the last line we have and there's still space, break early
		if i == len(reflowedLines)-1 {
			break
		}
	}

	// Fill remaining space only if needed to reach the bottom
	for displayedLines < available {
		fmt.Println()
		displayedLines++
	}

	// Display page information with file type
	var pageInfo string
	if d.fileType == "epub" {
		pageInfo = fmt.Sprintf("Chapter %d/%d (Actual %d/%d) - EPUB", d.currentPage+1, len(d.textPages), actualPage+1, d.doc.NumPage())
	} else {
		pageInfo = fmt.Sprintf("Page %d/%d (Actual %d/%d) - PDF", d.currentPage+1, len(d.textPages), actualPage+1, d.doc.NumPage())
	}

	if len(pageInfo) > termWidth {
		// Truncate if too long
		pageInfo = pageInfo[:termWidth-3] + "..."
	}

	if len(pageInfo) < termWidth {
		padding := (termWidth - len(pageInfo)) / 2
		fmt.Printf("%s%s", strings.Repeat(" ", padding), pageInfo)
	} else {
		fmt.Print(pageInfo)
	}
}

// reflowText takes raw document text and reflows it to fit the terminal width
func (d *DocumentViewer) reflowText(text string, termWidth int) []string {
	// Clean and normalize the text
	text = strings.ReplaceAll(text, "\r\n", "\n")
	text = strings.ReplaceAll(text, "\r", "\n")

	// For EPUB files, we might want to handle HTML entities or tags
	if d.fileType == "epub" {
		text = d.cleanEpubText(text)
	}

	// Split into paragraphs (double newlines or more indicate paragraph breaks)
	paragraphs := strings.Split(text, "\n\n")

	var reflowedLines []string

	for _, paragraph := range paragraphs {
		if strings.TrimSpace(paragraph) == "" {
			reflowedLines = append(reflowedLines, "")
			continue
		}

		// Clean the paragraph - remove extra whitespace and newlines within paragraph
		cleanParagraph := strings.ReplaceAll(paragraph, "\n", " ")
		cleanParagraph = d.normalizeWhitespace(cleanParagraph)

		if strings.TrimSpace(cleanParagraph) == "" {
			continue
		}

		// Wrap the paragraph to terminal width
		wrappedLines := d.wrapText(cleanParagraph, termWidth)
		reflowedLines = append(reflowedLines, wrappedLines...)

		// Add empty line after paragraph (except for last paragraph)
		reflowedLines = append(reflowedLines, "")
	}

	// Remove trailing empty lines
	for len(reflowedLines) > 0 && reflowedLines[len(reflowedLines)-1] == "" {
		reflowedLines = reflowedLines[:len(reflowedLines)-1]
	}

	return reflowedLines
}

// cleanEpubText removes common HTML entities and tags that might appear in EPUB text
func (d *DocumentViewer) cleanEpubText(text string) string {
	// Replace common HTML entities
	replacements := map[string]string{
		"&nbsp;":  " ",
		"&amp;":   "&",
		"&lt;":    "<",
		"&gt;":    ">",
		"&quot;":  "\"",
		"&apos;":  "'",
		"&#8217;": "'",
		"&#8220;": "\"",
		"&#8221;": "\"",
		"&#8230;": "...",
		"&#8212;": "—",
		"&#8211;": "–",
	}

	for entity, replacement := range replacements {
		text = strings.ReplaceAll(text, entity, replacement)
	}

	return text
}

// normalizeWhitespace replaces multiple spaces/tabs with single spaces
func (d *DocumentViewer) normalizeWhitespace(text string) string {
	var result strings.Builder
	var lastWasSpace bool

	for _, r := range text {
		if unicode.IsSpace(r) {
			if !lastWasSpace {
				result.WriteRune(' ')
				lastWasSpace = true
			}
		} else {
			result.WriteRune(r)
			lastWasSpace = false
		}
	}

	return strings.TrimSpace(result.String())
}

// wrapText wraps text to specified width using word wrapping
func (d *DocumentViewer) wrapText(text string, width int) []string {
	if width <= 0 {
		return []string{text}
	}

	words := strings.Fields(text)
	if len(words) == 0 {
		return []string{""}
	}

	var lines []string
	var currentLine strings.Builder

	for _, word := range words {
		// Check if adding this word would exceed the width
		proposedLength := currentLine.Len()
		if proposedLength > 0 {
			proposedLength += 1 // for the space
		}
		proposedLength += len(word)

		if proposedLength <= width || currentLine.Len() == 0 {
			// Add word to current line
			if currentLine.Len() > 0 {
				currentLine.WriteString(" ")
			}
			currentLine.WriteString(word)
		} else {
			// Start new line with this word
			if currentLine.Len() > 0 {
				lines = append(lines, currentLine.String())
				currentLine.Reset()
			}
			currentLine.WriteString(word)
		}
	}

	// Add the last line if it has content
	if currentLine.Len() > 0 {
		lines = append(lines, currentLine.String())
	}

	return lines
}

func (d *DocumentViewer) showHelp() {
	fmt.Print("\033[2J\033[H") // clear screen
	termWidth, _ := d.getTerminalSize()

	fmt.Println(strings.Repeat("=", termWidth))
	fmt.Printf("%s Viewer Help\n", strings.ToUpper(d.fileType))
	fmt.Println(strings.Repeat("=", termWidth))
	fmt.Println()
	fmt.Println("Navigation:")
	fmt.Println("  j or Space  - Next page/chapter")
	fmt.Println("  k           - Previous page/chapter")
	fmt.Println("  g           - Go to specific page/chapter")
	fmt.Println("  h           - Show this help")
	fmt.Println("  q           - Quit")
	fmt.Println()
	fmt.Println("Features:")
	fmt.Println("  - Text is reflowed to fit terminal width")
	fmt.Println("  - Only pages/chapters with extractable text are shown")
	fmt.Println("  - Paragraphs are preserved with proper spacing")
	if d.fileType == "epub" {
		fmt.Println("  - HTML entities are converted to readable text")
	}
	fmt.Println()
	fmt.Println("Supported formats: PDF, EPUB")
	fmt.Println()
	fmt.Println(strings.Repeat("=", termWidth))
	fmt.Println("Press any key to return...")
	d.readSingleChar()
}
