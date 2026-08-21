package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInputToFrameSlice_ReturnsEmptySliceForEmptyInput(test *testing.T) {
	input := ""
	expected := [][]string{}

	output := ParseInputToFrameSlice(input)

	assert.Equal(test, expected, output)
}

func TestParseInputToFrameSlice_ReturnsSingleSliceForOneCharInput(test *testing.T) {
	input := "X"
	expected := [][]string{{"X"}}

	output := ParseInputToFrameSlice(input)

	assert.Equal(test, expected, output)
}

func TestParseInputToFrameSlice_ReturnsSingleSliceForTwoCharInput(test *testing.T) {
	input := "23"
	expected := [][]string{{"23"}}

	output := ParseInputToFrameSlice(input)

	assert.Equal(test, expected, output)
}

func TestParseInputToFrameSlice_ReturnsMultipleSlicesForVariousCharInput(test *testing.T) {
	input := "23 X"
	expected := [][]string{{"23"}, {"X"}}

	output := ParseInputToFrameSlice(input)

	assert.Equal(test, expected, output)

	input = "3/ 45"
	expected = [][]string{{"3/"}, {"45"}}

	output = ParseInputToFrameSlice(input)

	assert.Equal(test, expected, output)

	input = "X 45 4/ 32"
	expected = [][]string{{"X"}, {"45"}, {"4/"}, {"32"}}

	output = ParseInputToFrameSlice(input)

	assert.Equal(test, expected, output)
}

func TestIsStrike_ReturnsFalseIfFrameHasMoreThanOneElement(test *testing.T) {
	input := []string{"3", "/"}
	expected := false

	output := IsStrike(input)

	assert.Equal(test, expected, output)
}

func TestIsStrike_ReturnsFalseForSingleNonStrikeRoll(test *testing.T) {
	input := []string{"9"}
	expected := false

	output := IsStrike(input)

	assert.Equal(test, expected, output)
}

func TestIsStrike_ReturnsFalseForNotAStrike(test *testing.T) {
	input := []string{"2", "/"}
	expected := false

	output := IsStrike(input)

	assert.Equal(test, expected, output)
}

func TestIsStrike_ReturnsTrueForIsAStrike(test *testing.T) {
	input := []string{"X"}
	expected := true

	output := IsStrike(input)

	assert.Equal(test, expected, output)
}

func TestIsSpare_ReturnsFalseIfFrameHasLessThanTwoElements(test *testing.T) {
	input := []string{"X"}
	expected := false

	output := IsSpare(input)

	assert.Equal(test, expected, output)
}

func TestIsSpare_ReturnsFalseIfFrameHasMoreThanTwoElements(test *testing.T) {
	input := []string{"3", "/", "2"}
	expected := false

	output := IsSpare(input)

	assert.Equal(test, expected, output)
}

func TestIsSpare_ReturnsFalseForNotASpare(test *testing.T) {
	input := []string{"2", "3"}
	expected := false

	output := IsSpare(input)

	assert.Equal(test, expected, output)
}

func TestIsSpare_ReturnsTrueForIsASpare(test *testing.T) {
	input := []string{"3", "/"}
	expected := true

	output := IsSpare(input)

	assert.Equal(test, expected, output)
}

func TestCalculateScoreForNoStrikeOrSpare_ReturnsZeroForNoScore(test *testing.T) {
	inputCurrentFrame := "00"
	expected := 0

	output := CalculateScoreForNoStrikeOrSpare(inputCurrentFrame)

	assert.Equal(test, expected, output)
}

func TestCalculateScoreForNoStrikeOrSpare_ReturnsScore(test *testing.T) {
	inputCurrentFrame := "45"
	expected := 9

	output := CalculateScoreForNoStrikeOrSpare(inputCurrentFrame)

	assert.Equal(test, expected, output)

	inputCurrentFrame = "32"
	expected = 5

	output = CalculateScoreForNoStrikeOrSpare(inputCurrentFrame)

	assert.Equal(test, expected, output)
}

func TestCalculateScoreForStrike_ReturnsScoreForCurrentAndNextFrame(test *testing.T) {
	inputCurrentFrame := "X"
	inputNextFrame := "45"
	expected := 19

	output := CalculateScoreForStrike(inputCurrentFrame, inputNextFrame)

	assert.Equal(test, expected, output)
}

func TestCalculateScoreForStrike_ReturnsZeroForNonStrikeCurrentFrame(test *testing.T) {
	inputCurrentFrame := "45"
	inputNextFrame := "32"
	expected := 0

	output := CalculateScoreForStrike(inputCurrentFrame, inputNextFrame)

	assert.Equal(test, expected, output)
}

func TestCalculateScoreForStrike_IncludesBothRollsWhenNextFrameIsASpare(test *testing.T) {
	inputCurrentFrame := "X"
	inputNextFrame := "4/"
	expected := 20

	output := CalculateScoreForStrike(inputCurrentFrame, inputNextFrame)

	assert.Equal(test, expected, output)
}

func TestCalculateScoreForStrike_IncludesNextStrikeBonus(test *testing.T) {
	inputCurrentFrame := "X"
	inputNextFrame := "X"
	expected := 20

	output := CalculateScoreForStrike(inputCurrentFrame, inputNextFrame)

	assert.Equal(test, expected, output)
}

func TestCalculateScoreForSpare_ReturnsScoreForCurrentAndNextFrameFirstRollOnly(test *testing.T) {
	inputCurrentFrame := "4/"
	inputNextFrame := "32"
	expected := 13

	output := CalculateScoreForSpare(inputCurrentFrame, inputNextFrame)

	assert.Equal(test, expected, output)
}

func TestCalculateScoreForSpare_ReturnsZeroForNonSpareCurrentFrame(test *testing.T) {
	inputCurrentFrame := "45"
	inputNextFrame := "32"
	expected := 0

	output := CalculateScoreForSpare(inputCurrentFrame, inputNextFrame)

	assert.Equal(test, expected, output)
}

func TestCalculateScoreForSpareIncludesStrikeAsNextRoll(test *testing.T) {
	inputCurrentFrame := "4/"
	inputNextFrame := "X"
	expected := 20

	output := CalculateScoreForSpare(inputCurrentFrame, inputNextFrame)

	assert.Equal(test, expected, output)
}

func TestCalculateScoreForTurnFrame_ReturnsZeroForNoScore(test *testing.T) {
	inputCurrentFrame := "00"
	inputNextFrame := ""
	expected := 0

	output := CalculateScoreForTurnFrame(inputCurrentFrame, inputNextFrame)

	assert.Equal(test, expected, output)
}

func TestCalculateScoreForTurnFrame_ReturnsScoreForNonStrikeAndNonSpare_SingleFrame(test *testing.T) {
	inputCurrentFrame := "45"
	inputNextFrame := ""
	expected := 9

	output := CalculateScoreForTurnFrame(inputCurrentFrame, inputNextFrame)

	assert.Equal(test, expected, output)

	inputCurrentFrame = "32"
	inputNextFrame = ""
	expected = 5

	output = CalculateScoreForTurnFrame(inputCurrentFrame, inputNextFrame)

	assert.Equal(test, expected, output)
}

func TestCalculateScoreForTurnFrame_ReturnsScoreForStrikeAndNextTurn(test *testing.T) {
	inputCurrentFrame := "X"
	inputNextFrame := "45"
	expected := 19

	output := CalculateScoreForTurnFrame(inputCurrentFrame, inputNextFrame)

	assert.Equal(test, expected, output)
}

func TestCalculateScoreForTurnFrame_ReturnsScoreForSpareAndNextTurn(test *testing.T) {
	inputCurrentFrame := "4/"
	inputNextFrame := "32"
	expected := 13

	output := CalculateScoreForTurnFrame(inputCurrentFrame, inputNextFrame)

	assert.Equal(test, expected, output)
}

func TestCalculateGameScoreTotal_ReturnsTotal(test *testing.T) {
	input := "X 45 4/ 32 00 00 00 00 00 00"
	expected := 46

	output := CalculateGameScoreTotal(input)

	assert.Equal(test, expected, output)
}

func TestCalculateGameScoreTotal_ReturnsZeroForAllGutters(test *testing.T) {
	input := "00 00 00 00 00 00 00 00 00 00"
	expected := 0

	output := CalculateGameScoreTotal(input)

	assert.Equal(test, expected, output)
}

func TestCalculateGameScoreTotal_ReturnsTwentyForAllOnes(test *testing.T) {
	input := "11 11 11 11 11 11 11 11 11 11"
	expected := 20

	output := CalculateGameScoreTotal(input)

	assert.Equal(test, expected, output)
}

func TestCalculateGameScoreTotal_IncludesSpareBonusFromStrike(test *testing.T) {
	input := "X 4/ 32 00 00 00 00 00 00 00"
	expected := 38

	output := CalculateGameScoreTotal(input)

	assert.Equal(test, expected, output)

	input = "4/ X 32 00 00 00 00 00 00 00"
	expected = 40

	output = CalculateGameScoreTotal(input)

	assert.Equal(test, expected, output)

	input = "X 45 32 00 00 00 00 00 00 00"
	expected = 33

	output = CalculateGameScoreTotal(input)

	assert.Equal(test, expected, output)

	input = "4/ 32 11 00 00 00 00 00 00 00"
	expected = 20

	output = CalculateGameScoreTotal(input)

	assert.Equal(test, expected, output)
}

func TestCalculateGameScoreTotal_IncludesOneBonusBallAfterTenthFrameSpare(test *testing.T) {
	input := "00 00 00 00 00 00 00 00 00 4/ 5"
	expected := 15

	output := CalculateGameScoreTotal(input)

	assert.Equal(test, expected, output)
}

func TestCalculateGameScoreTotal_IncludesStrikeBonusAfterTenthFrameSpare(test *testing.T) {
	input := "00 00 00 00 00 00 00 00 00 4/ X"
	expected := 20

	output := CalculateGameScoreTotal(input)

	assert.Equal(test, expected, output)
}

func TestCalculateGameScoreTotal_IncludesTwoBonusBallsAfterTenthFrameStrike(test *testing.T) {
	input := "00 00 00 00 00 00 00 00 00 X 5 4"
	expected := 19

	output := CalculateGameScoreTotal(input)

	assert.Equal(test, expected, output)
}

func TestCalculateGameScoreTotal_IncludesTwoStrikeBonusBallsAfterTenthFrameStrike(test *testing.T) {
	input := "00 00 00 00 00 00 00 00 00 X X X"
	expected := 30

	output := CalculateGameScoreTotal(input)

	assert.Equal(test, expected, output)
}

func TestCalculateGameScoreTotal_IncludesNumericSecondBonusBallAfterTenthFrameStrike(test *testing.T) {
	input := "00 00 00 00 00 00 00 00 00 X X 5"
	expected := 25

	output := CalculateGameScoreTotal(input)

	assert.Equal(test, expected, output)
}

func TestCalculateGameScoreTotal_IncludesNumericFirstBonusBallAfterTenthFrameStrike(test *testing.T) {
	input := "00 00 00 00 00 00 00 00 00 X 5 X"
	expected := 25

	output := CalculateGameScoreTotal(input)

	assert.Equal(test, expected, output)
}

func TestCalculateGameScoreTotal_ScoresOpenTenthFrame(test *testing.T) {
	input := "00 00 00 00 00 00 00 00 00 45"
	expected := 9

	output := CalculateGameScoreTotal(input)

	assert.Equal(test, expected, output)
}
