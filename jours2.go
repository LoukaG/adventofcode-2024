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

	fichier, err := os.Open("input/jours2.txt")
	if err != nil {
		log.Fatalf("Erreur en ouvrant le fichier : %v", err)
	}
	defer fichier.Close()

	scanner := bufio.NewScanner(fichier)

	var safeReportPart1 uint16 = 0
	var safeReportPart2 uint16 = 0
	for scanner.Scan() {
		levels := strings.Split(scanner.Text(), " ")

		if isSafePart1(levels) {
			safeReportPart1++
		}

		if isSafePart2(levels) {
			safeReportPart2++
		}

	}

	fmt.Println("Day 2, part 1:", safeReportPart1)
	fmt.Println("Day 2, part 2:", safeReportPart2)

}

func isSafePart1(levels []string) bool {
	var isSafe bool = true
	var isAscending bool = true
	for index := 1; index < len(levels); index++ {
		level, _ := strconv.Atoi(levels[index])
		levelPrev, _ := strconv.Atoi(levels[index-1])
		if index == 1 {
			isAscending = level > levelPrev
		}
		diff := level - levelPrev

		isSafe = (isAscending && diff > 0 && diff <= 3) || (!isAscending && diff < 0 && diff >= -3)

		if !isSafe {
			break
		}
	}

	return isSafe
}

func isSafePart2(levels []string) bool {
	var isSafe bool = isSafePart1(levels)
	if !isSafe {
		for index := 0; index < len(levels); index++ {
			cutLevels := make([]string, 0, len(levels)-1)
			cutLevels = append(cutLevels, levels[:index]...)
			cutLevels = append(cutLevels, levels[index+1:]...)
			if isSafePart1(cutLevels) {
				return true
			}
		}
	}
	return isSafe
}
