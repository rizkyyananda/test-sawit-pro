package service

import (
	"context"
	"errors"
	"net/http"
	"sort"
	"test_sawit_pro/dto/response"
	"test_sawit_pro/entity"
	"test_sawit_pro/repository"
)

type TreeService interface {
	AddTree(ctx context.Context, tree *entity.Tree) (string, error, int)
	GetTreeStats(ctx context.Context, estateID string) (*response.TreeStats, error, int)
}

type treeService struct {
	treeRepo   repository.TreeRepository
	estateRepo repository.EstateRepository
}

func NewTreeService(treeRepo repository.TreeRepository, estateRepo repository.EstateRepository) TreeService {
	return &treeService{
		treeRepo:   treeRepo,
		estateRepo: estateRepo,
	}
}

func (s *treeService) AddTree(ctx context.Context, tree *entity.Tree) (string, error, int) {
	// Validasi koordinat & tinggi
	if tree.X <= 0 || tree.Y <= 0 || tree.Height < 1 || tree.Height > 30 {
		return "", errors.New("invalid tree input"), http.StatusBadRequest
	}

	estate, err := s.estateRepo.GetEstateByID(ctx, tree.EstateID)
	if err != nil {
		return "", errors.New("estate not found"), http.StatusNotFound
	}
	if tree.X > estate.Width || tree.Y > estate.Length {
		return "", errors.New("tree coordinates out of bounds"), http.StatusBadRequest
	}

	res, err := s.treeRepo.AddTree(ctx, tree)
	if err != nil {
		return "", err, http.StatusNotFound
	}
	return res, nil, http.StatusOK
}

func (s *treeService) GetTreeStats(ctx context.Context, estateID string) (*response.TreeStats, error, int) {
	heights, err := s.treeRepo.GetHeightsByEstateID(ctx, estateID)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}

	stats := &response.TreeStats{}
	if len(heights) == 0 {
		return stats, errors.New("data not found"), http.StatusNotFound
	}

	// Hitung count, min, max, median
	stats.Count = len(heights)
	stats.Min, stats.Max = heights[0], heights[0]
	sum := 0
	for _, h := range heights {
		sum += h
		if h < stats.Min {
			stats.Min = h
		}
		if h > stats.Max {
			stats.Max = h
		}
	}

	sorted := append([]int{}, heights...)
	sort.Ints(sorted)
	mid := len(sorted) / 2
	if len(sorted)%2 == 0 {
		stats.Median = (sorted[mid-1] + sorted[mid]) / 2
	} else {
		stats.Median = sorted[mid]
	}

	return stats, nil, http.StatusOK
}
