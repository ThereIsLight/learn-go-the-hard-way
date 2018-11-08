package cg

import (
	"ipc"
	"sync"
	"encoding/json"
	"errors"
)

/*
中央服务器为全局唯一实例，从原则上需要承担以下责任：
- 在线玩家的状态管理
- 服务器管理(没有实现其他服务器，所以服务器管理这一块先空着)
- 聊天系统(聊天系统也先只实现了广播。要实现房间内聊天或者私聊，其实都可以根据当前的实现进行扩展)
*/

type Message struct {
	From string `json:"from"`  //可以理解为对结构体子段的注释，但是对反射机制可见。
	To string `json:"to"`
	Content string `json:"content"`
}

type CenterServer struct {
	severs map[string] ipc.Server
	players []*Player
	// rooms []*Room  // ?? 可能没写这部分代码
	mutex sync.Mutex  // 互斥锁
}

func NewCenterServer() *CenterServer {
	servers := make(map[string] ipc.Server)  // 为什么这个地方没有0,加上0也没有报错。
	players := make([]*Player, 0)
	//room
	return &CenterServer{severs:servers, players:players}
}
func (server *CenterServer)addPlayer(params string) error {
	player := NewPLayer()
	err := json.Unmarshal([]byte(params), &player)  // []byte(params)讲字符串转化为byte数组吗？？如何进行转化。
	// Unmarshal将json字符串解码到相应的数据结构
	if err != nil {
		return err
	}
	server.mutex.Lock()  //不允许其他的goroutine进行读写操作
	defer server.mutex.Unlock()  // 结束方法的时候关闭锁。
	server.players = append(server.players, player)  // 所有的数据都在server这个
	return nil
}
func (server *CenterServer)removePlayer(params string) error {
	//player := NewPLayer()
	////1.Unmarshal的第一个参数是json字符串，第二个参数是接受json解析的数据结构。
	////第二个参数必须是指针，否则无法接收解析的数据
	//err := json.Unmarshal([]byte(params), &player) // //player本身就是一个指针了，还要去取他的地址吗？第二个参数要求一定是指针。
	//if err != nil {
	//	return nil
	//}
	//for i,v := range server.players  {  // 测试的server是一个全局的示例
	//	if player.Name == v.Name {
	//		server.players = append(server.players[0:i], server.players[i+1:]...)
	//	}
	//}

	//上面是我自己写的代码，主要问题是：
	// 1.没有人规定params就一定是json字符串。
	// 2.这种删除操作一定要考虑到锁。
	server.mutex.Lock()
	defer server.mutex.Unlock()
	for i, v := range server.players {
		if v.Name == params {
			if len(server.players) == 1 {  //这里不用考虑0的情况吗？？还是说在建立server这个实例的时候，就已经确保了server.players是非空的？？
				server.players = make([]*Player, 0)
			}else if i == len(server.players) -1 {
				server.players = server.players[:i]  // 这里直接用赋值就可以了
				// server.players = append(server.players[:i])
			}else if i == 0 {
				server.players = server.players[1:]
			}else {
				server.players = append(server.players[0:i], server.players[i+1:]...)
			}
		}
		return nil
	}
	return errors.New("Player not found")
}

func (server *CenterServer) listPlayer(params string) (players string, err error) {
	server.mutex.Lock()
	defer server.mutex.Unlock()
	if len(server.players) > 0 {
		b, _ := json.Marshal(server.players)  //参数为接口，为什么能够接收一个切片类型，是因为所有的类型都实现了空接口吗？？
		players = string(b)  //不用定义players，也不用返回players!!!!
	}else {
		errors.New("NO Player Online")
	}
	return
}

func (server *CenterServer) broadcast (params string) error {
	var message Message
	err := json.Unmarshal([]byte(params), &message)
	if err != nil {
		return err
	}
	server.mutex.Lock()
	defer server.mutex.Unlock()
	if len(server.players) > 0 {
		for _, player := range server.players {
			// player.mq = &message
			player.mq <- &message  // chan
		}
	}else {
		err = errors.New("NO Player Online Now")
	}
	return err
}

func (server *CenterServer)Handle(method, params string) *ipc.Response {
	switch method {
	case "addplayer":
		err := server.addPlayer(params)
		if err != nil {
			return &ipc.Response{Code:err.Error()}
		}  // 不用回复200吗？？
	case "removeplayer":
		err := server.removePlayer(params)
		if err != nil {
			return &ipc.Response{Code:err.Error()}
		}  // 不用回复200吗？？
	case "listplayer":
		players, err := server.listPlayer(params)
		if err != nil {
			return &ipc.Response{Code:err.Error()}
		}
		return &ipc.Response{"200", players}
	case "broadcast":
		err := server.addPlayer(params)
		if err != nil {
			return &ipc.Response{Code:err.Error()}
		}
		return &ipc.Response{Code:"200"}
	default :
		return &ipc.Response{Code:"404", Body:method + ":" + params}
	}
}

func (server *CenterServer)Name() string {
	return "CenterServer"
}