package domain

import (
	"errors"
	"fmt"
)

type Reels [3][5]string

var Alphabet = map[string]struct{}{
	"A": {}, "B": {}, "C": {},
	"D": {}, "E": {}, "F": {},
	"G": {},
}

type WinLine struct {
	Index     int        `json:"line"`
	Positions []Position `json:"positions"`
}

type Position struct {
	Row int `json:"row"`
	Col int `json:"col"`
}
type Lines []WinLine

type PayoutSymbol struct {
	Symbol string `json:"symbol"`
	Payout [5]int `json:"payout"`
}

type Payouts []PayoutSymbol

type Line struct {
	Index  int `json:"line"`
	Payout int `json:"payout"`
}
type Result struct {
	Lines []Line `json:"lines"`
	Total int    `json:"total"`
}

type GameDate struct {
	Reels    *Reels
	WinLines *Lines
	Payouts  *Payouts // todo map
}

func (r *Reels) Validate() error {
	for indexR, row := range r {
		for indexC, column := range row {
			_, ok := Alphabet[column]
			if !ok {
				return fmt.Errorf("not valid in char in row:%v, column:%v", indexR, indexC)
			}
		}
	}
	return nil
}

// todo
func (g *GameDate) calculateRepeated(line [5]string) int {
	count := 1

	// todo test index out of range
	for i := 0; i < 4; i++ {
		if line[i] != line[i+1] {
			break
		}
		count++
	}

	return count
}
func (g *GameDate) calculateWinLinePayout(line [5]string) int {

	count := g.calculateRepeated(line)
	for _, payout := range *g.Payouts {
		if payout.Symbol == line[0] {
			return payout.Payout[count-1]
		}
	}

	return 0
}
func (g *GameDate) validate() error {
	if g.WinLines == nil {
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

func (g *GameDate) Calculate() (Result, error) {
	var result Result
	err := g.validate()
	if err != nil {
		return result, err
	}

	result.Lines = make([]Line, 0, len(*g.WinLines))
	for _, value := range *g.WinLines {
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
