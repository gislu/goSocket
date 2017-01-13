
package utils

import (
	"bytes"
	"encoding/binary"
)


/* A custom communication protocol between server and client;
   All the message will be composed by  header + the length of whole message(4 byte int) + message content;
   If the header is wrong or the length of message doesn't match, the message wouldn't be decoded

   一个简单的通讯协议，由 header + 信息长度 ＋ 信息内容组成
*/

const (
	ConstHeader         = "testHeader"
	ConstHeaderLength   = 10
	ConstMLength = 4
)

func Enpack(message []byte) []byte {
	return append(append([]byte(ConstHeader), IntToBytes(len(message))...), message...)
}


func Depack(buffer []byte) []byte {
	length := len(buffer)

	var i int
	data := make([]byte, 32)
	for i = 0; i < length; i = i + 1 {
		if length < i+ConstHeaderLength+ConstMLength {
			break
		}
		if string(buffer[i:i+ConstHeaderLength]) == ConstHeader {
			messageLength := BytesToInt(buffer[i+ConstHeaderLength : i+ConstHeaderLength+ConstMLength])
			if length < i+ConstHeaderLength+ConstMLength+messageLength {
				break
			}
			data = buffer[i+ConstHeaderLength+ConstMLength : i+ConstHeaderLength+ConstMLength+messageLength]

		}
	}

	if i == length {
		return make([]byte, 0)
	}
	return data
}

func IntToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}


func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}  