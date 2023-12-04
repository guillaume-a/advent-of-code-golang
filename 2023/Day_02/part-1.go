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

	reGameNr := regexp.MustCompile(`Game ([0-9]+):`)
	reGameColors := regexp.MustCompile(`(?:([0-9]+) (red|green|blue),?)+`)

	// Parcours chaque ligne du fichier
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		gameColors := reGameColors.FindStringSubmatch(line, -1)
		possible := true

		for _, v := range gameColors {
			fmt.Println(v[1], v[2])
			count, _ := strconv.Atoi(v[1]);

			switch(v[2]) {
			case "red":
				possible = count <= 12
				break
			case "green":
				possible = count <= 13
				break
			case "blue":
				possible = count <= 14
				break								
			}

			if !possible {
				break
			}
		}

		if possible {
			gameNr := reGameNr.FindStringSubmatch(line)
			fmt.Println("Game: ", gameNr[1])

			count, _ := strconv.Atoi(gameNr[1]);
			answer += count
		}
	}

    fmt.Println("Answer:", answer)
}



