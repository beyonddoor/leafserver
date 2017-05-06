package internal

import (
    // "github.com/name5566/leaf/log"
    "github.com/name5566/leaf/gate"
    "server/msg"
	"server/game/logic"
)

func init() {
    handler(&msg.C2S_EnterRoom{}, handle_C2S_EnterRoom)
    handler(&msg.C2S_LeaveRoom{}, handle_C2S_LeaveRoom)
	handler(&msg.C2S_Chat{}, handle_C2S_Chat)
}

func handle_C2S_EnterRoom(args []interface{}) {
}

func handle_C2S_LeaveRoom(args []interface{})  {
}

func handle_C2S_Chat(args []interface{}) {
    m := args[0].(*msg.C2S_Chat)
    a := args[1].(gate.Agent)
	logic.AddToRoom("def", &logic.User{"", a})
	logic.SendAll("def", &msg.S2C_Chat{"", m.Text})
}