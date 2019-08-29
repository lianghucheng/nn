package internal

import (
	"github.com/name5566/leaf/gate"
	"nn/msg"
	"reflect"
)

func init(){
	handler(&msg.C2S_Heartbeat{},handleHeartbeat)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleHeartbeat(args []interface{}){
	a:=args[1].(gate.Agent)
	if a.UserData()==nil{
		a.Close()
		return
	}
	user:=a.UserData().(*AgentInfo).user
	if user == nil{
		a.Close()
		return
	}
	user.heartbeatStop = false
}
