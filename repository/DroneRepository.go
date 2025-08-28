package repository

import (
	"context"
	"gorm.io/gorm"
	"test_sawit_pro/entity"
)

type DroneRepository interface {
	GetEstateSize(ctx context.Context, estateID string) (int, int, error)
	GetTrees(ctx context.Context, estateID string) ([]entity.Tree, error)
}

type droneConnection struct {
	connection *gorm.DB
}

func NewDroneRepository(db *gorm.DB) DroneRepository {
	return &droneConnection{
		connection: db,
	}
}

func (r *droneConnection) GetEstateSize(ctx context.Context, estateID string) (int, int, error) {
	var estate entity.Estate
	if err := r.connection.WithContext(ctx).First(&estate, "id = ?", estateID).Error; err != nil {
		return 0, 0, err
	}
	return estate.Width, estate.Length, nil
}

func (r *droneConnection) GetTrees(ctx context.Context, estateID string) ([]entity.Tree, error) {
	var trees []entity.Tree
	if err := r.connection.WithContext(ctx).
		Where("estate_id = ?", estateID).
		Find(&trees).Error; err != nil {
		return nil, err
	}
	return trees, nil
}
