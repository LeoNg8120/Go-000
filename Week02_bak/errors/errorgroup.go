package main

import (
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

func main() {
	g := new(errgroup.Group)
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.ws.com/",
		"http://www.somestupidname.com/",
	}

	for _,url := range urls {
		url := url
		g.Go(func() error {
			fmt.Println(url)
			time.Sleep(1*time.Second)
			//resp,err := http.Get(url)
			//if err != nil {
			//	resp.Body.Close()
			//}
			if url=="http://www.google.com/" {
				return errors.New("google")
			}
			if url=="http://www.ws.com/" {
				return errors.New("ws")
			}
			return nil
		})
	}
	if err:=g.Wait();err == nil {
		fmt.Println("Successfully fetched all urls.")
	}else{
		fmt.Println("Get urls error:",err)
	}
}
