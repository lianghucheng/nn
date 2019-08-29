package internal

import (
	"github.com/name5566/leaf/log"
	"gopkg.in/mgo.v2/bson"
	"nn/msg"
	"time"
)

func (user *User)tokenAuthorize(m *msg.C2S_TokenAuthorize){
	userData:=new(UserData)
	skeleton.Go(func() {
		db:=mongoDB.Ref()
		defer mongoDB.UnRef(db)

		err:=db.DB(DB).C("users").Find(bson.M{"token":m.Token,"expireat":bson.M{"$gt":time.Now().Unix()},"online":true}).One(userData)
		if err!=nil{
			log.Debug("find token %v error: %v", m.Token, err)
			userData=nil
			user.WriteMsg(&msg.S2C_Close{Error:msg.S2C_Close_TokenInvalid})
			user.Close()
		}
	}, func() {
		if userData==nil||user.state==userLogout{
			return
		}
		if oldUser,ok:=userIDUsers[userData.UserID];ok{
			log.Debug("accountid,account: %v  %v 已经登录", userData.AccountID, userData.Account)
			oldUser.Close()
			oldUser.baseData.userData.Token=user.baseData.userData.Token
			user.baseData=oldUser.baseData
			userData=oldUser.baseData.userData
		}
		userIDUsers[userData.UserID]=user
		user.baseData.userData=userData
		user.autoHeartbeat()
		user.WriteMsg(&msg.S2C_Authorize{})
		if m.Connect{
			if r,ok:=userIDRooms[user.baseData.userData.UserID];ok{
				//user.enterRoom(r)
				log.Debug("account: %v Token验证登录, 游戏在线人数: %v", userData.Account, len(userIDUsers),r)
				return
			}
			user.WriteMsg(&msg.S2C_Close{})
			user.Close()
			return
		}
		log.Debug("account: %v Token验证登录, 游戏在线人数: %v", userData.Account, len(userIDUsers))
		if _,ok:=userIDRooms[user.baseData.userData.UserID];ok{
			//user.enterRoom(r)
		}
	})
}