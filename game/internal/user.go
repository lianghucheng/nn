package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/timer"
	"nn/conf"
	"nn/msg"
	"time"
)

const (
	userLogin = iota
	userLogout
)

const (
	roleRobot = -2
	roleBlack = -1
	rolePlayer = 1
	roleAgent = 2
	roleAdmin = 3
	roleRoot = 4
)

var (
	userIDUsers=make(map[int]*User)
	userIDRooms=make(map[int]*NiuNiuRoom)
	systemOn=true
)

type User struct{
	gate.Agent
	state int
	baseData *BaseData
	heartbeatTimer *timer.Timer
	heartbeatStop bool
}

type BaseData struct{
	userData *UserData
	roomOwnerID int
	taskIDTaskDatas map[int]*TaskData
}

func newUser(a gate.Agent) *User{
	user:=new(User)
	user.Agent=a
	user.state=userLogin
	user.baseData=new(BaseData)
	user.baseData.userData=new(UserData)
	user.baseData.taskIDTaskDatas=make(map[int]*TaskData)
	return user
}

func (user *User)autoHeartbeat(){
	if user.heartbeatStop{
		log.Debug("userID: %v 心跳停止", user.baseData.userData.UserID)
		user.Close()
		return
	}
	user.heartbeatStop=true
	user.WriteMsg(&msg.S2C_Heartbeat{})

	user.heartbeatTimer=skeleton.AfterFunc(time.Duration(conf.GetCfgTimeOut().S2CHeartBeatTimeout)*time.Second, func() {
		user.autoHeartbeat()
	})
}