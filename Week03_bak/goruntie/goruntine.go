package main

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"time"
)

func tryErr()  {
	errs := recover()
	if errs == nil {
		return
	}
	now := time.Now()  //获取当前时间
	pid := os.Getpid() //获取进程ID

	timeStr := now.Format("20060102150405")                                                      //设定时间格式
	fName := fmt.Sprintf("panic-%d-%s-dump.log", pid, timeStr) //保存错误信息文件名:程序名-进程ID-当前时间（年月日时分秒）
	log.Println("dump to file ", fName)

	f, err := os.Create(fName)
	if err != nil {
		return
	}
	defer f.Close()

	_, _ = f.WriteString(fmt.Sprintf("%v\r\n", errs)) //输出panic信息
	_, _ = f.WriteString("========\r\n")

	_, _ = f.WriteString(string(debug.Stack())) //输出堆栈信息
}

func Run()  {
	defer tryErr()
	time.Sleep(3*time.Second)
	panic("Run err...")
}

func main() {
	go Run()

	select {}
}
