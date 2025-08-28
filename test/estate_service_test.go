package service_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"test_sawit_pro/entity"
	"test_sawit_pro/mocks"
	"test_sawit_pro/service"
)

func TestCreateEstate_Success(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockEstateRepository(ctrl)

	estate := &entity.Estate{
		ID:     "123",
		Width:  10,
		Length: 10,
	}

	// ekspektasi: CreateEstate dipanggil sekali dan return "123", nil
	mockRepo.EXPECT().
		CreateEstate(gomock.Any(), estate).
		Return("123", nil).
		Times(1)

	svc := service.NewEstateService(mockRepo)

	id, err, status := svc.CreateEstate(context.Background(), estate)
	fmt.Println(status)
	assert.NoError(t, err)
	assert.Equal(t, "123", id)
}

func TestCreateEstate_Error(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockEstateRepository(ctrl)

	estate := &entity.Estate{
		ID:     "123",
		Width:  10,
		Length: 10,
	}

	expectedErr := errors.New("db insert failed")

	// ekspektasi: CreateEstate return error
	mockRepo.EXPECT().
		CreateEstate(gomock.Any(), estate).
		Return("", expectedErr).
		Times(1)

	svc := service.NewEstateService(mockRepo)

	id, err, status := svc.CreateEstate(context.Background(), estate)
	fmt.Println(status)
	assert.Error(t, err)
	assert.EqualError(t, err, expectedErr.Error())
	assert.Empty(t, id)
}
