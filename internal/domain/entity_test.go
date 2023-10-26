package domain

import (
	"reflect"
	"testing"
)

func TestGameDate_Calculate(t *testing.T) {
	/*
		[
		  ["A", "B", "C", "D", "E"],
		  ["F", "A", "F", "B", "C"],
		  ["D", "E", "A", "G", "A"]
		]
	*/

	reels0 := Reels{
		[5]string{"A", "B", "C", "D", "E"},
		[5]string{"F", "A", "F", "B", "C"},
		[5]string{"D", "E", "A", "G", "A"},
	}

	/*
			[
		    {
		        "line": 1,
		        "positions": [
		            {"row": 0, "col": 0},
		            {"row": 1, "col": 1},
		            {"row": 2, "col": 2},
		            {"row": 1, "col": 3},
		            {"row": 0, "col": 4}
		        ]
		    },
		    {
		        "line": 2,
		        "positions": [
		            {"row": 2, "col": 0},
		            {"row": 1, "col": 1},
		            {"row": 0, "col": 2},
		            {"row": 1, "col": 3},
		            {"row": 2, "col": 4}
		        ]
		    },
		    {
		        "line": 3,
		        "positions": [
		            {"row": 1, "col": 0},
		            {"row": 2, "col": 1},
		            {"row": 1, "col": 2},
		            {"row": 0, "col": 3},
		            {"row": 1, "col": 4}
		        ]
		    }
		]
	*/

	winLines0 := WinLines{WinLine{
		Line:      1,
		Positions: []Position{{0, 0}, {1, 1}, {2, 2}, {1, 3}, {0, 4}},
	},

		WinLine{
			Line:      2,
			Positions: []Position{{2, 0}, {1, 1}, {0, 2}, {1, 3}, {2, 4}},
		},

		WinLine{
			Line:      3,
			Positions: []Position{{1, 0}, {2, 1}, {1, 2}, {0, 3}, {1, 4}},
		},
	}

	/*
		[
		    {
		        "symbol": "A",
		        "payout": [0, 0, 50, 100, 200]
		    },
		    {
		        "symbol": "B",
		        "payout": [0, 0, 40, 80, 160]
		    },
		    {
		        "symbol": "C",
		        "payout": [0, 0, 30, 60, 120]
		    },
		    {
		        "symbol": "D",
		        "payout": [0, 0, 20, 40, 80]
		    },
		    {
		        "symbol": "E",
		        "payout": [0, 0, 10, 20, 40]
		    },
		    {
		        "symbol": "F",
		        "payout": [0, 0, 5, 10, 20]
		    },
		    {
		        "symbol": "G",
		        "payout": [0, 0, 2, 5, 10]
		    }
		]
	*/
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

	/*
		{
		    "lines": [
		        {"line": 1, "payout": 50},
		        {"line": 2, "payout": 0},
		        {"line": 3, "payout": 0}
		    ],
		    "total": 50
		}
	*/

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
		winLines WinLines
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
			if got := data.Calculate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateWinLinePayout(t *testing.T) {
	type args struct {
		line         [5]string
		payoutMatrix Payouts
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateWinLinePayout(tt.args.line, tt.args.payoutMatrix); got != tt.want {
				t.Errorf("calculateWinLinePayout() = %v, want %v", got, tt.want)
			}
		})
	}
}
