package domain

import "fmt"

const MaxColumn = 5
const MaxRows = 2

type Lines []WinLine

type WinLine struct {
	Index     int        `json:"line"`
	Positions []Position `json:"positions"`
}

type Position struct {
	Row int `json:"row"`
	Col int `json:"col"`
}

// Validate Lines
func (w Lines) Validate() error {
	for _, line := range w {
		for index, value := range line.Positions {
			if value.Row > MaxRows {
				return fmt.Errorf("row:%v", index)
			}

			if value.Col > MaxColumn {
				return fmt.Errorf("col:%v", index)
			}
		}
	}

	return nil
}
