package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/go-vgo/robotgo"
	_ "github.com/robotn/gohook"
	hook "github.com/robotn/gohook"
)

func main() {

	/*clipboard.WriteAll("你好啊")
	text, _ := clipboard.ReadAll()
	fmt.Println(text)*/

	/*
	var k = robotgo.AddEvents("ctrl","c")
	if k  {
		fmt.Println("you press... ", "ctrl+c")
	}*/

	for {
		robotgo.EventHook(hook.KeyDown, []string{"ctrl", "c"}, func(e hook.Event) {
			//fmt.Println("ctrl-c")

			content, err := clipboard.ReadAll()

			if err == nil {
				fmt.Println(content)
			}

			robotgo.EventEnd()
		})

		s := robotgo.EventStart()
		<-robotgo.EventProcess(s)
	}
}