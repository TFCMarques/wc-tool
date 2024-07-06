package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func countLines(data []byte) int {
	buffer := bytes.NewBuffer(data)
	scanner := bufio.NewScanner(buffer)

	lines := 0

	for scanner.Scan() {
		lines++
	}

	return lines
}

func countWords(data []byte) int {
	words := 0
	atWord := false

	for _, b := range data {
		if unicode.IsSpace(rune(b)) {
			if atWord {
				words++
				atWord = false
			}
		} else {
			atWord = true
		}
	}

	if atWord {
		words++
	}

	return words
}

func countChars(data []byte) int {
	return utf8.RuneCount(data)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ccwc <option> [<filename>]")
		return
	}

	var data []byte
	var err error
	var filename string

	if len(os.Args) == 2 && os.Args[1][0] != '-' {
		filename = os.Args[1]
		data, err = os.ReadFile(filename)
	} else if len(os.Args) > 2 {
		filename = os.Args[2]
		data, err = os.ReadFile(filename)
	} else {
		filename = ""
		data, err = io.ReadAll(os.Stdin)
	}

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	option := ""

	if len(os.Args) > 1 {
		option = os.Args[1]
	}

	switch option {
	case "-c":
		fmt.Printf("%d %s\n", len(data), filename)
	case "-l":
		lines := countLines(data)
		fmt.Printf("%d %s\n", lines, filename)
	case "-w":
		words := countWords(data)
		fmt.Printf("%d %s\n", words, filename)
	case "-m":
		chars := countChars(data)
		fmt.Printf("%d %s\n", chars, filename)
	case "":
		lines := countLines(data)
		words := countWords(data)
		fmt.Printf("%d %d %d %s\n", lines, words, len(data), filename)
	default:
		fmt.Println("Unknown option:", option)
	}

}
