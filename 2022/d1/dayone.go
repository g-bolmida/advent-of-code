package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var activeElf, elfTotals, elfPodium []int

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		if line != "" {
			lineInt, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			activeElf = append(activeElf, lineInt)
		} else {
			elfSum()
		}
	}

	elfCounter := 0
	for elfCounter != 3 {
		max := 0
		index := 0
		for i, val := range elfTotals {
			if val > max {
				max = val
				index = i
			}
		}
		elfPodium = append(elfPodium, max)
		elfTotals[index] = 0
		elfCounter++
	}

	fmt.Printf("the top elf is carrying %v calories\n\n", elfPodium[0])
	fmt.Printf("the second elf is carrying %v calories\n\n", elfPodium[1])
	fmt.Printf("the third elf is carrying %v calories\n\n", elfPodium[2])
	podiumTotal := elfPodium[0] + elfPodium[1] + elfPodium[2]
	fmt.Printf("the total of the top three is %v calories\n\n", podiumTotal)
}

func elfSum() {
	elfTotal := 0
	for _, val := range activeElf {
		elfTotal += val
	}
	elfTotals = append(elfTotals, elfTotal)
	activeElf = nil
}
