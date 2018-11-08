package cg

import (
	"ipc"
	"encoding/json"
	"errors"
)

/*
调用这个服务器的功能
*/

type CenterClient struct {
	*ipc.IPCClient  // 这种结构体定义代表着什么？？这两个结构体代表着什么？？有点类似于继承。父类与子类。
	// CenterClient匿名组合了IpcClient，这样就可以直接在代码中直接调用IpcClient的功能了。
}
// 大写的方法名  小写的方法名
// json.Marshal()参数为接口，那么传入结构体实体与传入结构体指针有什么区别吗？
func (client *CenterClient)AddPlayer(player *Player) error {
	b, err := json.Marshal(player)  //
	if err != nil {
		return err
	}
	resp, err := client.Call("addplayer", string(b))  // 为什么可以直接调用别的结构体的方法？？？
	// Call()函数是关键？？？
	if err == nil && resp.Code == "200" {
		return nil
	}
	return err
}

func (client *CenterClient)RemovePLayer(name string) error {
	ret, _ := client.Call("removeplayer", name)  //ret与resq有什么区别？？
	if ret.Code == "200" {
		return nil
	}
	return errors.New(ret.Code)
}

func (client *CenterClient)ListPlayer(params string)(ps []*Player, err error) {
	resp, _ := client.Call("listplayer", params)
	if resp.Code != "200" {
		err = errors.New(resp.Code)
		return
	}
	err = json.Unmarshal([]byte(resp.Code), &ps)
	return
}
func (client *CenterClient)Broadcast(message string) error {
	 m := &Message{Content:message}
	 b, err := json.Marshal(m)
	 if err != nil {
	 	return err
	 }
	 resp, _ := client.Call("broadcast", string(b))
	 if resp.Code == "200" {
	 	return nil
	 }
	 return errors.New(resp.Code)

}