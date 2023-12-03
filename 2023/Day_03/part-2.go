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
		fmt.Println("Usage: go run part-2.go <input>")
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

	gears := make(map[int]int)

	for row, line := range lines {
	
		inNumber = false

		for col, char := range line[:] {
			isDigit := regexp.MustCompile(`\d`).MatchString(string(char))
			if(isDigit) {
				//First digit
				if(!inNumber) {
					start = col
				}
			} else {
				// Last digit was found previously
				if(inNumber) {
					end = col
					//size := end-start
					number, _ := strconv.Atoi(string(line[start:end]))

					fmt.Println(number, start, end, row)

					for x:=max(0,start-1); x < min(width,end+1); x++ {
						for y:= max(0,row-1); y <= min(len(lines)-1,row+1); y++ {
							fmt.Println(y, x, string(lines[y][x]))

							if(string(lines[y][x]) == "*") {
								fmt.Println("Number:", number, start, end)

								key := width * y + x
								fmt.Println("Key", key)

								// Did we found the first operand ?
								first, exists := gears[key];
								if exists {
									fmt.Println("Exists in map", first)
									answer += first * number
								} else {
									fmt.Println("Wait for second", first)
									gears[key] = number
								}
							}

						}
					}
					fmt.Println("============")
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
