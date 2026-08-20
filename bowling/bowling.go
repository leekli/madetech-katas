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

func CalculateScoreForTurnFrame(currentFrame string, nextFrame string) int {
	frameTotal := 0

	// If the frame is not a strike or spare, get the numeric total of what was knocked down
	if len(currentFrame) >= 1 && string(currentFrame[0]) != "X" && len(currentFrame) >= 2 && string(currentFrame[1]) != "/" {
		currentFrameTotal := CalculateScoreForNoStrikeOrSpare(currentFrame)

		frameTotal += currentFrameTotal
	}

	return frameTotal
}

func CalculateScoreForNoStrikeOrSpare(currentFrame string) int {
	frameTotal := 0

	// Scores come in as a string pair, e.g. "13". Split them out and convert to numbers where possible
	currentFrameScoresSplit := strings.Split(currentFrame, "")
	currentFrameIntScores := convertNumericStringsToIntArray(currentFrameScoresSplit)

	// Calculate the current frame total
	for _, num := range currentFrameIntScores {
		frameTotal += num
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
