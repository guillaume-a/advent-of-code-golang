package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
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

	multipliers := make(map[int]int)

	// Parcours chaque ligne du fichier
	for scanner.Scan() {
		card := scanner.Text()
		re := regexp.MustCompile(`Card (?:[\s]*)([\d]+): ([\d\s]+) \| ([\d\s]+)`);

		game := re.FindStringSubmatch(card)

		id, _ := strconv.Atoi(game[1])
		found := strings.Split(game[2], " ")
		winning := strings.Split(game[3], " ")

		multipliers[id] ++

		fmt.Println(id, found, winning)
		
		matches := 0

		for _, n := range found {
			if n == "" {
				continue
			}

			check := slices.Contains(winning, n)
			//fmt.Println("Test:", n, check)

			if check {
				matches ++
			}
		}

		fmt.Println(matches, "number match")

		if matches > 0 {
			for i:= id+1; i < id+1+matches; i++ {
				multipliers[i] += multipliers[id]
			}
			
			//fmt.Println(multipliers)
		}
	}

	fmt.Println(multipliers)

	for _, count := range multipliers {
		answer += count
	}

    fmt.Println("Answer:", answer)
}
