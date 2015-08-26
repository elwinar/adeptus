package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const ()

func main() {
	filename := "test.txt"

	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Unable to open file %s. Please ensure the file exists and is readable.\n", filename)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		if len(strings.TrimSpace(text)) == 0 {
			continue
		}
		l := NewLine(text)
		s := rawSession{
			date:       l.GetDate(),
			experience: l.GetXp(),
			label:      l.GetLabel(),
		}
		fmt.Println(s)
		for scanner.Scan() {
			text = scanner.Text()
			if len(strings.TrimSpace(text)) == 0 {
				break
			}
			l := NewLine(text)
			u := rawUpgrade{
				mark:       l.GetMark(),
				experience: l.GetXp(),
				value:      l.GetValue(),
				label:      l.GetLabel(),
			}
			fmt.Println(u.Format())
		}

	}
	if scanner.Err() != nil {
		fmt.Println(err)
	}
}
