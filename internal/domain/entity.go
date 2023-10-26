package domain

import "errors"

type Reels [3][5]string

type WinLine struct {
	Line      int        `json:"line"`
	Positions []Position `json:"positions"`
}

type Position struct {
	Row int `json:"row"`
	Col int `json:"col"`
}
type WinLines []WinLine

type PayoutSymbol struct {
	Symbol string `json:"symbol"`
	Payout [5]int `json:"payout"`
}

type Payouts []PayoutSymbol

type Line struct {
	X      int `json:"line"`
	Payout int `json:"payout"`
}
type Result struct {
	Lines []Line `json:"lines"`
	Total int    `json:"total"`
}

type GameDate struct {
	Reels    *Reels
	WinLines *WinLines
	Payouts  *Payouts // todo map
}

// todo
func calculateWinLinePayout(line [5]string, payoutMatrix Payouts) int {
	var count int

	for i := 0; i < 4; i++ {
		if line[i] != line[i+1] {
			break
		}
		count++
	}
	for _, payout := range payoutMatrix {
		if payout.Symbol == line[0] {
			// todo
			return payout.Payout[count]
		}
	}

	return 0
}
func (data *GameDate) Validate() error {
	if data.WinLines == nil {
		return errors.New("lines is empty")
	}

	if data.Payouts == nil {
		return errors.New("payouts is empty")
	}

	if data.Reels == nil {
		return errors.New("reels is empty")
	}

	return nil
}
func (data *GameDate) Calculate() Result {
	var result Result
	result.Lines = make([]Line, 0, len(*data.WinLines))
	for _, value := range *data.WinLines {
		// todo
		var r [5]string
		for index, position := range value.Positions {
			r[index] = data.Reels[position.Row][position.Col]
		}
		win := calculateWinLinePayout(r, *data.Payouts)
		result.Total += win

		result.Lines = append(result.Lines, Line{
			X:      value.Line,
			Payout: win,
		})
	}

	return result
}
