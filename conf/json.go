package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/name5566/leaf/log"
)

var opts *Config
var Server CfgLeaf

//Config 配置类型
type Config struct{
	CfgLeaf CfgLeaf
	CfgTimeOut CfgTimeOut
	Matchs map[string]Match
	GameTimeout GameTimeout
	Private map[string]Match
	CfgRedis CfgRedis
	CfgNN CfgNN
}

type CfgLeaf struct {
	LogLevel    string
	LogPath     string
	WSAddr      string
	CertFile    string
	KeyFile     string
	TCPAddr     string
	MaxConnNum  int
	ConsolePort int
	ProfilePath string
	DBMaxConnNum int
	DBUrl string
}

type CfgTimeOut struct{
	C2SConnectTimeout int
	S2CHeartBeatTimeout int
}

type Match struct {

}

type GameTimeout struct{

}

type CfgRedis struct{

}

type CfgNN struct{

}

func init() {
	opts=new(Config)
	_,err:=toml.DecodeFile("conf/server.toml",opts)
	if err!=nil{
		log.Fatal("配置文件解析错误:%v", err)
	}
	Server=opts.CfgLeaf
}

func GetCfgTimeOut() *CfgTimeOut{
	return &opts.CfgTimeOut
}

func GetCfgMatch()map[string]Match{
	return opts.Matchs
}

func GetGameTimeout()*GameTimeout{
	return &opts.GameTimeout
}

func GetCfgPrivate()map[string]Match{
	return opts.Private
}

func GetCfgRedis()*CfgRedis{
	return &opts.CfgRedis
}

func GetCfgNN()*CfgNN{
	return &opts.CfgNN
}
