package main

import (
	"github.com/mattermost/mattermost-server/model"
)

type DefaultProvider struct{}

func (provider *DefaultProvider) GenerateProfileImage(user *model.User) []byte {
	// TODO
	return nil
}
