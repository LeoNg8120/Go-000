package main

//基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"
	//"go-project-layout/server/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"golang.org/x/sync/errgroup"
)


func main() {
	ip := "127.0.0.1"
	g, ctx := errgroup.WithContext(context.Background())
	svr := http.Server{Addr:ip+":"+strconv.Itoa(1)}

	// http server
	g.Go(func() error {
		log.Println("http start listen")
		go func() {
			<-ctx.Done()
			log.Println("http server get ctx done")
			err:=svr.Shutdown(context.TODO())
			if err != nil {
				log.Println("svr shutdown err")
			}
			log.Println("http server Shutdown success")
		}()
		return svr.ListenAndServe()
	})

	// signal
	g.Go(func() error {
		exitSignals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT} // SIGTERM is POSIX specific
		sig := make(chan os.Signal, len(exitSignals))
		signal.Notify(sig, exitSignals...)
		for {
			log.Println("signal register")
			select {
			case <-ctx.Done():
				log.Println("signal ctx done")
				return ctx.Err()
			case <-sig:
				// do something
				return nil
			}
		}
	})
	time.Sleep(3*time.Second)
	// inject error
	g.Go(func() error {
		log.Println("start inject error")
		time.Sleep(time.Second)
		log.Println("inject finish")
		return errors.New("inject error")
	})

	err := g.Wait() // first error return
	log.Println("first error return from errGroup:",err)
	time.Sleep(3*time.Second)
}