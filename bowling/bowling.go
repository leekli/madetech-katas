package main

import "fmt"

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
