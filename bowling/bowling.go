package main

import (
	"fmt"
	"strconv"
	"strings"
)

/* Game:
- A game has 10 'turns' or 'frames'
- Bowling has 10 pins to knock down
- Each turn/try is 2 rolls of the ball each
- X means 'strike' - all 10 knocked down in one go
- / means 'spare' - all 10 knocked down across two go's
*/

/* Scoring:
- If in 2 tries fails knock all 10 down: total is +1 for each they did knock down
- If in 2 tries, all 10 pins knocked down: this is 'spare': score is 10 + number of pins knocked down in next throw in their first throw only
- If in 1 try, all 10 pins knocked down: this is 'strike': score is 10 + number of pins knocked down innext throw over both tries
- If bowler gets spare or stike in last game (10th one), then bowler gets 1 or 2 more bonus balls (1 for spare, 2 for stike): Total pins hit taken from this to add to [last game - 1] game frame total.
- Game score is total of all frame scores.
*/

func main() {
	fmt.Println("Hello!")
}

func ParseInputToFrameSlice(input string) [][]string {
	if len(input) == 0 {
		return [][]string{}
	}

	delimString := strings.Split(input, " ")

	frameSlice := [][]string{}

	for i := 0; i < len(delimString); i++ {
		singleFrame := []string{delimString[i]}

		frameSlice = append(frameSlice, singleFrame)
	}

	return frameSlice
}

func IsStrike(frame []string) bool {
	if len(frame) > 1 {
		return false
	}

	if frame[0] != "X" {
		return false
	}

	if frame[0] == "X" {
		return true
	}

	return false
}

func IsSpare(frame []string) bool {
	if len(frame) <= 1 {
		return false
	}

	if len(frame) > 2 {
		return false
	}

	if frame[1] == "/" {
		return true
	}

	return false
}

func CalculateGameScoreTotal(gameInput string) int {
	gameTotal := 0

	splitGameScores := strings.Split(gameInput, " ")

	// Get the scores of every game up to but not including the final turn
	for i := 0; i < (len(splitGameScores) - 1); i++ {
		frameScore := CalculateScoreForTurnFrame(splitGameScores[i], splitGameScores[i+1])

		gameTotal += frameScore
	}

	// Now get the score of the final turn
	frameScore := CalculateScoreForTurnFrame(splitGameScores[len(splitGameScores)-1], "")

	gameTotal += frameScore

	return gameTotal
}

func CalculateScoreForTurnFrame(currentFrame string, nextFrame string) int {
	frameTotal := 0

	// Split out current frame into individual chars
	currentFrameScoresSplit := strings.Split(currentFrame, "")

	// If the frame is not a Strike (X) or Spare (/), get the numeric total of what was knocked down
	if !IsStrike(currentFrameScoresSplit) && !IsSpare(currentFrameScoresSplit) {
		currentFrameTotal := CalculateScoreForNoStrikeOrSpare(currentFrame)

		frameTotal += currentFrameTotal
	}

	if IsStrike(currentFrameScoresSplit) {
		currentFrameTotal := CalculateScoreForStrike(currentFrame, nextFrame)

		frameTotal += currentFrameTotal
	}

	if IsSpare(currentFrameScoresSplit) {
		currentFrameTotal := CalculateScoreForSpare(currentFrame, nextFrame)

		frameTotal += currentFrameTotal
	}

	return frameTotal
}

func CalculateScoreForNoStrikeOrSpare(currentFrame string) int {
	frameTotal := 0

	// Scores come in as a string pair, e.g. "13". Split them out and convert to numbers where possible
	currentFrameScoresSplit := strings.Split(currentFrame, "")

	// If the frame does not include a Strike (X) or Spare (/) then calculate the single current frames total score and add to the total
	currentFrameIntScores := convertNumericStringsToIntArray(currentFrameScoresSplit)

	// Calculate the current frame total for int scores only
	for _, num := range currentFrameIntScores {
		frameTotal += num
	}

	return frameTotal
}

func CalculateScoreForStrike(currentFrame string, nextFrame string) int {
	frameTotal := 0

	if currentFrame != "X" {
		return 0
	}

	// Assign a score of 10 for the initial strike
	if currentFrame == "X" {
		frameTotal += 10
	}

	// Then find the total of the next turn (nextFrame)
	// Split out current frame into individual chars
	nextFrameScoresSplit := strings.Split(nextFrame, "")

	// If the frame is not a Strike (X) or Spare (/), get the numeric total of what was knocked down
	if !IsStrike(nextFrameScoresSplit) && !IsSpare(nextFrameScoresSplit) {
		nextFrameTotal := CalculateScoreForNoStrikeOrSpare(nextFrame)

		frameTotal += nextFrameTotal
	}

	return frameTotal
}

func CalculateScoreForSpare(currentFrame string, nextFrame string) int {
	frameTotal := 0

	if !strings.Contains(currentFrame, "/") {
		return 0
	}

	// Assign a score of 10 for the initial strike
	if strings.Contains(currentFrame, "/") {
		frameTotal += 10
	}

	// Then find the total of the next turn (nextFrame)
	// Split out current frame into individual chars
	nextFrameScoresSplit := strings.Split(nextFrame, "")

	// If the frame is not a Strike (X) or Spare (/), get the numeric total of what was knocked down
	if !IsStrike(nextFrameScoresSplit) && !IsSpare(nextFrameScoresSplit) {
		// Get only the first roll of the 2nd frame/turn
		firstRollOnly := nextFrameScoresSplit[0]

		nextFrameTotal := CalculateScoreForNoStrikeOrSpare(firstRollOnly)

		frameTotal += nextFrameTotal
	}

	return frameTotal
}

func convertNumericStringsToIntArray(frame []string) []int {
	frameIntScores := []int{}

	for _, value := range frame {
		num, err := strconv.Atoi(value)

		if err != nil {
			fmt.Println("Error: ", err)
		}

		frameIntScores = append(frameIntScores, num)
	}

	return frameIntScores
}
