package gate

import (
	"nn/game"
	"nn/login"
	"nn/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.C2S_TokenAuthorize{},login.ChanRPC)

	msg.Processor.SetRouter(&msg.C2S_Heartbeat{},game.ChanRPC)
}
