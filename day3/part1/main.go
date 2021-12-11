package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/bits"
	"os"
	"strconv"
)

// scanInput scans input of base 2 integers, returns slice of []uint64, along with bitsLength.
func scanInput(r io.Reader) ([]uint64, int, error) {
	var (
		input      []uint64
		bitsLength int
		done       bool
	)
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		if !done {
			bitsLength = len(scanner.Text())
			done = true
		}
		// break this loop if we have empty string.
		if scanner.Text() == "" {
			break
		}
		v, err := strconv.ParseUint(scanner.Text(), 2, 64)
		if err != nil {
			return nil, 0, fmt.Errorf("unable to scan input: %w", err)
		}

		input = append(input, v)
	}

	return input, bitsLength, nil
}

func main() {
	input, bitsLength, err := scanInput(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	bitMap := make([][]int, len(input))
	for i := range input {
		bitMap[i] = make([]int, bitsLength)
		for j := 0; j < bitsLength; j++ {
			bitMap[i][bitsLength-j-1] = bit(input[i], j+1)
		}
	}

	gammaRate := ""
	epsilonRate := ""
	rho := len(bitMap[0])
	for i := 0; i < rho; i++ {
		zeroes := 0
		ones := 0
		for j := 0; j < len(bitMap); j++ {
			if bitMap[j][i] == 0 {
				zeroes++
			} else {
				ones++
			}
		}

		if ones > zeroes {
			gammaRate += "1"
			epsilonRate += "0"
		} else {
			gammaRate += "0"
			epsilonRate += "1"
		}
	}

	gammaRateV, err := strconv.ParseUint(gammaRate, 2, 64)
	if err != nil {
		log.Fatalf("Unable to parse gammaRate: %v", err)
	}
	epsilonRateV, err := strconv.ParseUint(epsilonRate, 2, 64)
	if err != nil {
		log.Fatalf("Unable to parse epsilonRate: %v", err)
	}

	fmt.Println(gammaRateV * epsilonRateV)
}

// bit returns n'th bit of the val.
func bit(val uint64, n int) int {
	if int(val&(1<<(n-1))) == 0 || bits.Len64(val) < n {
		return 0
	}
	return 1
}
