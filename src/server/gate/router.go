package gate


import (
    "server/game"
    "server/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.Hello{}, game.ChanRPC)
    msg.Processor.SetRouter(&msg.Register{}, game.ChanRPC)

    msg.Processor.SetRouter(&msg.C2S_Chat{}, game.ChanRPC)
    // msg.Processor.SetRouter(&msg.Register{}, game.ChanRPC)

    msg.Processor.SetRouter(&msg.S2C_Chat{}, game.ChanRPC)
}
