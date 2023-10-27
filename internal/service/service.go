package service

import (
	"errors"
	"releaseband/internal/domain"
)

type Repository interface {
	Get(id string) (*domain.GameDate, bool)
	Set(id string, input *domain.GameDate)
}

type Service struct {
	repository Repository
}

func New(repository Repository) *Service {
	return &Service{repository: repository}
}
func (s *Service) CreateOrUpdate(id string, input *domain.GameDate) error {
	s.repository.Set(id, input)
	return nil
}

func (s *Service) GetCalculateDate(id string) (domain.Result, error) {
	game, ok := s.repository.Get(id)
	if !ok {
		return domain.Result{}, errors.New("not found")
	}

	return game.Calculate()
}
