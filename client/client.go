package main

import (

	"fmt"
	"net"
	"os"
	"time"
	"strconv"
	"./utils"
	"encoding/json"
)


type Msg struct {
	Meta   map[string]interface{} `json:"meta"`
	Content interface{}            `json:"content"`
}

func send(conn net.Conn) {
	for i := 0; i <6; i++ {
		session:=GetSession()
		message := &Msg{
			Meta:map[string]interface{}{
			"meta":"test",
			"ID":strconv.Itoa(i),
			},
		Content: Msg{
			Meta:map[string]interface{}{
			"author":"nucky lu",
			},
		Content:session,
		},
		}
		result,_ :=	json.Marshal(message)
		conn.Write(utils.Enpack((result)))
		//conn.Write([]byte(message))
		time.Sleep(1 * time.Second)
	}
	fmt.Println("send over")
	defer conn.Close()
}

func GetSession() string{
	gs1:=time.Now().Unix()
	gs2:=strconv.FormatInt(gs1,10)
	return gs2
}

func main() {
	server := "localhost:1024"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}


	fmt.Println("connect success")
	send(conn)


}

