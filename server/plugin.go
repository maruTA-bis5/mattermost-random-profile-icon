package main

import (
	"fmt"
	"sync"

	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
)

type Plugin struct {
	plugin.MattermostPlugin

	// configurationLock synchronizes access to the configuration.
	configurationLock sync.RWMutex

	// configuration is the active plugin configuration. Consult getConfiguration and
	// setConfiguration for usage.
	configuration *configuration
}

// UserHasBeenCreated hook will update user profile image to generated one.
func (p *Plugin) UserHasBeenCreated(c *plugin.Context, user *model.User) {
	p.API.LogInfo("UserHasBeenCreated called")
	defer p.API.LogInfo("UserHasBeenCreated exited")
	p.configurationLock.RLock()
	defer p.configurationLock.RUnlock()

	imageProvider := NewImageProvider(p.getConfiguration().ImageProvider)
	if imageProvider == nil {
		// TODO error log and return
		p.API.LogError(fmt.Sprintf("ImageProvider(%s) does not instantiated. abort", p.configuration.ImageProvider))
		return
	}
	data := imageProvider.GenerateProfileImage(user)
	if data == nil {
		// no profile image generated. abort
		p.API.LogWarn(fmt.Sprintf("ImageProvider(%s) does not generate profile image. abort", p.configuration.ImageProvider))
		return
	}
	uid := user.Id
	err := p.API.SetProfileImage(uid, data)
	if err != nil {
		// TODO error log
		p.API.LogError(fmt.Sprintf("ERR: %s", err))
		return
	}
}
