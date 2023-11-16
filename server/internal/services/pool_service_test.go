package services_test

import (
	"context"
	"gobooru/internal/dtos"
	"gobooru/internal/mocks"
	"gobooru/internal/models"
	"gobooru/internal/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPoolCreation(t *testing.T) {
	poolRepository := mocks.NewMockPoolRepository(t)
	poolService := services.NewPoolService(services.PoolRepositoryConfig{
		PoolRepository: poolRepository,
	})

	ctx := context.Background()
	dto := dtos.CreatePoolDTO{
		Name:        "test_pool",
		Description: "This is a test pool",
	}

	poolRepository.On("Create", ctx, mock.Anything).Return(models.Pool{ID: 1}, nil)

	resp, err := poolService.Create(ctx, dto)
	assert.NoError(t, err)
	assert.Equal(t, resp.Pool.ID, 1)

	poolRepository.AssertExpectations(t)
}
