package cmd

import (
	log "github.com/Sirupsen/logrus"
	"github.com/t-tran/lucas/backend"
	"github.com/t-tran/lucas/backend/types"
)

type CommandLine struct {
	Opts *types.Opts
}

func (cl *CommandLine) Run() {
	if &cl.Opts != nil && cl.Opts.DebugMode == true {
		log.SetLevel(log.DebugLevel)
	}
	lbe := &backend.LucasServer{}
	lbe.Start(cl.Opts)
}
