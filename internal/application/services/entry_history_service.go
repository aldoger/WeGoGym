package services

import (
	"go-kpl/internal/application/dto"
	"go-kpl/internal/domain/repository"

	"golang.org/x/net/context"
)

type (
	EntryHistoryService interface {
		GetEntryAll(ctx context.Context) ([]dto.EntryHistoryAllResponseDto, error)
		GetEntryByUserId(ctx context.Context, userId string) ([]dto.EntryHistoryUserResponseDto, error)
	}

	entryHistoryService struct {
		entryHistoryRepository repository.EntryHistoryRepository
		userRepository         repository.UserRepository
	}
)

func NewEntryHistoryService(entryHistoryRepository repository.EntryHistoryRepository, userRepository repository.UserRepository) EntryHistoryService {
	return &entryHistoryService{entryHistoryRepository: entryHistoryRepository, userRepository: userRepository}
}

func (s *entryHistoryService) GetEntryAll(ctx context.Context) ([]dto.EntryHistoryAllResponseDto, error) {

	allEntry, err := s.entryHistoryRepository.GetEntryAll(ctx, nil)
	if err != nil {
		return []dto.EntryHistoryAllResponseDto{}, err
	}

	var results []dto.EntryHistoryAllResponseDto

	for _, entry := range allEntry {

		user, err := s.userRepository.GetById(ctx, nil, entry.UserId.String())
		if err != nil {
			return nil, err
		}

		results = append(results, dto.EntryHistoryAllResponseDto{
			UserName:  user.Username,
			EntryTime: entry.EntryTime,
		})
	}

	return results, nil
}

func (s *entryHistoryService) GetEntryByUserId(ctx context.Context, userId string) ([]dto.EntryHistoryUserResponseDto, error) {

	userEntry, err := s.entryHistoryRepository.GetEntryByUserId(ctx, nil, userId)
	if err != nil {
		return []dto.EntryHistoryUserResponseDto{}, err
	}

	var results []dto.EntryHistoryUserResponseDto
	for _, entry := range userEntry {

		results = append(results, dto.EntryHistoryUserResponseDto{
			EntryTime: entry.EntryTime,
		})
	}

	return results, nil
}
