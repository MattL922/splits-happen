package main

import (
    "flag"
    "fmt"
    "strconv"
)

const (
    Strike = "X"
    Spare = "/"
    Miss = "-"
)

var linevar string

func main() {
    flag.StringVar(&linevar, "l", "--------------------", "A valid bowling line: '-' for a miss, '/' for a spare, 'X' for a strike, and 1-9 for their respective values")
    flag.Parse()
    line := NewLine(linevar)
    fmt.Printf("Score: %d\n", line.Score())
}

// Roll represents the score of a single roll of the ball down the lane. It acts
// as a string alias with some added functionality.
type Roll string

// IsStrike returns true if the Roll is a strike, false otherwise.
func (r Roll) IsStrike() bool {
    return r == Strike
}

// IsSpare returns true if the Roll is a spare, false otherwise.
func (r Roll) IsSpare() bool {
    return r == Spare
}

// IsMiss returns true if the Roll is a miss, false otherwise.
func (r Roll) IsMiss() bool {
    return r == Miss
}

// Value returns the numeric value of the Roll score.
func (r Roll) Value() int {
    if r.IsStrike() || r.IsSpare() {
        return 10
    } else if r.IsMiss() {
        return 0
    } else {
        // Inputs are guaranteed to be valid so we ignore the error here!
        val, _ := strconv.Atoi(string(r))
        return val
    }
}

// Line holds the rolls taken during 10 frames of a game plus any bonus rolls.
type Line struct {
    Rolls []Roll
}

// NewLine creates and returns a new Line from the given string by splitting it
// into a new Roll for each character.
func NewLine(line string) Line {
    l := Line{}
    for _, roll := range line {
        l.Rolls = append(l.Rolls, Roll(roll))
    }
    return l
}

// Score calculates the total score for the Line.
func (l Line) Score() int {
    score := 0
    // Used to determine how many frames have passed - strikes count as 2 rolls.
    // This makes every game an even 20 rolls, making frames deterministic.
    rollcount := 0
    for i, roll := range l.Rolls {
        score += roll.Value()
        rollcount++
        if roll.IsStrike() { // Strike bonus points
            score += (l.Rolls[i+1].Value() + l.Rolls[i+2].Value())
            // Spare values include the first roll in the frame, so subtract it
            if l.Rolls[i+2].IsSpare() {
                score -= l.Rolls[i+1].Value()
            }
            rollcount++ // extra roll for a strike
        } else if roll.IsSpare() { // Spare bonus points
            // Spare values include the first roll in the frame, so subtract it
            score += (-l.Rolls[i-1].Value() + l.Rolls[i+1].Value())
        }
        // Stop calculating at the end of the 10th frame (don't double count
        // bonus rolls)
        if rollcount == 20 { break }
    }
    return score
}
