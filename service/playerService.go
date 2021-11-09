package service

import (
	"g/go/allsports/domain"
	"g/go/allsports/errs"
)

type PlayerService interface {
	GetAllPlayers() ([]domain.Player, *errs.AppError)
	GetPlayer(string) (*domain.Player, *errs.AppError)
}

type DefaultPlayerService struct {
	repo domain.PlayerRepository
}

func (s DefaultPlayerService) GetPlayer(id string) (*domain.Player, *errs.AppError) {

	return s.repo.ById(id)
}

func (s DefaultPlayerService) GetAllPlayers() ([]domain.Player, *errs.AppError) {

	return s.repo.FindAll()
}

func NewPlayerService(repositery domain.PlayerRepository) DefaultPlayerService {
	return DefaultPlayerService{repo: repositery}
}
