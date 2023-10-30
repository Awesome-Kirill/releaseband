package domain

import "testing"

func TestReels_Validate(t *testing.T) {
	tests := []struct {
		r       Reels
		wantErr bool
	}{{
		wantErr: false,
		r:       Reels{{"A", "B", "C", "D", "E"}, {"F", "A", "F", "B", "C"}, {"D", "E", "A", "G", "A"}},
	},
		{
			wantErr: true,
			r:       Reels{{"x", "B", "C", "D", "E"}, {"F", "A", "F", "B", "C"}, {"D", "E", "A", "G", "A"}},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if err := tt.r.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
