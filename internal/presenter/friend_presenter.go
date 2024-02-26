package presenter

import (
	"fmt"
	"kominfo-go/internal/entities"
)

type FriendPresenter interface {
	ShowFriend(friend entities.Friend)
}

type friendPresenter struct{}

func NewFriendPresenter() FriendPresenter {
	return &friendPresenter{}
}

func (p *friendPresenter) ShowFriend(friend entities.Friend) {
	fmt.Println("Data teman dengan absen", friend.Absen)
	fmt.Println("Nama:", friend.Nama)
	fmt.Println("Alamat:", friend.Alamat)
	fmt.Println("Pekerjaan:", friend.Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang:", friend.Alasan)
}
