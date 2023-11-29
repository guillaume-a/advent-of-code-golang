package aoc

import (
	"bufio"
	"os"
)

func ReadFile(fileName string, addNewLine bool) ([]string, error) {
	// Ouvre le fichier en lecture
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Crée un scanner pour lire le fichier ligne par ligne
	scanner := bufio.NewScanner(file)

	var lines []string

	// Parcours chaque ligne du fichier
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	// Vérifie s'il y a des erreurs pendant la lecture du fichier
	if err := scanner.Err(); err != nil {
		return nil, err
	}

    if(addNewLine) {
        lines = append(lines, "")
    }

	return lines, nil
}
