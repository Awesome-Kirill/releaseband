package domain

import "fmt"

type Reels [3][5]string

var alphabet = map[string]struct{}{
	"A": {}, "B": {}, "C": {},
	"D": {}, "E": {}, "F": {},
	"G": {},
}

// Validate input reels. Available only A-G symbol
func (r *Reels) Validate() error {
	for indexR, row := range r {
		for indexC, column := range row {
			_, ok := alphabet[column]
			if !ok {
				return fmt.Errorf("not valid in char in row:%v, column:%v", indexR, indexC)
			}
		}
	}
	return nil
}
