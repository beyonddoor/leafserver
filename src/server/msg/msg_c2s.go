package msg

type C2S_Chat struct {
	Text string 
}

type C2S_EnterRoom struct {
	RoomId string
}

type C2S_LeaveRoom struct {
	RoomId string
}

