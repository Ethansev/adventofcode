package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
--- Day 2: Cube Conundrum ---
You're launched high into the atmosphere! The apex of your trajectory just barely reaches the surface of a large island floating in the sky. You gently land in a fluffy pile of leaves. It's quite cold, but you don't see much snow. An Elf runs over to greet you.

The Elf explains that you've arrived at Snow Island and apologizes for the lack of snow. He'll be happy to explain the situation, but it's a bit of a walk, so you have some time. They don't get many visitors up here; would you like to play a game in the meantime?

As you walk, the Elf shows you a small bag and some cubes which are either red, green, or blue. Each time you play this game, he will hide a secret number of cubes of each color in the bag, and your goal is to figure out information about the number of cubes.

To get information, once a bag has been loaded with cubes, the Elf will reach into the bag, grab a handful of random cubes, show them to you, and then put them back in the bag. He'll do this a few times per game.

You play several games and record the information from each game (your puzzle input). Each game is listed with its ID number (like the 11 in Game 11: ...) followed by a semicolon-separated list of subsets of cubes that were revealed from the bag (like 3 red, 5 green, 4 blue).

For example, the record of a few games might look like this:

Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
In game 1, three sets of cubes are revealed from the bag (and then put back again). The first set is 3 blue cubes and 4 red cubes; the second set is 1 red cube, 2 green cubes, and 6 blue cubes; the third set is only 2 green cubes.

The Elf would first like to know which games would have been possible if the bag contained only 12 red cubes, 13 green cubes, and 14 blue cubes?

In the example above, games 1, 2, and 5 would have been possible if the bag had been loaded with that configuration. However, game 3 would have been impossible because at one point the Elf showed you 20 red cubes at once; similarly, game 4 would also have been impossible because the Elf showed you 15 blue cubes at once. If you add up the IDs of the games that would have been possible, you get 8.

Determine which games would have been possible if the bag had been loaded with only 12 red cubes, 13 green cubes, and 14 blue cubes. What is the sum of the IDs of those games?
*/

func main() {
	// condition to sum is 12 red, 13 green, and 14 blue cubes
	games, err := parseInput2()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	sum := 0
	for _, game := range games {
		parts := strings.SplitN(game, ": ", 2)
		gameID, _ := strconv.Atoi(strings.TrimPrefix(parts[0], "Game "))
		reveals := strings.Split(parts[1], "; ")

		maxRed, maxGreen, maxBlue := 0, 0, 0
		for _, reveal := range reveals {
			red, green, blue := countCubes(reveal)
			maxRed = max(maxRed, red)
			maxGreen = max(maxGreen, green)
			maxBlue = max(maxBlue, blue)
		}

		if maxRed <= 12 && maxGreen <= 13 && maxBlue <= 14 {
			sum += gameID
		}
	}

	fmt.Println("Sum of IDs:", sum)

}

func parseInput2() ([]string, error) {
	file, err := os.Open("2_input.txt")
	if err != nil {
		fmt.Println("Could not open the file", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var inputsList []string
	for scanner.Scan() {
		inputsList = append(inputsList, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error while scanning", err)
		return nil, err
	}

	return inputsList, nil
}

func countCubes(reveal string) (int, int, int) {
	redCount, greenCount, blueCount := 0, 0, 0

	parts := strings.Split(reveal, ", ")

	for _, part := range parts {
		// Splitting 16 blue into separate pieces
		piece := strings.Split(part, " ")
		if len(piece) != 2 {
			continue
		}

		count, err := strconv.Atoi(piece[0])
		if err != nil {
			fmt.Println("Oh no something happened...", err)
			continue
		}

		switch piece[1] {
		case "red":
			redCount += count
		case "green":
			greenCount += count
		case "blue":
			blueCount += count
		}
	}

	return redCount, greenCount, blueCount
}

// TODO: add this to a utility function
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
