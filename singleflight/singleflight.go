package main

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/sync/singleflight"
)

func singleFlight1()  {
	var singleSetCache singleflight.Group

	getAndSetCache:=func (requestID int,cacheKey string) (string,bool,error) {
		log.Printf("request %v start to get and set cache...",requestID)

		//do的入参key，可以直接使用缓存的key，这样同一个缓存，只有一个协程会去读DB
		value,err, shared :=singleSetCache.Do(cacheKey, func() (ret interface{}, err error) {
			log.Printf("request %v is setting cache...",requestID)
			time.Sleep(3*time.Second)
			log.Printf("request %v set cache success!",requestID)
			return "VALUE",nil
		})
		if err != nil {
			log.Println(err)
			return "",shared,err
		}
		return value.(string),shared,nil
	}

	cacheKey:="cacheKey"
	for i:=1;i<10;i++{//模拟多个协程同时请求
		go func(requestID int) {
			value,shared,_:=getAndSetCache(requestID,cacheKey)
			log.Printf("request %v get value: %v shared:%v",requestID,value,shared)
		}(i)
	}
	time.Sleep(20*time.Second)
}

func singleFlight2()  {
	g := singleflight.Group{}
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			val, err, shared := g.Do("a", a)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("index: %d, val: %d, shared: %v\n", j, val, shared)
		}(i)
	}

	wg.Wait()
}

var (
	count = int64(0)
)

// 模拟接口方法
func a() (interface{}, error) {
	time.Sleep(time.Millisecond * 15000)
	return atomic.AddInt64(&count, 1), nil
}
//
//// 部分输出，shared表示是否共享了其他请求的返回结果
//index: 2, val: 1, shared: false
//index: 71, val: 1, shared: true
//index: 69, val: 1, shared: true
//index: 73, val: 1, shared: true
//index: 8, val: 1, shared: true
//index: 24, val: 1, shared: true


func main() {
	singleFlight1()
}
