package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
	"william/Go-000/Week06/rolling"
)

func main() {
	wg := sync.WaitGroup{}
	slide := rolling.NewSlideWindow(10)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	wg.Add(3)
	go func() {
		for i := 0; i < 10000; i++ {
			slide.Success.Increment(r.Float64())
			time.Sleep(10 * time.Millisecond)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 10000; i++ {
			slide.Failed.Increment(r.Float64())
			time.Sleep(15 * time.Millisecond)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 10000; i++ {
			slide.Timeout.Increment(r.Float64())
			time.Sleep(20 * time.Millisecond)
		}
		wg.Done()
	}()

	for {
		now := time.Now()
		fmt.Printf("success max:%f sum:%f avg:%f\n", slide.Success.Max(now), slide.Success.Sum(now), slide.Success.Avg(now))
		fmt.Printf("failed max:%f sum:%f avg:%f\n", slide.Failed.Max(now), slide.Failed.Sum(now), slide.Failed.Avg(now))
		fmt.Printf("timeout max:%f sum:%f avg:%f\n\n", slide.Timeout.Max(now), slide.Timeout.Sum(now), slide.Timeout.Avg(now))
		time.Sleep(100 * time.Millisecond)
	}
	wg.Wait()
	fmt.Println("end")
}
