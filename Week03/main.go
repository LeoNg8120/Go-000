package main

import (
	"context"
	"errors"
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
	g,ctx := errgroup.WithContext(context.Background())
	ip := "127.0.0.1"

	signalChan:=make(chan os.Signal,1)
	signal.Notify(signalChan,syscall.SIGINT,syscall.SIGTERM)

	//1. 建议使用 errgroup 控制整个流程
	g.Go(func() error {
		select {
		//2. 除了这个信号量之外，如果其中一个 server 启动的时候，报错也需要考虑。
		case <-signalChan:
			log.Println("recv SIGTERM,service will exit...")
			// inject error
			g.Go(func() error {
				log.Println("start inject error")
				time.Sleep(time.Second)
				log.Println("inject finish")
				return errors.New("inject error")
			})
		case <-ctx.Done():
			log.Println("signal ctx done")
			return ctx.Err()
		}
		return nil
	})

	http.HandleFunc("/",Handler)
	for port:=1000;port<1050;port++ {
		tmp:=port
		g.Go(func() error {
			fmt.Printf("HttpService start,port is %d\n",tmp)
			srv := &http.Server{Addr:ip+":"+strconv.Itoa(tmp)}

			go func(svr *http.Server) {
				<-ctx.Done()
				log.Println("HttpService get ctx done")
				err:=svr.Shutdown(context.TODO())
				if err != nil {
					log.Println("svr shutdown err")
				}
				log.Println("HttpService Shutdown success")
			}(srv)

			if err:=srv.ListenAndServe();err!= nil {
				fmt.Printf("HttpService(%v) ListenAndServe error:%v\n",srv.Addr,err.Error())
			}
			time.Sleep(1)
			return nil
		})
	}

	if err:=g.Wait();err == nil {
		fmt.Println("HttpService g.wait All Successfully exit.")
	}else{
		fmt.Println("HttpService g.wait error:",err)
	}
	time.Sleep(5*time.Second)
}

