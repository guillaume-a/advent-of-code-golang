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

	// Parcours chaque ligne du fichier
	for scanner.Scan() {
		line := scanner.Text()

        re := regexp.MustCompile(`[^0-9]`)

        number := re.ReplaceAll([]byte(line), []byte(""))
        firstDigit := string(number[0])
        lastDigit := string(number[len(number)-1])

	    value, err := strconv.Atoi(firstDigit + lastDigit)

        // Vérifie s'il y a des erreurs pendant la lecture du fichier
        if err != nil {
            panic(err)
        }

        answer += value
	}


    fmt.Println(answer)
}

