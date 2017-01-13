package utils
import (
	"fmt"
	"net"
	"encoding/json"
)

/*
	in this part, we try to decouple the whole code by a route-controller structure;
	before this server running, all the controller would be written in the router by function init();
	when the client send a json, this server decode this json and decide which controller to process this message;

	我在Server的内部加入一层Router,通过Router对通过Socket发来的信息，通过我们设定的规则进行解析判断后，调用相关的Controller进行任务的分发处理。
	在这个过程中不仅Controller彼此独立，匹配规则和Controller之间也是相互独立的。

*/


type Msg struct {
	Meta   map[string]interface{} `json:"meta"`
	Content interface{}            `json:"content"`
}


type Controller interface {
	Excute(message Msg) []byte
}

var routers [][2]interface{}

func Route(pred interface{} ,controller Controller) {
	switch pred.(type) {
	case func(entry Msg)bool:{
		var arr [2]interface{}
		arr[0] = pred
		arr[1] = controller
		routers = append(routers,arr)
	}
	case map[string]interface{}:{
		defaultPred:= func(entry Msg)bool{
			for keyPred , valPred := range pred.(map[string]interface{}){
				val, ok := entry.Meta[keyPred]
				if !ok {
					return false
				}
				if val != valPred {
					return false
				}
			}
			return true
		}
		var arr [2]interface{}
		arr[0] = defaultPred
		arr[1] = controller
		routers = append(routers,arr)
		fmt.Println(routers)
	}
	default:
		fmt.Println("didn't find requested controller")
	}
}


func TaskDeliver(postdata []byte,conn net.Conn){
	for _ ,v := range routers{
		pred := v[0]
		act := v[1]
		var entermsg Msg
		err := json.Unmarshal(postdata,&entermsg)
		if err != nil {
			Log(err)
		}
		if pred.(func(entermsg Msg)bool)(entermsg) {
			result := act.(Controller).Excute(entermsg)
			conn.Write(result)
			return
		}
	}
}

/*
	this is a sample of how to setup a controller;
	please pay attention: all the controller must be registered in the function init()

	一个controller实例, 注意： 所有的controller必须在init()函数内注册后才能被router分配
*/


type EchoController struct  {

}

func (this *EchoController) Excute(message Msg)[]byte {
	mirrormsg,err :=json.Marshal(message)
	Log("echo the message:", string(mirrormsg))
	CheckError(err)
	return mirrormsg
}


func init() {
	var echo EchoController
	routers = make([][2]interface{} ,0 , 20)
	Route(func(entry Msg)bool{
		if entry.Meta["meta"]=="test"{
			return true}
		return  false
	},&echo)
}