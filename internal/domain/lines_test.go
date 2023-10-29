package domain

import "testing"

func TestLines_Validate(t *testing.T) {
	// todo
	tests := []Lines{{{Index: 1, Positions: []Position{{Row: 17}}}}}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if err := tt.Validate(); err == nil {
				t.Errorf("Validate() error")
			}
		})
	}
}
