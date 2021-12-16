package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func calculate_decimal(arr [12]int, zero int, one int) int {
	var dec int
	for pos, el := range arr {
		zo := zero
		if el > 0 {
			zo = one
		}

		dec += zo * int(math.Pow(2, float64(12-(pos+1))))
	}

	return dec
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

	epsilon := calculate_decimal(statut, 1, 0)
	rate := calculate_decimal(statut, 0, 1)

	fmt.Println("Submarine:", epsilon*rate)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
