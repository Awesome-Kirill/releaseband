package domain

import "fmt"

type Lines []WinLine

type WinLine struct {
	Index     int        `json:"line"`
	Positions []Position `json:"positions"`
}

func (w Lines) Validate() error {

	for _, line := range w {
		for index, value := range line.Positions {
			if value.Row > 4 {
				return fmt.Errorf("row:%v", index)
			}

			if value.Col > 5 {
				return fmt.Errorf("col:%v", index)
			}
		}
	}

	return nil
}
