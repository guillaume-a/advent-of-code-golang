package main

import (
	"aoc"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	// Vérifie s'il y a suffisamment d'arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part-2.go <input>")
		return
	}
    
    // Récupère le contenu du fichier sous forme de tableau
    lines, err := aoc.ReadFile(os.Args[1], true)
    if err != nil {
        fmt.Println(err)
        return
    }
    
    // Ton traitement spécifique à solve-1.go
    currentWeight := 0
    var weight []int

    for _, line := range lines {
        if(line == "") {
            weight = append(weight, currentWeight)
            currentWeight = 0;
            continue;
        }

        weight, err := strconv.Atoi(line)
        if err != nil {
            // Gestion de l'erreur si la conversion échoue
            fmt.Println("Erreur lors de la conversion de la chaîne en nombre entier:", err)
            return
        }
                
        currentWeight += weight
    }

    sort.Sort(sort.Reverse(sort.IntSlice(weight)))

	total := 0
	for _, num := range weight[:3] {
		total += num
	}

    fmt.Println(total)
}
