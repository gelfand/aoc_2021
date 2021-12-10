package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Op struct {
	opcode Opcode
	units  int
}

func scanfile(filename string) ([]Op, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var ops []Op
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		splitted := strings.Split(scanner.Text(), " ")
		opcode := OpcodeFromString(splitted[0])
		units, err := strconv.Atoi(splitted[1])
		if err != nil {
			return nil, fmt.Errorf("unable to scan units, line: %s, err: %w", scanner.Text(), err)
		}

		op := Op{
			opcode: opcode,
			units:  units,
		}
		ops = append(ops, op)
	}

	return ops, nil
}

func main() {
	ops, err := scanfile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	type state struct {
		pos   int
		aim   int
		depth int
	}

	s := state{
		pos:   0,
		aim:   0,
		depth: 0,
	}

	for _, op := range ops {
		switch op.opcode {
		case down:
			s.aim += op.units
		case up:
			s.aim -= op.units
		case forward:
			s.pos += op.units
			s.depth += s.aim * op.units
		}
	}

	fmt.Println(s.pos * s.depth)
}
