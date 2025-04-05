package service

import "siem-sistem/internal/repository"

type Service struct {
	Repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{Repo: repo}
}

func (s *Service) GetAllItems() (string, error) {
	return s.Repo.GetAllItems()
}
