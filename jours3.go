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
	fichier, err := os.Open("input/jours3.txt")
	if err != nil {
		log.Fatalf("Erreur en ouvrant le fichier : %v", err)
	}
	defer fichier.Close()

	scanner := bufio.NewScanner(fichier)

	var resultPart1 int = 0
	var resultPart2 int = 0
	var enable bool = true

	for scanner.Scan() {
		var text string = scanner.Text()
		for len(text) > 0 {
			if strings.HasPrefix(text, "mul(") {
				text = text[4:]

				var firstNumber int = getNumber(text)

				//fmt.Println(text)
				text = text[len(strconv.Itoa(firstNumber)):]

				if firstNumber > -1 && text[0] == ',' {
					text = text[1:]
					var secondNumber int = getNumber(text)

					text = text[len(strconv.Itoa(secondNumber)):]

					if text[0] == ')' {
						text = text[1:]

						fmt.Println("mul(", firstNumber, ",", secondNumber, ")=", (firstNumber * secondNumber))
						resultPart1 += firstNumber * secondNumber
						if enable {
							resultPart2 += firstNumber * secondNumber
						}
					}
				}

			} else if strings.HasPrefix(text, "do()") {
				enable = true
				text = text[4:]
			} else if strings.HasPrefix(text, "don't()") {
				enable = false
				text = text[7:]
			} else {
				text = text[1:]
			}
		}

	}

	fmt.Println("Day 3, part 1:", resultPart1)
	fmt.Println("Day 3, part 2:", resultPart2)
}

func getNumber(text string) int {
	var firstNumber string = ""
	for pos := 0; pos < 3; pos++ {
		if len(text) > 0 && text[0] >= '0' && text[0] <= '9' {
			firstNumber += string(text[0])
			text = text[1:]
		} else {
			break
		}
	}

	number, err := strconv.Atoi(firstNumber)

	if err != nil {
		return -1
	} else {
		return number
	}
}
