package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

/*
	基于 errgroup 实现多个 http server 的启动和关闭 ，
	以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出。
 */

func Handler(writer http.ResponseWriter, request *http.Request)  {
	t := time.Now()
	timeStr := fmt.Sprintf("Hello this time is %s\n",t.String())
	writer.Write([]byte(timeStr))
}

func main() {
	g := new(errgroup.Group)
	ip := "127.0.0.1"

	signalChan:=make(chan os.Signal,1)
	srvChan:=make(chan *http.Server,50)

	signal.Notify(signalChan,syscall.SIGINT,syscall.SIGTERM)

	go func() {
		select {
		case <-signalChan:
			log.Println("recv SIGTERM,service will exit...")
			for server := range srvChan {
				if err:=server.Shutdown(context.Background());err!=nil {
					fmt.Printf("HttpService(%v) Shutdown err:%v\n",server.Addr,err)
				}
			}
		}
	}()

	http.HandleFunc("/",Handler)
	for port:=1000;port<1050;port++ {
		tmp:=port
		g.Go(func() error {
			fmt.Printf("HttpService start,port is %d\n",tmp)
			srv := &http.Server{Addr:ip+":"+strconv.Itoa(tmp)}
			srvChan<-srv
			if err:=srv.ListenAndServe();err!= nil {
				fmt.Printf("HttpService(%v) ListenAndServe error:%v\n",srv.Addr,err.Error())
			}
			return nil
		})
	}

	if err:=g.Wait();err == nil {
		fmt.Println("HttpService g.wait All Successfully exit.")
	}else{
		fmt.Println("HttpService g.wait error:",err)
	}
}

