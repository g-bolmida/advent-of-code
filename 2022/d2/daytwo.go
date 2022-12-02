package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var points int
	var points2 int

	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		plays := strings.Split(scanner.Text(), " ")
		switch plays[0] {
		case "A":
			if plays[1] == "X" {
				points = points + 4
				points2 = points2 + 3
			} else if plays[1] == "Y" {
				points = points + 8
				points2 = points2 + 4
			} else if plays[1] == "Z" {
				points = points + 3
				points2 = points2 + 8
			}
		case "B":
			if plays[1] == "X" {
				points = points + 1
				points2 = points2 + 1
			} else if plays[1] == "Y" {
				points = points + 5
				points2 = points2 + 5
			} else if plays[1] == "Z" {
				points = points + 9
				points2 = points2 + 9
			}
		case "C":
			if plays[1] == "X" {
				points = points + 7
				points2 = points2 + 2
			} else if plays[1] == "Y" {
				points = points + 2
				points2 = points2 + 6
			} else if plays[1] == "Z" {
				points = points + 6
				points2 = points2 + 7
			}
		default:
		}
	}

	fmt.Printf("the score without encryption is %v and with the decrypted second column is %v\n", points, points2)
}
