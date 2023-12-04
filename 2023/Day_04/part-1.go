package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strings"
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

	// Parcours chaque ligne du fichier
	for scanner.Scan() {
		card := scanner.Text()
		re := regexp.MustCompile(`Card ([\d\s]+): ([\d\s]+) \| ([\d\s]+)`);

		game := re.FindStringSubmatch(card)

		//fmt.Println("Card:", game[1], "Found:", game[2], "Winning:", game[3])

		found := strings.Split(game[2], " ")
		winning := strings.Split(game[3], " ")

		fmt.Println(found, winning)
		
		matches := 0

		for _, n := range found {
			if n == "" {
				continue
			}

			check := slices.Contains(winning, n)
			fmt.Println("Test:", n, check)

			if check {
				matches ++
			}
		}

		fmt.Println(matches, "number match")

		if matches > 0 {
			score := int(math.Pow(2, float64(matches-1)))
			fmt.Println(score, "points")
			answer += score
		}
		
	}

    fmt.Println("Answer:", answer)
}
