package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func gamma_rate(arr [12]int) int {
	var gamma_rate int
	for pos, el := range arr {
		if el > 0 {
			gamma_rate += 1 * int(math.Pow(2, float64(12-(pos+1))))
		} else {
			gamma_rate += 0 * int(math.Pow(2, float64(12-(pos+1))))
		}
	}

	return gamma_rate
}

func gamma_epsilon(arr [12]int) int {
	var epsilon int
	for pos, el := range arr {
		if el > 0 {
			epsilon += 0 * int(math.Pow(2, float64(12-(pos+1))))
		} else {
			epsilon += 1 * int(math.Pow(2, float64(12-(pos+1))))
		}
	}

	return epsilon
}

func main() {
	file, err := os.Open("./values")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var statut [12]int

	for scanner.Scan() {
		for pos, char := range scanner.Text() {
			switch char {
			case '0':
				statut[pos]--
			case '1':
				statut[pos]++
			}
		}
	}

	epsilon := gamma_epsilon(statut)
	rate := gamma_rate(statut)

	fmt.Println("Submarine:", epsilon*rate)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
