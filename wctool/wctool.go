package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	argsWithProg := os.Args
	if argsWithProg[0] == "ccwc" {
		if (len(argsWithProg) == 3 && (argsWithProg[1] == "-c") || (argsWithProg[1] == "-l") || (argsWithProg[1] == "-w") || (argsWithProg[1] == "-m")) || len(argsWithProg) == 2 {
			data, err := os.Open("./" + argsWithProg[len(argsWithProg)-1])
			if err != nil {
				fmt.Println("Failed to open data")
				return
			}
			reader := bufio.NewReader(data)

			byteCount := 0
			wordCount := 0
			lineCount := 0
			charCount := 0

			insideWord := false
			for {
				char, size, err := reader.ReadRune()
				if err != nil {
					if err == io.EOF {
						break // Terminate the loop if end of file is reached
					} else {
						fmt.Println("Failed to read data")
						return
					}
				}
				byteCount += size
				charCount++
				if unicode.IsSpace(char) {
					insideWord = false
					if char == '\n' {
						lineCount++
					}
				} else {
					if !insideWord {
						wordCount++
						insideWord = true
					}
				}
			}

			if len(argsWithProg) == 2 {
				fmt.Println(lineCount, wordCount, byteCount, argsWithProg[1])
			} else if argsWithProg[1] == "-c" {
				//chars / bytes
				fmt.Println(byteCount, argsWithProg[2])
			} else if argsWithProg[1] == "-w" {
				// words
				fmt.Println(wordCount, argsWithProg[2])
			} else if argsWithProg[1] == "-l" {
				//lines
				fmt.Println(lineCount, argsWithProg[2])
			} else if argsWithProg[1] == "-m" {
				//lines
				fmt.Println(charCount, argsWithProg[2])
			}
		}
	}
}
