package msg

import (
	"github.com/name5566/leaf/network/json"
)

var Processor =json.NewProcessor()

func init() {
	Processor.Register(&C2S_TokenAuthorize{})

	Processor.Register(&C2S_Heartbeat{})
}

type C2S_Heartbeat struct{}

type S2C_Heartbeat struct{}