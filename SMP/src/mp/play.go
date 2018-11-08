package mp

import "fmt"

type Player interface {
	Play(source string)
}

func Play(source, mtype string) { //接口的意义何在？？就是为了省了一个string吗？？这两个play方法直接没有关系，没有实现这个player接口。
	var p Player

	switch mtype {
	case "MP3":
		p = &MP3Player{}  //interface接口类型的意义就是描述数据类型的行为，以及数据类型的共性特征
	case "WAV":
		p = &WAVPlayer{}  //接口是引用类型
	default:
		fmt.Println("Unsupported music type", mtype)
		return
	}
	p.Play(source)
}