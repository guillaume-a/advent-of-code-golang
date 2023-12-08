package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	reData := regexp.MustCompile(`(?:\s+)(\d+)`)
	
	var times []int
	var distances []int
	

	for scanner.Scan() {
		line := scanner.Text()

		values := reData.FindAllStringSubmatch(line, -1)
		fmt.Println(values)

		var raw []int

		for _, value := range(values) {
			n, _ := strconv.Atoi(value[1])
			raw = append(raw, n)
		}
		fmt.Println(raw)
		
		if strings.HasPrefix(line, "Time:") {
			// Start new mapping recording
			fmt.Println("Time")

			times = raw
		}

		if strings.HasPrefix(line, "Distance:") {
			// Start new mapping recording
			fmt.Println("Distance")

			distances = raw
		}
	}
	
	fmt.Println(times, distances)

	answer := 1

	for i, time := range(times) {
		fmt.Println("Race", i, "Time", time, "Distance", distances[i])

		wins := 0
		for push := 1; push < time; push ++ {
			fmt.Println("Try", push, "Distance", push * (time - push))
			if push * (time - push) > distances[i] {
				wins ++
			}
		}

		answer *= wins
	}

	fmt.Println("Answer:", answer)
}
