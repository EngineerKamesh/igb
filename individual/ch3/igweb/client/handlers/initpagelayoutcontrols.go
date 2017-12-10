package handlers

import (
	"github.com/EngineerKamesh/igb/igweb/client/common"
	"github.com/EngineerKamesh/igb/igweb/shared/cogs/notify"
)

func InitializePageLayoutControls(env *common.Env) {

	n := notify.NewNotify()
	err := n.Start()
	if err != nil {
		println("Error encountered when attempting to start the notify cog: ", err)
	}

}
