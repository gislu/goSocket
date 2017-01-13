package utils

import (
	"net"
	"time"
)


//HeartBeating, determine if client send a message within set time by GravelChannel
// 心跳计时，根据GravelChannel判断Client是否在设定时间内发来信息

func HeartBeating(conn net.Conn, readerChannel chan byte,timeout int) {
	select {
	case _ = <-readerChannel:
		Log(conn.RemoteAddr().String(), "get message, keeping heartbeating...")
		conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
		break
	case <-time.After(time.Second*5):
		Log("It's really weird to get Nothing!!!")
		conn.Close()
	}

}

func GravelChannel(n []byte,mess chan byte){
	for _ , v := range n{
		mess <- v
	}
	close(mess)
}

