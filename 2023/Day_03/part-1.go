package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// Vérifie s'il y a suffisamment d'arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part-1.go <input>")
		return
	}
    
    // Ouvre le fichier
	file, err := os.Open(os.Args[1])
	if err != nil {
        panic(err)
	}
	defer file.Close()

	// Crée un scanner pour lire le fichier ligne par ligne
	scanner := bufio.NewScanner(file)

    answer := 0
	width := 0
	// motor := ""

	var lines []string

	// Parcours chaque ligne du fichier
	for scanner.Scan() {
		line := scanner.Text()
		width = len(line)
		lines = append(lines, line + ".")
		fmt.Println(line, width)
	}

	inNumber := false
	start := 0
	end := 0

	for row, line := range lines {
	
		//fmt.Println(row)
		inNumber = false

		for col, char := range line[:] {
			isDigit := regexp.MustCompile(`\d`).MatchString(string(char))
			if(isDigit) {
				//First digit
				if(!inNumber) {
					start = col
					//fmt.Println("Start number:", col)
				}
			} else {
				// Last digit was found previously
				if(inNumber) {
					end = col
					size := end-start
					number, _ := strconv.Atoi(string(line[start:end]))

					fmt.Println("Number:", number, start, end)

					// Get borders
					top := ""
					left := ""
					right := ""
					bottom := ""
					
					left += string(line[max(0,start-1):start])
					right += string(line[start+size:min(width,end+1)])
	

					if(row > 0) {
						top += string(lines[row-1][max(0,start-1):min(width,end+1)])
					}

					if(row < len(lines) - 1) {
						top += string(lines[row+1][max(0,start-1):min(width,end+1)])
					}

					border := top + left + right + bottom;

					// fmt.Println("Left:", left)
					// fmt.Println("Top:", top)
					// fmt.Println("Right:", right)
					// fmt.Println("Bottom:", bottom)
			
					fmt.Println("Border:", string(border))

					// Clean border to find only symbols (not . or digit)
					hasSymbol := regexp.MustCompile(`[\d\.]`)
					border = hasSymbol.ReplaceAllString(border, "")
					
					fmt.Println("Border:", string(border))

					// Add if symbol found
					if(len(border) > 0) {
					 	answer += number
					}
				}
			}

			inNumber = isDigit
		}

	}

/*

. . + + + + + . . . / . . + 1 1 1 + . . . / . . + + + + + . . .

width = 10
start = 13
len = 3

check : [
	start - 1 // Juste avant
	start + len // Juste apres

	start - width - 1 -> start - width + len
	start + width - 1 -> start + width + len
]

*/

    fmt.Println("Answer:", answer)
}
