package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
