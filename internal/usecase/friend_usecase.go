package usecase

import (
	"kominfo-go/internal/entities"
	"kominfo-go/internal/repositories"
)

type FriendUsecase interface {
	GetFriendByAbsen(absen int) entities.Friend
}

type friendUsecase struct {
	friendRepo repositories.FriendRepository
}

func NewFriendUsecase(repo repositories.FriendRepository) FriendUsecase {
	return &friendUsecase{friendRepo: repo}
}

func (uc *friendUsecase) GetFriendByAbsen(absen int) entities.Friend {
	return uc.friendRepo.GetFriendByAbsen(absen)
}
