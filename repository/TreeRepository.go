package repository

import (
	"context"
	"gorm.io/gorm"
	"test_sawit_pro/entity"
)

type TreeRepository interface {
	AddTree(ctx context.Context, tree *entity.Tree) (string, error)
	GetHeightsByEstateID(ctx context.Context, estateID string) ([]int, error)
	IsPlotOccupied(ctx context.Context, estateID string, x, y int) (bool, error)
}

type treeConnection struct {
	connection *gorm.DB
}

func NewTreeRepository(db *gorm.DB) TreeRepository {
	return &treeConnection{
		connection: db,
	}
}

func (r *treeConnection) AddTree(ctx context.Context, tree *entity.Tree) (string, error) {
	occupied, err := r.IsPlotOccupied(ctx, tree.EstateID, tree.X, tree.Y)
	if err != nil {
		return "", err
	}
	if occupied {
		return "", gorm.ErrDuplicatedKey
	}
	if err := r.connection.WithContext(ctx).Create(&tree).Error; err != nil {
		return "", err
	}
	return tree.ID, nil
}

func (r *treeConnection) GetHeightsByEstateID(ctx context.Context, estateID string) ([]int, error) {
	var heights []int
	if err := r.connection.WithContext(ctx).
		Model(&entity.Tree{}).
		Where("estate_id = ?", estateID).
		Pluck("height", &heights).Error; err != nil {
		return nil, err
	}
	return heights, nil
}

func (r *treeConnection) IsPlotOccupied(ctx context.Context, estateID string, x, y int) (bool, error) {
	var count int64
	if err := r.connection.WithContext(ctx).
		Model(&entity.Tree{}).
		Where("estate_id = ? AND x = ? AND y = ?", estateID, x, y).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
