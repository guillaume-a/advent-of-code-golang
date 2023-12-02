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

	reGameColors := regexp.MustCompile(`(?:([0-9]+) (red|green|blue),?)+`)

	// Parcours chaque ligne du fichier
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		gameColors := reGameColors.FindAllStringSubmatch(line, -1)
		
		fewer := map[string]int{
			"red": 0, 
			"green": 0, 
			"blue": 0,
    	}

		for _, v := range gameColors {
			fmt.Println(v[1], v[2])
			count, _ := strconv.Atoi(v[1]);

			if count > fewer[v[2]] {
				fewer[v[2]] = count
			}
		}

		gameScore := 1
		for k, v := range fewer {
			fmt.Println(k, v)
			gameScore *= v
		}

		fmt.Println("gameScore:", gameScore)

		answer += gameScore
		
	}

    fmt.Println("Answer:", answer)
}



