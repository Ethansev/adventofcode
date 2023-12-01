package main

// Go won't auto format the comments, I'm sorry :'(
/*
https://adventofcode.com/2023/day/1
--- Day 1: Trebuchet?! ---
Something is wrong with global snow production, and you've been selected to take a look. The Elves have even given you a map; on it, they've used stars to mark the top fifty locations that are likely to be having problems.

You've been doing this long enough to know that to restore snow operations, you need to check all fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

You try to ask why they can't just use a weather machine ("not powerful enough") and where they're even sending you ("the sky") and why your map looks mostly blank ("you sure ask a lot of questions") and hang on did you just say the sky ("of course, where do you think snow comes from") when you realize that the Elves are already loading you into a trebuchet ("please hold still, we need to strap you in").

As they're making the final adjustments, they discover that their calibration document (your puzzle input) has been amended by a very young Elf who was apparently just excited to show off her art skills. Consequently, the Elves are having trouble reading the values on the document.

The newly-improved calibration document consists of lines of text; each line originally contained a specific calibration value that the Elves now need to recover. On each line, the calibration value can be found by combining the first digit and the last digit (in that order) to form a single two-digit number.

For example:

1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
In this example, the calibration values of these four lines are 12, 38, 15, and 77. Adding these together produces 142.

Consider your entire calibration document. What is the sum of all of the calibration values?
*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	// I need to use two pointers on either end until they both find their first number.
	// TODO: does not handle the case where a number does not exist in the string
	// I GOT A GOLD STAR WOOOOOOOOOOOO

	inputs, err := parseInput()
	if err != nil {
		fmt.Println("We goofed up somewhere", err)
		return
	}

	answer := calculate(inputs)

	fmt.Println(answer)
}

func parseInput() ([]string, error) {
	file, err := os.Open("1_input.txt")
	if err != nil {
		fmt.Println("Could not open the file", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	inputsList := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		inputsList = append(inputsList, line)
	}
	return inputsList, nil
}

func calculate(example []string) int {
	sum := 0
	for i := 0; i < len(example); i++ {
		l := 0
		r := len(example[i]) - 1
		var left string
		var right string

		for l < r {
			for left == "" {
				if unicode.IsDigit(rune(example[i][l])) {
					left = string(example[i][l])
				}
				l++
			}
			for right == "" {
				if unicode.IsDigit(rune(example[i][r])) {
					right = string(example[i][r])
				}
				r--
			}
			leftInt, _ := strconv.Atoi(left)
			rightInt, _ := strconv.Atoi(right)
			sum += leftInt*10 + rightInt
			break
		}
	}
	return sum
}
