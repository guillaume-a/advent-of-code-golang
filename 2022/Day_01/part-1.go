package main

import (
	"aoc"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Vérifie s'il y a suffisamment d'arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part-1.go <input>")
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
    maxWeight := 0

    for _, line := range lines {
        if(line == "") {
            currentWeight = 0;
            continue;
        }

        weight, err := strconv.Atoi(line)
        if err != nil {
            // Gestion de l'erreur si la conversion échoue
            fmt.Println("Erreur lors de la conversion de la chaîne en nombre entier:", err)
            return
        }
                
        // fmt.Println(weight)
        currentWeight += weight

        if(currentWeight > maxWeight) {
            maxWeight = currentWeight;
        }
    }

    fmt.Println(maxWeight)
}
