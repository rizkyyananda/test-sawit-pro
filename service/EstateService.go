package service

import (
	"context"
	"errors"
	"net/http"
	"test_sawit_pro/entity"
	"test_sawit_pro/repository"
)

type EstateService interface {
	CreateEstate(ctx context.Context, estate *entity.Estate) (string, error, int)
	GetEstateByID(ctx context.Context, id string) (*entity.Estate, error, int)
}

type estateService struct {
	repo repository.EstateRepository
}

func NewEstateService(r repository.EstateRepository) EstateService {
	return &estateService{
		repo: r,
	}
}

func (s *estateService) CreateEstate(ctx context.Context, estate *entity.Estate) (string, error, int) {
	if estate.Width <= 0 || estate.Length <= 0 || estate.Width > 50000 || estate.Length > 50000 {
		return "", errors.New("invalid estate dimensions"), http.StatusInternalServerError
	}
	res, err := s.repo.CreateEstate(ctx, estate)
	if err != nil {
		return "", err, http.StatusInternalServerError
	}

	return res, nil, http.StatusOK
}

func (s *estateService) GetEstateByID(ctx context.Context, id string) (*entity.Estate, error, int) {
	var data *entity.Estate
	data, err := s.repo.GetEstateByID(ctx, id)
	if err != nil {
		return data, err, http.StatusInternalServerError
	}
	return data, nil, http.StatusOK
}
