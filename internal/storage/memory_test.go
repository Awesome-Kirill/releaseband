package storage

import (
	"reflect"
	"releaseband/internal/domain"
	"testing"
)

func TestInMemory_GetGame(t *testing.T) {
	id := "test_game_key"
	m := New()
	m.SetPayouts(id, &domain.Payouts{})

	tests := []struct {
		name  string
		id    string
		want  *domain.GameDate
		want1 bool
	}{
		{name: "exist", id: id, want: &domain.GameDate{
			Payouts: &domain.Payouts{},
		}, want1: true},
		{name: "not-exist", id: "not-exist", want: nil, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := m.GetGame(tt.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGame() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetGame() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestInMemory_SetReels(t *testing.T) {
	tests := []struct {
		name  string
		id    string
		reels domain.Reels
		want  *domain.GameDate
	}{
		{name: "test", id: "not-empty", reels: domain.Reels{}, want: &domain.GameDate{Reels: &domain.Reels{}}},
	}

	m := New()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.SetReels(tt.id, &tt.reels)

			got, _ := m.GetGame(tt.id)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGame() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInMemory_SetLines(t *testing.T) {
	tests := []struct {
		name  string
		id    string
		lines domain.Lines
		want  *domain.GameDate
	}{
		{name: "test", id: "not-empty", lines: domain.Lines{}, want: &domain.GameDate{Lines: &domain.Lines{}}},
	}

	m := New()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.SetLines(tt.id, &tt.lines)

			got, _ := m.GetGame(tt.id)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGame() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInMemory_SetPayouts(t *testing.T) {
	tests := []struct {
		name string
		id   string
		pay  domain.Payouts
		want *domain.GameDate
	}{
		{name: "test", id: "not-empty", pay: domain.Payouts{}, want: &domain.GameDate{Payouts: &domain.Payouts{}}},
	}

	m := New()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.SetPayouts(tt.id, &tt.pay)

			got, _ := m.GetGame(tt.id)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGame() got = %v, want %v", got, tt.want)
			}
		})
	}
}
