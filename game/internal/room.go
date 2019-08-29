package internal
type room struct {
	state           int
	positionUserIDs map[int]int // key: 座位号, value: userID
	creatorUserID   int         // 创建者 userID
	ownerUserID     int         // 房主 userID
	number          string
	desc            string
}