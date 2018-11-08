package cg

import "fmt"

/*
实现在线玩家的管理
 */

type Player struct {
	Name string
	Level int
	Exp int
	Room int
	mq chan *Message
}

//新建一个角色的时候为什么要用协程？？ 为了防止数据的冲突吗？
func NewPLayer() *Player {
	m := make(chan *Message, 1024)  //这个message里面的数据代表着什么？
	player := &Player{"", 0, 0, 0, m}
	go func(p *Player) {  // 这里使用协程的目的是什么？？ 首先协程的目的是什么？？
		for {
			msg := <-p.mq
			fmt.Println(p.Name, "received message:", msg.Content)
		}
	}(player)
	return player
}
