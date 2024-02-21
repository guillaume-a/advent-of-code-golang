package main

import (
	"bufio"
	"fmt"
	"math"
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
	
	var time int
	var distance int
	
	for scanner.Scan() {
		line := scanner.Text()

		values := reData.FindAllStringSubmatch(line, -1)
		fmt.Println(values)

		raw := ""

		for _, value := range(values) {
			raw += value[1]
		}

		fmt.Println(raw)
		
		if strings.HasPrefix(line, "Time:") {
			// Start new mapping recording
			fmt.Println("Time")

			time, _ = strconv.Atoi(raw)
		}

		if strings.HasPrefix(line, "Distance:") {
			// Start new mapping recording
			fmt.Println("Distance")

			distance, _ = strconv.Atoi(raw)
		}
	}
	
	fmt.Println(time, distance)

	// https://www.maxicours.com/se/cours/resoudre-une-equation-du-second-degre---terminale/
	// Delta
	a := float64(-1)
	b := float64(time)
	c := float64(-distance)


	// Debug
	// a = float64(-4)
	// b = float64(-3)
	// c = float64(10)

	delta := (b * b) - (4 * a * c)

	fmt.Println("Solve", a, b, c, delta)

	x1 := (-b - math.Sqrt(delta)) / (2 * a)
	x2 := (-b + math.Sqrt(delta)) / (2 * a)
	
	fmt.Println("Solutions", x1, x2)

	answer := math.Abs(math.Ceil(x1) - math.Ceil(x2))


	fmt.Println("Answer:", answer)
}
