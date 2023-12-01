package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Reverse(s string) (result string) {
  for _,v := range s {
    result = string(v) + result
  }
  return 
}

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

    numbers := map[string]string{
        "0": "0", "1": "1", "2": "2", "3": "3", "4": "4", "5": "5", "6": "6", "7": "7", "8": "8", "9": "9",
        "zero": "0", "one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9",
    }

    answer := 0

	// Parcours chaque ligne du fichier
	for scanner.Scan() {
		line := scanner.Text()
		

		firstDigit := ""
		lastDigit := ""
		firstIndex := int(^uint(0) >> 1) 
		lastIndex := int(^uint(0) >> 1)

		fmt.Println("--- Line:", line, Reverse(line))

		index := -1

        for k, v := range numbers {
			index = strings.Index(line, k)
		    
			if(index != -1) {
				fmt.Println("Found:", k, index)

				if(index < firstIndex) {
					firstIndex = index
					firstDigit = v
				}
			}

			index = strings.Index(Reverse(line), Reverse(k))
			if(index != -1) {
				fmt.Println("Found:", k, index)

				if(index < lastIndex) {
					lastIndex = index
					lastDigit = v
				}
			}
	    }
		value, err := strconv.Atoi(firstDigit + lastDigit)

    	// Vérifie s'il y a des erreurs pendant la lecture du fichier
        if err != nil {
            panic(err)
        }

		fmt.Println("--- Digits:")
		fmt.Println("First:", firstDigit)
		fmt.Println("Last", lastDigit)
		fmt.Println("Value", value)
		fmt.Println("---")

        answer += value
	}


    fmt.Println(answer)
}

