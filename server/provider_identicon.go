package main

import (
	identicon "github.com/dgryski/go-identicon"
	"github.com/mattermost/mattermost-server/model"
)

type IdenticonProvider struct{}

func (provider *IdenticonProvider) GenerateProfileImage(user *model.User) []byte {
	uid := user.Id
	key := []byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF}
	icon := identicon.New5x5(key)
	data := []byte(uid)
	image := icon.Render(data)
	return image
}
