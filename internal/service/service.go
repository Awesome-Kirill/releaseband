package service

import (
	"errors"
	"fmt"
	"releaseband/internal/domain"
)

type Repository interface {
	GetGame(id string) (*domain.GameDate, bool)
	SetLines(id string, input *domain.Lines)
	SetPayouts(id string, pay *domain.Payouts)
	SetReels(id string, reels *domain.Reels)
}

type Service struct {
	repository Repository
}

func New(repository Repository) *Service {
	return &Service{repository: repository}
}

// CreateLines create lines
func (s *Service) CreateLines(id string, lines domain.Lines) error {
	err := lines.Validate()
	if err != nil {
		return fmt.Errorf("error create lines:%v|%w", id, err)
	}
	s.repository.SetLines(id, &lines)
	return nil
}

// CreateReels create reels
func (s *Service) CreateReels(id string, reels *domain.Reels) error {
	err := reels.Validate()
	if err != nil {
		return fmt.Errorf("error reels lines:%v|%w", id, err)
	}
	s.repository.SetReels(id, reels)
	return nil
}

// CreatePayouts create payouts
func (s *Service) CreatePayouts(id string, payouts domain.Payouts) error {
	err := payouts.Validate()
	if err != nil {
		return fmt.Errorf("error payouts lines:%v|%w", id, err)
	}
	s.repository.SetPayouts(id, &payouts)
	return nil
}

// GetCalculateDate get win date
func (s *Service) GetCalculateDate(id string) (domain.Result, error) {
	game, ok := s.repository.GetGame(id)
	if !ok {
		return domain.Result{}, errors.New("not found")
	}

	return game.Calculate()
}
