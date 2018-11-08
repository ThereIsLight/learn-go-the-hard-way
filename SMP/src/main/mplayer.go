package main
import(
	"mlib"
	"fmt"
	"strconv"
	"mp"
	"bufio"
	"os"
	"strings"
)
// 为什么mlib是红色的。解释：需要在GoPath下面添加项目的路径名。
var lib *mlib.MusicManager  // 为什么前面是个包名，同一个包下面是不是不能函数、变量、结构体等不能同名？？
var id int = 1  //??
var ctrl, signal chan int  // 应该是为了携程而设计的，但是目前的代码根本没有用到。

func handleLibCommands(tokens []string) {  //t
	switch tokens[1] {
	case "list":
		if lib.Len() == 0 {
			fmt.Println("There is no music in the library")
		}
		for i:=0;i<lib.Len();i++ {
			e, _ := lib.Get(i)  // 忽略了错误？？
			fmt.Println(i+1, ":", e.Name, e.Artist, e.Source, e.Type)
		}
	case "add":
		if len(tokens) == 6 {  //
			id++
			lib.Add(&mlib.MusicEntry{strconv.Itoa(id), tokens[2], tokens[3], tokens[4], tokens[5]})
		} else {
			fmt.Println("USAGE: lib add <name><artist><source><type>")
		}
	case "remove":
		if len(tokens) == 3 {
			lib.Remove(tokens[2])
		} else {
			fmt.Println("USAGE: lib remove <name>")
		}
	default:
		fmt.Println("Unrecognized lib command:", tokens[1])
	}
}
//tokens 代表什么？？

func handlePlayCommands(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("USAGE: play <name>")
		return
	}
	e := lib.Find(tokens[1])
	if e == nil {
		fmt.Println("The music", tokens[1], "does not exist.")
		return
	}
	mp.Play(e.Source, e.Type)
}

func main() {
	fmt.Println(`
		Enter following commands to control the player:
		lib list -- View the existing music lib
		lib add <name><artist><source><type> -- Add a music to the music lib
		lib remove <name> -- Remove the specified music from the lib
		play <name> -- Play the specified music
	`)

	lib = mlib.NewMusicManager()
	r := bufio.NewReader(os.Stdin)  // 带着er的往往都是接口
	for {
		fmt.Print("Enter command->: ")
		rawline, _, _ := r.ReadLine()
		line := string(rawline)
		if line == "q" || line == "e" {
		 	break
		}
		tokens := strings.Split(line, " ")  //列举出tokens的可能取值

		if tokens[0] == "lib" {
			handleLibCommands(tokens)
		} else if tokens[0] == "play" {
			handlePlayCommands(tokens)
		} else {
			fmt.Println("Unrecognized command:", tokens[0])
		}
	}
}
