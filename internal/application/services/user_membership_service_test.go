package services

import (
	"context"
	"testing"
	"time"

	"go-kpl/internal/application/dto"
	"go-kpl/internal/domain/models"
	"go-kpl/internal/domain/repository"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userMembershipRepoMock    = &repository.UserMembershipRepositoryMock{}
	userRepositoryMock        = &repository.UserRepositoryMock{}
	membershipRepositoryMock  = &repository.MembershipRepositoryMock{}
	userMembershipServiceTest = userMembershipService{
		userMembershipRepository: userMembershipRepoMock,
		userRepository:           userRepositoryMock,
		membershipRepository:     membershipRepositoryMock,
	}
)

func TestUserMembershipService_CreateUserMembership(t *testing.T) {
	ctx := context.Background()

	userId := uuid.New()
	membershipId := uuid.New()

	mockMembership := models.Membership{
		Id:       membershipId,
		Type:     "Premium",
		Duration: 30,
		Price:    100000,
	}

	expectedExpiredAt := time.Now().Add(time.Hour * 24 * time.Duration(mockMembership.Duration))

	mockUserMembership := models.UserMembership{
		UserId:    userId,
		MemberId:  membershipId,
		ExpiredAt: expectedExpiredAt,
	}

	membershipRepositoryMock.Mock.On("GetById", ctx, mock.Anything, membershipId.String()).
		Return(mockMembership, nil)

	userMembershipRepoMock.Mock.On("Create", ctx, mock.Anything, mock.MatchedBy(func(um models.UserMembership) bool {
		return um.UserId == userId &&
			um.MemberId == membershipId
	})).Return(mockUserMembership, nil)

	result, err := userMembershipServiceTest.CreateUserMembership(ctx, dto.CreateUserMembershipRequestDto{
		UserId:       userId.String(),
		MembershipId: membershipId.String(),
	})

	// Assert
	assert.NoError(t, err)
	assert.WithinDuration(t, expectedExpiredAt, result.ExpiredAt, time.Second)

}
