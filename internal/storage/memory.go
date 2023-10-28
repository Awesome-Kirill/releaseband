package storage

import (
	"releaseband/internal/domain"
	"sync"
)

type InMemory struct {
	mu   sync.RWMutex
	data map[string]*domain.GameDate
}

func New() *InMemory {
	return &InMemory{
		mu:   sync.RWMutex{},
		data: make(map[string]*domain.GameDate),
	}
}

// Get by id all three game date
func (m *InMemory) Get(id string) (*domain.GameDate, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	v, ok := m.data[id]
	return v, ok
}

// Set (create or update) all three date
func (m *InMemory) Set(id string, input *domain.GameDate) {
	m.mu.Lock()
	defer m.mu.Unlock()
	v, ok := m.data[id]
	var game domain.GameDate
	if ok {
		game = *v
	}

	if input.WinLines != nil {
		game.WinLines = input.WinLines
	}

	if input.Payouts != nil {
		game.Payouts = input.Payouts
	}

	if input.Reels != nil {
		game.Reels = input.Reels
	}

	m.data[id] = &game
}
