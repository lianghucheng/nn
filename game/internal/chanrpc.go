package internal

import (
	"github.com/name5566/leaf/gate"
	"nn/conf"
	"nn/msg"
	"strings"
	"time"
)

type AgentInfo struct{
	user *User
}

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
	skeleton.RegisterChanRPC("TokenAuthorize",rpcTokenAuthorize)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)

	a.SetUserData(new(AgentInfo))

	skeleton.AfterFunc(time.Duration(conf.GetCfgTimeOut().C2SConnectTimeout)*time.Second, func() {
		if a.UserData()!=nil{
			agentInfo:=a.UserData().(*AgentInfo)
			if agentInfo!=nil&&agentInfo.user==nil{
				a.Close()
			}
		}
	})
	a.Close()
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	user:=a.UserData().(*AgentInfo).user
	a.SetUserData(nil)
	if user==nil{
		return
	}
	if user.state==userLogin{
		user.state=userLogout
		//user.logout()
	}
}

func rpcTokenAuthorize(args []interface{}){
	a:=args[0].(gate.Agent)
	m:=args[1].(*msg.C2S_TokenAuthorize)
	if a.UserData()==nil||a.UserData().(*AgentInfo).user!=nil{
		a.Close()
		return
	}
	if strings.TrimSpace(m.Token)==""{
		a.WriteMsg(&msg.S2C_Close{Error:msg.S2C_Close_TokenInvalid})
		a.Close()
		return
	}
	if !systemOn{
		a.WriteMsg(&msg.S2C_Close{Error:msg.S2C_Close_SystemOff})
		a.Close()
		return
	}
	newUser:=newUser(a)
	a.UserData().(*AgentInfo).user=newUser
	newUser.tokenAuthorize(m)
}