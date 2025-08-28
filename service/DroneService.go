package service

import (
	"context"
	"math"
	"net/http"
	"test_sawit_pro/dto/response"
	"test_sawit_pro/repository"
)

type DroneService interface {
	CalculateDroneDistance(ctx context.Context, estateID string) (int, error, int)
	MaxDistance(ctx context.Context, estateID string, maxDistance *int) (coordinate response.Coordinate, err error, status int)
}

type droneService struct {
	repo repository.DroneRepository
}

func NewDroneService(r repository.DroneRepository) DroneService {
	return &droneService{
		repo: r,
	}
}

func (s *droneService) CalculateDroneDistance(ctx context.Context, estateID string) (int, error, int) {
	width, length, err := s.repo.GetEstateSize(ctx, estateID)
	if err != nil {
		return 0, err, http.StatusInternalServerError
	}
	trees, err := s.repo.GetTrees(ctx, estateID)
	if err != nil {
		return 0, err, http.StatusInternalServerError
	}

	// Hitung jarak horizontal (zig-zag setiap baris)
	horizontal := (width - 1) * 10 * length

	// Hitung vertical naik/turun dari 0 ke setiap plot lalu turun ke 0 lagi
	treeMap := make(map[[2]int]int)
	for _, t := range trees {
		treeMap[[2]int{t.X, t.Y}] = t.Height
	}

	currentHeight := 0
	vertical := 0
	for y := 1; y <= length; y++ {
		var xStart, xEnd, xStep int
		if y%2 == 1 {
			xStart, xEnd, xStep = 1, width+1, 1
		} else {
			xStart, xEnd, xStep = width, 0, -1
		}
		for x := xStart; x != xEnd; x += xStep {
			h := 1
			if tree, ok := treeMap[[2]int{x, y}]; ok {
				h = tree + 1
			}
			vertical += int(math.Abs(float64(currentHeight - h)))
			currentHeight = h
		}
	}
	vertical += currentHeight

	calculate := horizontal + vertical

	return calculate, nil, http.StatusOK
}

func (s *droneService) MaxDistance(ctx context.Context, estateID string, maxDistance *int) (coordinate response.Coordinate, err error, status int) {
	width, length, err := s.repo.GetEstateSize(ctx, estateID)
	if err != nil {
		return coordinate, err, http.StatusInternalServerError
	}

	trees, err := s.repo.GetTrees(ctx, estateID)
	if err != nil {
		return coordinate, err, http.StatusInternalServerError
	}

	// Vertical: up/down height adjustments
	treeMap := make(map[[2]int]int)
	for _, t := range trees {
		treeMap[[2]int{t.X, t.Y}] = t.Height
	}

	currentHeight := 0
	done := false
	totalDistance := 0

	for y := 1; y <= length && !done; y++ {
		var xStart, xEnd, xStep int
		if y%2 == 1 {
			xStart, xEnd, xStep = 1, width+1, 1
		} else {
			xStart, xEnd, xStep = width, 0, -1
		}
		for x := xStart; x != xEnd; x += xStep {
			if x != xStart {
				totalDistance += 10
			}

			// Vertical movement
			h := 1
			if tree, ok := treeMap[[2]int{x, y}]; ok {
				h = tree + 1
			}
			totalDistance += int(math.Abs(float64(currentHeight - h)))
			currentHeight = h

			// Check maxDistance
			if maxDistance != nil && totalDistance > *maxDistance {
				coordinate = response.Coordinate{X: x, Y: y, TotalDistance: totalDistance}
				done = true
				break
			}
		}
	}

	// Turun ke permukaan
	if !done {
		totalDistance += currentHeight
		if maxDistance != nil && totalDistance > *maxDistance {
			coordinate = response.Coordinate{X: width, Y: length, TotalDistance: totalDistance}
		} else if maxDistance != nil {
			coordinate = response.Coordinate{X: width, Y: length, TotalDistance: totalDistance}
		}
	}

	return coordinate, nil, http.StatusOK
}
