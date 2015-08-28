package adeptus

import (
	"bufio"
	"fmt"
	"io"
)

type Sheet struct {
	Header   Header
	Sessions []Session
}

func ParseSheet(file io.Reader) (Sheet, error) {

	sheet := Sheet{}

	scanner := bufio.NewScanner(file)
	buffer := [][]Line{}
	block := []Line{}

	// Scan each line
	i := 0
	for scanner.Scan() {
		i++
		line := Line{
			Number: i,
			Text:   scanner.Text(),
		}

		// If line is comment, loop
		if line.IsComment() {
			continue
		}

		// If line is not empty, stock it
		if line.IsEmpty() {
			block = append(block, line)
			continue
		}

		// At this point, the script has reached an empty line,
		// which means the block is ready to be processed.
		// If the block is not empty, append it to the buffer and empty it
		if len(block) != 0 {
			buffer = append(buffer, block)
			block = []Line{}
		}
	}

	// In case of error, return now
	if scanner.Err() != nil {
		return sheet, fmt.Errorf("error while reading the sheet: %s", scanner.Err())
	}
	
	// Check there is at least one block
	if len(buffer) == 0 {
		return sheet, fmt.Errorf("invalid sheet: sheet should contain at least a complete header") 
	}
	
	// Parse the header
	h, err := ParseHeader(buffer[0])
	if err != nil {
		return sheet, fmt.Errorf("unable to parse sheet: %s", err)
	}
	sheet.Header = h

	// Remove the header block from the buffer
	buffer = buffer[1:]

	// For each remaining block, parse it as a session block
	for _, block := range buffer {
		s, err := ParseSession(block)
		if err != nil {
			return Sheet{}, fmt.Errorf("unable to parse sheet: %s", err)
		}
		sheet.Sessions = append(sheet.Sessions, s)
	}
	return sheet, nil
}

