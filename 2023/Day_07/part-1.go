package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part-1.go <input>")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	answer := 1


	power := map[string]int{
		"A": 14,
		"K": 13,
		"Q": 12,
		"J": 11,
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
	}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		hand := line[0]
		bid, _ := strconv.Atoi(line[1])

		cards := make(map[string]int, 5)

		// Count cards for this hand
		for _, card := range(hand) {
			cards[string(card)] ++
		}

		fmt.Println(hand, cards, bid)

		full := false
		pairs := 0;

		for card, count := range(cards) {
			if count >= 4 {
				answer += count * bid
			}
			if count == 3 {
				full = true
			}
			if count == 2 {
				pairs += count * bid
			}
		}
/////////////////
		if(pairs > 0) {

		} else {
			answer += count * bid
		}


		if(full) {
			if(pairs > 0) {

			} else {

			}
		} else {
			answer += pairs
		}

	}


	fmt.Println("Answer:", answer)
}
