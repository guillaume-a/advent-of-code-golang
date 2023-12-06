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

	var seeds []int
	var nextMapping map[int]int

	reMap := regexp.MustCompile(`to-([\w]*) map:$`)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			// Store new seed locations
			if nextMapping == nil {
				continue
			}

			fmt.Println("Seeds", nextMapping)

			// Add all values not mapped
			for _, seed := range seeds {
				_, found := nextMapping[seed]
				if !found {
					nextMapping[seed] = seed
				}
			}

			fmt.Println("Seeds remapped", nextMapping)

			i := 0
			seeds = make([]int, len(seeds))
			for _, seed := range nextMapping {
				fmt.Println("new seeds", i, seed)
				seeds[i] = seed
				i++
			}

			continue
		}

		if regexp.MustCompile(`^seeds: `).MatchString(line) {
			// Load seeds

			data := strings.Split(line[7:], " ")
			seeds = make([]int, len(data))

			// Store them in int
			for i, v := range data {
				seed, _ := strconv.Atoi(v)
				seeds[i] = seed

			}

			fmt.Println("Detect seeds", seeds)
			continue
		}

		if reMap.MatchString(line) {
			// Start new mapping recording
			fmt.Println("Reset mapping")
			nextMapping = make(map[int]int)

			toMap := reMap.FindStringSubmatch(line)
			fmt.Println("Start mapping", string(toMap[1]), nextMapping)

			continue
		}

		// Store mapping data
		fmt.Println("Store and process map")

		mapping := strings.Split(line, " ")
		fmt.Println(mapping)

		for _, seed := range seeds {
			sourceRange, _ := strconv.Atoi(mapping[1])
			lengthRange, _ := strconv.Atoi(mapping[2])
			destinationRange, _ := strconv.Atoi(mapping[0])

			if seed >= sourceRange && seed < sourceRange+lengthRange {
				fmt.Println("Process", seed, sourceRange, lengthRange, destinationRange)
				fmt.Println("Found", seed, destinationRange+seed-sourceRange)
				nextMapping[seed] = destinationRange + seed - sourceRange
			}
		}
	}

	//Find lowers

	answer := int(^uint(0) >> 1) 

	// Add all values not mapped
	for _, seed := range seeds {
		value, found := nextMapping[seed]
		if !found {
			answer = min(answer, seed)
		} else {
			answer = min(answer, value)
		}
	}

	fmt.Println("Answer:", answer)
}
