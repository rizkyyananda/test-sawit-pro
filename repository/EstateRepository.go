package repository

import (
	"context"
	"gorm.io/gorm"
	"test_sawit_pro/entity"
)

type EstateRepository interface {
	CreateEstate(ctx context.Context, estate *entity.Estate) (string, error)
	GetEstateSize(ctx context.Context, id string) (int, int, error)
	GetEstateByID(ctx context.Context, id string) (*entity.Estate, error)
}

type estateConnection struct {
	connection *gorm.DB
}

func NewEstateRepository(db *gorm.DB) EstateRepository {
	return &estateConnection{
		connection: db,
	}
}

func (r *estateConnection) CreateEstate(ctx context.Context, estate *entity.Estate) (string, error) {
	if err := r.connection.WithContext(ctx).Create(&estate).Error; err != nil {
		return "", err
	}
	return estate.ID, nil
}

func (r *estateConnection) GetEstateSize(ctx context.Context, id string) (int, int, error) {
	var estate entity.Estate
	if err := r.connection.WithContext(ctx).First(&estate, "id = ?", id).Error; err != nil {
		return 0, 0, err
	}
	return estate.Width, estate.Length, nil
}

func (r *estateConnection) GetEstateByID(ctx context.Context, id string) (*entity.Estate, error) {
	var estate entity.Estate
	if err := r.connection.WithContext(ctx).First(&estate, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &estate, nil
}
