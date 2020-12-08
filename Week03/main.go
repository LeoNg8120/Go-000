package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)

/*
	基于 errgroup 实现多个 http server 的启动和关闭 ，
	以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出。
 */



func main() {
	g := new(errgroup.Group)

	err:=HttpRegister()
	if err != nil {
		panic("HttpRegister err:"+err.Error())
	}
	signalChan:=make(chan os.Signal,1)
	signal.Notify(signalChan,syscall.SIGINT,syscall.SIGTERM)

	for port:=1000;port<1050;port++ {
		tmp:=port
		g.Go(func()error {
			fmt.Printf("httpservice start,port is %d\n",tmp)
			err=HttpRun(tmp)
			if err != nil {
				fmt.Printf("httpService run port:%d error,err:%s",tmp,err.Error())
				return err
			}
			return nil
		})
	}

	if err:=g.Wait();err == nil {
		fmt.Println("HttpService Successfully exit.")
	}else{
		fmt.Println("HttpService error:",err)
	}
}

