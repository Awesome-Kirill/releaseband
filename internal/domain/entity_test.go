package domain

import (
	"reflect"
	"testing"
)

func TestGameDate_Calculate(t *testing.T) {
	reels0 := Reels{
		[5]string{"A", "B", "C", "D", "E"},
		[5]string{"F", "A", "F", "B", "C"},
		[5]string{"D", "E", "A", "G", "A"},
	}

	winLines0 := Lines{WinLine{
		Index:     1,
		Positions: []Position{{0, 0}, {1, 1}, {2, 2}, {1, 3}, {0, 4}},
	},

		WinLine{
			Index:     2,
			Positions: []Position{{2, 0}, {1, 1}, {0, 2}, {1, 3}, {2, 4}},
		},

		WinLine{
			Index:     3,
			Positions: []Position{{1, 0}, {2, 1}, {1, 2}, {0, 3}, {1, 4}},
		},
	}

	payouts0 := Payouts{{
		Symbol: "A",
		Payout: [5]int{0, 0, 50, 100, 200},
	},
		{
			Symbol: "B",
			Payout: [5]int{0, 0, 40, 80, 160},
		},
		{
			Symbol: "C",
			Payout: [5]int{0, 0, 30, 60, 120},
		},
		{
			Symbol: "D",
			Payout: [5]int{0, 0, 20, 40, 80},
		},
		{
			Symbol: "E",
			Payout: [5]int{0, 0, 10, 20, 40},
		},
		{
			Symbol: "F",
			Payout: [5]int{0, 0, 5, 10, 20},
		},
		{
			Symbol: "G",
			Payout: [5]int{0, 0, 2, 5, 10},
		},
	}

	result0 := Result{
		Lines: []Line{
			{1, 50},
			{2, 0},
			{3, 0},
		},
		Total: 50,
	}
	type fields struct {
		reels    Reels
		winLines Lines
		payouts  Payouts
	}
	tests := []struct {
		name   string
		fields fields
		want   Result
	}{
		{name: "first", fields: fields{
			reels:    reels0,
			winLines: winLines0,
			payouts:  payouts0,
		}, want: result0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &GameDate{
				Reels:    &tt.fields.reels,
				WinLines: &tt.fields.winLines,
				Payouts:  &tt.fields.payouts,
			}
			// todo err
			if got, _ := data.Calculate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
