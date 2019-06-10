package main

import (
	"github.com/mattermost/mattermost-server/model"
)

type ImageProvider interface {
	GenerateProfileImage(*model.User) []byte
}

func NewImageProvider(providerType string) ImageProvider {
	switch providerType {
	case "default":
		return &DefaultProvider{}
	case "identicon":
		return &IdenticonProvider{}
	}
	// TODO default case -> error
	return nil
}
