package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fichier, err := os.Open("input/jours1.txt")
	if err != nil {
		log.Fatalf("Erreur en ouvrant le fichier : %v", err)
	}
	defer fichier.Close()

	scanner := bufio.NewScanner(fichier)

	var idsLeft []int
	var idsRight []int
	for scanner.Scan() {
		ids := strings.Split(scanner.Text(), "   ")

		idLeft, _ := strconv.Atoi(ids[0])
		idRight, _ := strconv.Atoi(ids[1])
		idsLeft = append(idsLeft, idLeft)
		idsRight = append(idsRight, idRight)
	}

	// Part 2
	var resultPart2 int = 0
	for _, idLeft := range idsLeft {
		occurence := 0
		for _, idRight := range idsRight {
			if idLeft == idRight {
				occurence++
			}
		}
		resultPart2 += idLeft * occurence
	}

	// Part 1
	var resultPart1 int = 0
	for len(idsLeft) > 0 {
		minLeft := math.MaxInt32
		minRight := math.MaxInt32
		idMinLeft := -1
		idMinRight := -1
		for pos := 0; pos < len(idsLeft); pos++ {
			if idsLeft[pos] < minLeft {
				minLeft = idsLeft[pos]
				idMinLeft = pos
			}

			if idsRight[pos] < minRight {
				minRight = idsRight[pos]
				idMinRight = pos
			}
		}

		if minLeft > minRight {
			resultPart1 += minLeft - minRight
		} else {
			resultPart1 += minRight - minLeft
		}

		idsLeft = append(idsLeft[:idMinLeft], idsLeft[idMinLeft+1:]...)
		idsRight = append(idsRight[:idMinRight], idsRight[idMinRight+1:]...)
	}

	fmt.Println("Day 1, part 1:", resultPart1)
	fmt.Println("Day 1, part 2:", resultPart2)

	if err := scanner.Err(); err != nil {
		log.Fatalf("Erreur pendant la lecture : %v", err)
	}
}
