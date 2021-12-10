package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var input []int
	for scanner.Scan() {
		if scanner.Text() == "\n" {
			break
		}

		v, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, v)
	}

	answer := 0
	latestWindow := 0
	for i := 0; i < len(input)-2; i++ {
		window := input[i] + input[i+1] + input[i+2]
		if i == 0 {
			latestWindow = window
			continue
		}

		if window > latestWindow {
			answer++
		}
		latestWindow = window
	}

	fmt.Println(answer)
}
