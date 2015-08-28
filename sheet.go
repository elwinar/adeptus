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
	if scanner.Err() != nil {
		return sheet, fmt.Errorf("Error during scan.")
	}
	
	// Parse header
	if len(buffer) < 1 {
		return sheet, fmt.Errorf("Empy file. Expecting at least a complete header."), 
	}
	h, err := ParseHeader(buffer[0])
	if err != nil {
		return sheet, err
	}
	sheet.Header = h
	
	// Parse sessions
	if len(buffer) < 2 {
		return sheet, nil
	}
	buffer = buffer[1:]
	for i, block := range buffer {
		s, err := ParseSession(block)
		if err != nil {
			return Sheet{}, err
		}
		sheet.Sessions = append(sheet.Sessions, s)
	}
	return sheet, nil

}
