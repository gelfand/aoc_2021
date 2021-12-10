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

	answer := 0
	var state *int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		v, err := strconv.Atoi(scanner.Text())
		if err != nil {
			continue
		}

		if state == nil {
			state = &v
			continue
		}

		if *state < v {
			answer++
		}

		state = &v
	}

	fmt.Println(answer)
}
