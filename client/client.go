package main

import (
    "encoding/binary"
    "net"
	"time"
	"io"
	"os"
	"fmt"
	"bufio"
	"strings"
)

func sendData(data []byte, wr io.Writer)  {
    // len + data
    m := make([]byte, 2+len(data))

    // 默认使用大端序
    binary.BigEndian.PutUint16(m, uint16(len(data)))

    copy(m[2:], data)

    // 发送消息
    wr.Write(m)
}

func chat(msg string, conn io.Writer) {
    sendData([]byte(fmt.Sprintf(`{
        "C2S_Chat": {
            "Text": "%s"
        }
    }`, msg)), conn)
}

func readline(conn io.Writer)  {
    reader := bufio.NewReader(os.Stdin)
    for {
        str, _ := reader.ReadString('\n')
        str = strings.Replace(str, "\n", "", -1)
        chat(str, conn)
    }
}

func main() {
    conn, err := net.Dial("tcp", "127.0.0.1:3563")
    if err != nil {
        panic(err)
    }

    // Hello 消息（JSON 格式）
    // 对应游戏服务器 Hello 消息结构体
    // sendData([]byte(`{
    //     "Hello": {
    //         "Name": "leaf"
    //     }
    // }`), conn)

    // sendData([]byte(`{
    //     "Register" :{
    //         "Name": "jhon",
    //         "Age": 20,
    //         "Male": true
    //         }}`), conn)

    go readline(conn)

    io.Copy(os.Stdout, conn)

    time.Sleep(time.Second * 3)
}