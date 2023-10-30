package storage

import (
	"releaseband/internal/domain"
	"sync"
)

type InMemory struct {
	mu   *sync.RWMutex
	data map[string]*domain.GameDate
}

func New() *InMemory {
	return &InMemory{
		mu:   &sync.RWMutex{},
		data: make(map[string]*domain.GameDate),
	}
}

// GetGame by id all three game date
func (m *InMemory) GetGame(id string) (*domain.GameDate, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	v, ok := m.data[id]
	return v, ok
}

// SetPayouts Set Payouts
func (m *InMemory) SetPayouts(id string, payouts domain.Payouts) {
	m.mu.Lock()
	defer m.mu.Unlock()
	v, ok := m.data[id]
	var game domain.GameDate
	if ok {
		game = *v
	}

	game.Payouts = &payouts

	m.data[id] = &game
}

// SetLines set lines
func (m *InMemory) SetLines(id string, lines domain.Lines) {
	m.mu.Lock()
	defer m.mu.Unlock()
	v, ok := m.data[id]
	var game domain.GameDate
	if ok {
		game = *v
	}

	game.Lines = &lines

	m.data[id] = &game
}

// SetReels set reels
func (m *InMemory) SetReels(id string, reels domain.Reels) {
	m.mu.Lock()
	defer m.mu.Unlock()
	v, ok := m.data[id]
	var game domain.GameDate
	if ok {
		game = *v
	}

	game.Reels = &reels

	m.data[id] = &game
}
