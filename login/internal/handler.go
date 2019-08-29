package internal

import (
	"github.com/name5566/leaf/gate"
	"nn/game"
	"nn/msg"
	"reflect"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handler(&msg.C2S_TokenAuthorize{},handleTokenAuthorize)
}

func handler(m interface{},h interface{}){
	skeleton.RegisterChanRPC(reflect.TypeOf(m),h)
}

func handleTokenAuthorize(args []interface{}){
	m:=args[0].(*msg.C2S_TokenAuthorize)
	a:=args[1].(gate.Agent)
	game.ChanRPC.Go("TokenAuthorize",a,m)
}

