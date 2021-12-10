package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	s := struct {
		x int
		y int
	}{
		x: 0,
		y: 0,
	}

	for scanner.Scan() {
		text := scanner.Text()
		splitted := strings.Split(text, " ")

		op := OpcodeFromString(splitted[0])
		units, err := strconv.Atoi(splitted[1])
		if err != nil {
			log.Fatal(err)
		}

		switch op {
		case forward:
			s.x += units
		case down:
			s.y += units
		case up:
			s.y -= units
		}
	}

	fmt.Println(s.x * s.y)
}
