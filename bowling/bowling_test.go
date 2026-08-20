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
