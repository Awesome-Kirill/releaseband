package domain

import (
	"errors"
)

type Position struct {
	Row int `json:"row"`
	Col int `json:"col"`
}

type Line struct {
	Index  int `json:"line"`
	Payout int `json:"payout"`
}
type Result struct {
	Lines []Line `json:"lines"`
	Total int    `json:"total"`
}

type GameDate struct {
	Reels   *Reels
	Lines   *Lines
	Payouts *Payouts // todo map
}

// count of repeat char(first) in array char
func (g *GameDate) getRepeatedCount(line [5]string) int {
	count := 1
	for i := 0; i < 4; i++ {
		if line[i] != line[i+1] {
			break
		}
		count++
	}

	return count
}

// calculate one line win
func (g *GameDate) calculateWinLinePayout(line [5]string) int {

	count := g.getRepeatedCount(line)
	for _, payout := range *g.Payouts {
		if payout.Symbol == line[0] {
			return payout.Payout[count-1]
		}
	}

	return 0
}

// check all three struct(reels lines payouts) is not empty
func (g *GameDate) validate() error {
	if g.Lines == nil {
		return errors.New("lines is empty")
	}

	if g.Payouts == nil {
		return errors.New("payouts is empty")
	}

	if g.Reels == nil {
		return errors.New("reels is empty")
	}

	return nil
}

// Calculate game
func (g *GameDate) Calculate() (Result, error) {
	var result Result
	err := g.validate()
	if err != nil {
		return result, err
	}

	result.Lines = make([]Line, 0, len(*g.Lines))
	for _, value := range *g.Lines {
		var row [5]string
		for index, position := range value.Positions {
			row[index] = g.Reels[position.Row][position.Col]
		}
		win := g.calculateWinLinePayout(row)
		result.Total += win

		result.Lines = append(result.Lines, Line{
			Index:  value.Index,
			Payout: win,
		})
	}

	return result, nil
}
