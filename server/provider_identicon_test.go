package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/mattermost/mattermost-server/model"
)

func TestGenerateIdenticon(t *testing.T) {
	assert := assert.New(t)
	provider := NewImageProvider("identicon")
	user := &model.User{
		Id: model.NewId(),
	}

	result := provider.GenerateProfileImage(user)
	assert.NotNil(result)
	assert.NotEqual(t, len(result), 0, "[]byte is empty")
	// TODO 画像であることの検証
}
