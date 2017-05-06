package msg

// import (
// 	"github.com/name5566/leaf/network"
// )

import (
    "github.com/name5566/leaf/network/json"
	"fmt"
)

// var Processor network.Processor
var Processor = json.NewProcessor()

func init() {
    Processor.Register(&Hello{})
    Processor.Register(&Register{})

    Processor.Register(&C2S_Chat{})
    Processor.Register(&C2S_EnterRoom{})
    Processor.Register(&C2S_LeaveRoom{})

    Processor.Register(&S2C_Chat{})
    Processor.Register(&S2C_EnterRoom{})
    Processor.Register(&S2C_LeaveRoom{})
}


type Hello struct {
    Name string
}

type Register struct {
    Name string
    Age int
    Male bool
}

func DebugString(obj interface{}) string {
    return ""
}

func (obj *Register)String()string  {
    return fmt.Sprintf("Name: %s Age: %d Male: %t", obj.Name, obj.Age, obj.Male)
}
