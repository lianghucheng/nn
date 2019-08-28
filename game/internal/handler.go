package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"reflect"
	"nn/msg"
)

func init(){
	handler(&msg.Test{},handleTest)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleTest(args []interface{}){
	a:=args[0].(gate.Agent)
	log.Debug("123")
	a.Close()
	log.Debug("123")
}