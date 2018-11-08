package mlib

import "errors"

type MusicEntry struct {
	Id string
	Name string
	Artist string
	Source string
	Type string
}

type MusicManager struct {
	musics []MusicEntry  // 结构体的命名规范到底是什么？？其中的变量到底是大写还是小写？？
}

func NewMusicManager() *MusicManager {  //为什么这里的返回值一定要是指针呢？？
	return &MusicManager{make([]MusicEntry, 0)}  //MusicManager实际上是对结构体进行赋值的语句，只不过这个结构里面只有一个变量。
}
// 补充make与new
// new使用的并不多，但是不去考虑
// makemake用于内存分配的，但是和new不同，它只用于chan、map以及切片的内存创建，
// 而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。

func (m *MusicManager) Len() int {
	return len(m.musics)
}

func (m *MusicManager) Get(index int) (music *MusicEntry, err error) {
	if index<0 || index>=len(m.musics) {
		return nil, errors.New("Index out of range.")
	}
	return &m.musics[index], nil
}

func (m *MusicManager) Find(name string) *MusicEntry {
	if len(m.musics) == 0 {
		return nil
	}
	for _, m := range m.musics {  // the first variable is index
		if m.Name == name {
			return &m
		}
	}
	return nil
}

func (m *MusicManager) Add(music *MusicEntry) {
	m.musics = append(m.musics, *music)
}

func (m *MusicManager) Remove(name string) *MusicEntry {
	if len(m.musics) == 0 {
		return nil
	}
	var index = 0
	for i, m := range m.musics {  // the first variable is index
		if m.Name == name {
			index = i
		}
	}
	removedMusic := &m.musics[index]
	m.musics = append(m.musics[:index], m.musics[index+1:]...)
	return removedMusic
}
// append函数：
// func append(s []T, x ...T) []T
// append函数将 x追加到切片 s的末尾，并且在必要的时候增加容量。
// 如果是要将一个切片追加到另一个切片尾部，需要使用 ...语法将第2个参数展开为参数列表。
