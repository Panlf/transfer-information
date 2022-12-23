package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-vgo/robotgo/clipboard"
	hook "github.com/robotn/gohook"
)

func main() {
	logFile, err := os.OpenFile("monitor.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Ldate)

	// 异步操作
	hook.Register(hook.KeyDown, []string{"ctrl", "c"}, func(e hook.Event) {
		content, _ := clipboard.ReadAll()
		content = handleConent(content)
		clipboard.WriteAll(content)
	})

	hook.Register(hook.KeyDown, []string{"esc"}, func(e hook.Event) {
		fmt.Println("esc")
		hook.End()
	})

	s := hook.Start()
	<-hook.Process(s)

	fmt.Println("the program is over!")
}

func handleConent(content string) string {
	log.Printf("处理前的字符串 : %s", content)
	//content = strings.Trim(content, " ")
	content = replaceWrapLeft(content)
	content = replaceWrapRight(content)
	log.Printf("处理后的字符串 : %s", content)
	return content
}

func replaceWrapLeft(content string) string {
	if content == "" {
		return content
	}
	num := 0
	len := len(content)
	for i := 0; i < len; i++ {
		if content[i] == '\r' || content[i] == '\t' || content[i] == '\n' || content[i] == ' ' {
			num++
		} else {
			break
		}
	}
	return content[num:len]
}

func replaceWrapRight(content string) string {
	if content == "" {
		return content
	}
	num := 0
	len := len(content)
	for i := len - 1; i > 0; i-- {
		if content[i] == '\r' || content[i] == '\t' || content[i] == '\n' || content[i] == ' ' {
			num++
		} else {
			break
		}
	}
	return content[0 : len-num]
}
