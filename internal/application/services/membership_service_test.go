package services

import (
	"go-kpl/internal/domain/models"
	"go-kpl/internal/domain/repository"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
)

var membershipRepository = &repository.MembershipRepositoryMock{}
var membershipServiceTest = membershipService{membershipRepository: membershipRepository}

func TestMembershipGetById(t *testing.T) {
	ctx := context.Background()
	expectedId := uuid.New().String()

	expectedMembership := models.Membership{
		Id:       uuid.MustParse(expectedId),
		Type:     "classic",
		Duration: 30,
		Price:    150000,
	}

	membershipRepository.Mock.On("GetById", ctx, mock.Anything, expectedId).
		Return(expectedMembership, nil)

	result, err := membershipServiceTest.GetByIdMembership(ctx, expectedId)

	assert.NoError(t, err)
	assert.Equal(t, expectedMembership.Id.String(), result.Id)
	assert.Equal(t, expectedMembership.Type, "classic")
	assert.Equal(t, expectedMembership.Duration, result.Duration)
	assert.Equal(t, expectedMembership.Price, result.Price)

	// Validate expectations
	membershipRepository.Mock.AssertExpectations(t)
}

func TestMembershipService_GetAllMembership(t *testing.T) {
	ctx := context.Background()

	// Mock data dari database
	mockMemberships := []models.Membership{
		{
			Id:       uuid.New(),
			Type:     "Basic",
			Duration: 30,
			Price:    100000,
		},
		{
			Id:       uuid.New(),
			Type:     "Premium",
			Duration: 90,
			Price:    250000,
		},
	}

	// Setup expected calls ke mock
	membershipRepository.Mock.On("GetAll", ctx, mock.Anything).Return(mockMemberships, nil)

	// Call service
	result, err := membershipServiceTest.GetAllMembership(ctx)

	// Assertions
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, mockMemberships[0].Id.String(), result[0].Id)
	assert.Equal(t, mockMemberships[0].Type, result[0].Type)
	assert.Equal(t, mockMemberships[0].Duration, result[0].Duration)
	assert.Equal(t, mockMemberships[0].Price, result[0].Price)

	assert.Equal(t, mockMemberships[1].Id.String(), result[1].Id)
	assert.Equal(t, mockMemberships[1].Type, result[1].Type)
	assert.Equal(t, mockMemberships[1].Duration, result[1].Duration)
	assert.Equal(t, mockMemberships[1].Price, result[1].Price)

	membershipRepository.Mock.AssertExpectations(t)
}
