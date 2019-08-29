package internal

import "github.com/name5566/leaf/timer"

type NiuNiuRoom struct {
	room
	//rule              *poker.Rule
	userIDPlayerDatas map[int]*PlayerData // Key: userID
	cards             []int               // 洗好的牌
	rests             []int               // 剩余的牌

	bidUserIDs   []int // 叫庄的玩家 userID
	dealerUserID int   // 庄家 userID，初始化为 -1

	startTimestamp int64
	startTimer     *timer.Timer
	prepareTimer   *timer.Timer
	bidTimer       *timer.Timer
	doubleTimer    *timer.Timer
	showTimer      *timer.Timer

	winnersPos []int
	losersPos  []int

	nextWaitTime int
}

// 玩家数据
type PlayerData struct {
	user      *User
	state     int // 玩家状态
	gameState int // 玩家游戏中的状态
	position  int // 用户在桌子上的位置，从 0 开始

	owner     bool  // 房主
	dealer    bool  // 庄
	double    bool  // 加倍
	show      bool  // 展示牌
	showFifth bool  // 展示第五张牌
	multiple  int   // 倍数
	gain      int64 // 获得的筹码（底分x倍数）

	hands    []int // 手牌
	//analyzer *poker.Analyzer

	actionTimestamp int64 // 记录房间动作时间戳

	vipChips  int64
	redPacket float64

	online bool
}