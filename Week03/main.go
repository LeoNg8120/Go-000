package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"strconv"
)

/*
	基于 errgroup 实现一个 http server 的启动和关闭 ，
	以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出。
 */


func main() {
	g := new(errgroup.Group)

	for port:=1000;port<1050;port++ {
		tmp:=port
		g.Go(func() error {
			httpService,err:=NewHttpService("Service"+strconv.Itoa(tmp),tmp)
			if err != nil {
				return err
			}
			httpService.Run()
			return nil
		})

	}

	if err:=g.Wait();err == nil {
		fmt.Println("HttpService Successfully exit.")
	}else{
		fmt.Println("HttpService error:",err)
	}
}

