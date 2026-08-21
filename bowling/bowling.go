package main

import (
	"fmt"
	"strconv"
	"strings"
)

/* Bowling Game Kata:
- Kata link: https://learn.madetech.com/technology/katas/bowling/

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
	input := "X 45 4/ 32 4/ 4/ X 32 44 12"

	score := CalculateGameScoreTotal(input)

	fmt.Println("==========================")
	fmt.Println("The provided game is:", input)
	fmt.Println("The total score for this game is:", score)
	fmt.Println("==========================")
}

func ParseInputToFrameSlice(input string) [][]string {
	if len(input) == 0 {
		return [][]string{}
	}

	delimString := strings.Split(input, " ")

	frameSlice := [][]string{}

	for _, str := range delimString {
		singleFrame := []string{str}

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

	// Every game is 10 tries: Get the scores of every game up to but not including the final 10th turn
	for i := range 9 {
		frameScore := CalculateScoreForTurnFrame(splitGameScores[i], splitGameScores[i+1])

		gameTotal += frameScore
	}

	// Now deal with the 10th turn
	tenthTryResult := strings.Split(splitGameScores[9], "")

	if IsStrike(tenthTryResult) {
		// Strike: Add 10 points for hitting a strike on 10th turn
		gameTotal += 10

		// Strike: Deal with bonus ball 1 out of 2
		firstBonusThrowResult := splitGameScores[10]

		if firstBonusThrowResult == "X" {
			gameTotal += 10
		} else {
			gameTotal += CalculateScoreForTurnFrame(firstBonusThrowResult, "")
		}

		// Strike: Deal with bonus ball 2 out of 2
		secondBonusThrowResult := splitGameScores[11]

		if secondBonusThrowResult == "X" {
			gameTotal += 10
		} else {
			gameTotal += CalculateScoreForTurnFrame(secondBonusThrowResult, "")
		}
	} else if IsSpare(tenthTryResult) {
		// Spare: Add 10 points for hitting a spare on 10th turn
		gameTotal += 10

		// Spare: Deal with bonus ball 1 out of 1
		firstBonusThrowResult := splitGameScores[10]

		if firstBonusThrowResult == "X" {
			gameTotal += 10
		} else {
			gameTotal += CalculateScoreForTurnFrame(firstBonusThrowResult, "")
		}
	} else {
		// If 10th turn is not a strike or spare, get the result
		gameTotal += CalculateScoreForTurnFrame(splitGameScores[9], "")
	}

	return gameTotal
}

func CalculateScoreForTurnFrame(currentFrame string, nextFrame string) int {
	frameTotal := 0

	// Split out current frame into individual chars
	currentFrameScoresSplit := strings.Split(currentFrame, "")

	// If the frame is not a Strike (X) or Spare (/), get the numeric total of what was knocked down
	if !IsStrike(currentFrameScoresSplit) && !IsSpare(currentFrameScoresSplit) {
		frameTotal += CalculateScoreForNoStrikeOrSpare(currentFrame)
	}

	if IsStrike(currentFrameScoresSplit) {
		if nextFrame == "" {
			frameTotal += CalculateScoreForStrike(currentFrame, "")
		}

		frameTotal += CalculateScoreForStrike(currentFrame, nextFrame)
	}

	if IsSpare(currentFrameScoresSplit) {
		if nextFrame == "" {
			frameTotal += CalculateScoreForSpare(currentFrame, "")
		}

		frameTotal += CalculateScoreForSpare(currentFrame, nextFrame)
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
		frameTotal += CalculateScoreForNoStrikeOrSpare(nextFrame)
	}

	if IsStrike(nextFrameScoresSplit) || IsSpare(nextFrameScoresSplit) {
		frameTotal += 10
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

		frameTotal += CalculateScoreForNoStrikeOrSpare(firstRollOnly)
	}

	if IsStrike(nextFrameScoresSplit) || IsSpare(nextFrameScoresSplit) {
		frameTotal += 10
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
