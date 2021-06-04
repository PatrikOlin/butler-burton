package cmd

import (
	"github.com/PatrikOlin/butler-burton/util"
)

func Auth() error {
	util.SharepointAuth()
	return nil
}
