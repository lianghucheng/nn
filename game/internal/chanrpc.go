package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)

	log.Debug("111")
	a.Close()
	log.Debug("222")
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	log.Debug("111")
	a.Close()
	log.Debug("222")
}
