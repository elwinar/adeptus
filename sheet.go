package adeptus

import "io"

type Sheet struct {
	Header   map[string]string
	Sessions []Session
}

func ParseSheet(file *io.Reader) (Sheet, error) {
	
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
			Text: scanner.Text(),
		}
		
		// If line is valid, push it to the block and loop
		if line.isValid() {
			block = append(block, line)
			continue
		}
		
		// If the block is not empty, append it to the buffer
		if len(block) != 0 {
			buffer = append(buffer, block)
		}
		
		// Empty the block
		block = []Line{}
		
	}
	if scanner.Err() != nil {
		return Sheet{}, err
	}
	
	for
}
