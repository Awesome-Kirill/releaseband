package domain

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
	reels    Reels
	winLines WinLines
	payouts  Payouts // todo map
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
func (data *GameDate) Calculate() Result {
	var result Result
	for _, value := range data.winLines {
		var r [5]string
		for index, position := range value.Positions {
			r[index] = data.reels[position.Row][position.Col]
		}
		win := calculateWinLinePayout(r, data.payouts)
		result.Total += win

		result.Lines = append(result.Lines, Line{
			X:      value.Line,
			Payout: win,
		})
	}

	return result
}
