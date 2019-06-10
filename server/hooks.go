package main

import (
	"fmt"

	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
)

// UserHasBeenCreated hook will update user profile image to generated one.
func (p *Plugin) UserHasBeenCreated(c *plugin.Context, user *model.User) {
	p.API.LogInfo("UserHasBeenCreated called")
	defer p.API.LogInfo("UserHasBeenCreated exited")
	
	providerName := p.getConfiguration().ImageProvider
	imageProvider := NewImageProvider(providerName)
	if imageProvider == nil {
		// TODO error log and return
		p.API.LogInfo(fmt.Sprintf("ImageProvider(%s) does not instantiated. abort", providerName))
		return
	}
	data := imageProvider.GenerateProfileImage(user)
	if data == nil {
		// no profile image generated. abort
		p.API.LogInfo(fmt.Sprintf("ImageProvider(%s) does not generate profile image. abort", providerName))
		return
	}
	uid := user.Id
	err := p.API.SetProfileImage(uid, data)
	if err != nil {
		// TODO error log
		p.API.LogInfo(fmt.Sprintf("ERR: %s", err))
		return
	}
}