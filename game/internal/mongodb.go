package internal

import (
	"github.com/name5566/leaf/db/mongodb"
	"github.com/name5566/leaf/log"
	"nn/conf"
)

var mongoDB *mongodb.DialContext

const DB = "lhc_gametoken"

func init(){
	if conf.Server.DBMaxConnNum<=0{
		conf.Server.DBMaxConnNum=100
	}
	db,err:=mongodb.Dial(conf.Server.DBUrl,conf.Server.DBMaxConnNum)
	if err!=nil{
		log.Fatal("dial mongodb error: %v", err)
	}
	mongoDB=db
	initCollection()
}

func initCollection(){
	var err error
	db:=mongoDB
	err=db.EnsureCounter(DB,"counters","users")
	if err!=nil{
		log.Fatal("ensure counter error: %v", err)
	}
	err=db.EnsureCounter(DB,"counters","configs")
	if err!=nil{
		log.Fatal("ensure counter error: %v",err)
	}
	err=db.EnsureUniqueIndex(DB,"users",[]string{"accountid"})
	if err!=nil{
		log.Fatal("ensure index error: %v",err)
	}
	err=db.EnsureUniqueIndex(DB,"users",[]string{"account"})
	if err!=nil{
		log.Fatal("ensure index error: %v",err)
	}
}

func mongoDBDestroy(){
	mongoDB.Close()
	mongoDB=nil
}

func mongoDBNextSeq(id string)(int,error){
	return mongoDB.NextSeq(DB,"counters",id)
}