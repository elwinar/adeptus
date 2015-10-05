package main

import (
	"bufio"
	"fmt"
	"io"
)

// Sheet holds the informations of the character sheet: the character definition
// in the header, and the history of the character in the sessions.
type Sheet struct {
	Header          Header
	Sessions        []Session
	Characteristics Characteristics
}

// ParseSheet parse a Sheet from a io.Reader.
func ParseSheet(file io.Reader) (Sheet, error) {
	scanner := bufio.NewScanner(file)
	buffer := [][]line{}
	block := []line{}

	// Scan each line.
	i := 0
	for scanner.Scan() {
		i++
		l := line{
			Number: i,
			Text:   scanner.Text(),
		}

		// If line is comment, loop.
		if l.IsComment() {
			continue
		}

		// If line is not empty, stock it.
		if !l.IsEmpty() {
			block = append(block, l)
			continue
		}

		// At this point, the script has reached an empty line,
		// which means the block is ready to be processed.
		// If the block is not empty, append it to the buffer and empty it.
		if len(block) != 0 {
			buffer = append(buffer, block)
			block = []line{}
		}
	}

	// In case of error, return now
	if scanner.Err() != nil {
		panic(fmt.Sprintf("unable to read sheet: %s", scanner.Err()))
	}

	// Append the last block to the buffer
	if len(block) != 0 {
		buffer = append(buffer, block)
	}

	// Check there is at least two blocks
	if len(buffer) < 2 {
		return Sheet{}, NewError(InvalidCharacterSheet)
	}

	// Parse the first block as header
	header, err := parseHeader(buffer[0])
	if err != nil {
		return Sheet{}, err
	}

	// Parse the second block as Characteristics
	characteristics, err := parseCharacteristics(buffer[1])
	if err != nil {
		return Sheet{}, err
	}

	// Parse the other blocks as sessions
	sessions := []Session{}
	for _, block := range buffer[2:] {
		session, err := parseSession(block)
		if err != nil {
			return Sheet{}, err
		}

		sessions = append(sessions, session)
	}

	return Sheet{
		Header:          header,
		Sessions:        sessions,
		Characteristics: characteristics,
	}, nil
}
