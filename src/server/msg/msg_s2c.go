package msg

type S2C_Chat struct {
	Uid,Text string
}

type S2C_EnterRoom struct {
	Uid, RoomId string
}

type S2C_LeaveRoom struct {
	Uid, RoomId string
}

