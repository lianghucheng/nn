package gate

import (
	"nn/game"
	"nn/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.Test{},game.ChanRPC)
}
