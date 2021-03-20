package main_test

import (
	"sync"
	"sync/atomic"
	"testing"
)

type config struct {
	a []int
}

func (c *config)T()  {
	
}

//go test atomic_test.go -v -run TestMutex
func TestMutex(t *testing.T)  {
	var mutex sync.Mutex
	cfg := &config{}
	go func() {
		i:=0
		for  {
			i++
			mutex.Lock()
			cfg.a=[]int{i,i+1,i+2,i+3}
			mutex.Unlock()
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 100 ; i++ {
				mutex.Lock()
				t.Logf("%v\n",cfg)
				mutex.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestRwMutex(t *testing.T)  {
	var rwmutex sync.RWMutex
	cfg := &config{}
	go func() {
		i:=0
		for  {
			i++
			rwmutex.Lock()
			cfg.a=[]int{i,i+1,i+2,i+3}
			rwmutex.Unlock()
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 100 ; i++ {
				rwmutex.RLock()
				t.Logf("%v\n",cfg)
				rwmutex.RUnlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestAtomic(t *testing.T)  {
	var v atomic.Value
	v.Store(&config{})
	go func() {
		i:=0
		for  {
			i++
			cfg:=&config{a: []int{i,i+1,i+2,i+3}}
			v.Store(cfg)
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 100 ; i++ {
				cfg:=v.Load().(*config)
				cfg.T()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}


func BenchmarkMutex(b *testing.B)  {
	var mutex sync.Mutex
	cfg := &config{}
	b.ResetTimer()
	go func() {
		i:=0
		for  {
			i++
			mutex.Lock()
			cfg.a=[]int{i,i+1,i+2,i+3}
			mutex.Unlock()
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < b.N ; i++ {
				mutex.Lock()
				cfg.T()
				mutex.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	b.StopTimer()
}

func BenchmarkRwMutex(b *testing.B)  {
	var rwmutex sync.RWMutex
	cfg := &config{}
	b.ResetTimer()
	go func() {
		i:=0
		for  {
			i++
			rwmutex.Lock()
			cfg.a=[]int{i,i+1,i+2,i+3}
			rwmutex.Unlock()
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < b.N ; i++ {
				rwmutex.RLock()
				cfg.T()
				rwmutex.RUnlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	b.StopTimer()
}

//go test -bench=.
func BenchmarkAtomic(b *testing.B)  {
	var v atomic.Value
	v.Store(&config{})
	b.ResetTimer()
	go func() {
		i:=0
		for  {
			i++
			cfg:=&config{a: []int{i,i+1,i+2,i+3}}
			v.Store(cfg)
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < b.N ; i++ {
				cfg:=v.Load().(*config)
				cfg.T()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	b.StopTimer()
}


