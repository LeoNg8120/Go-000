package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type HttpService struct {
	name string
	port int
}

func NewHttpService(name string,port int) (*HttpService,error) {
	return &HttpService{name: name,port: port},nil
}

func (h *HttpService)Run()error{
	fmt.Printf("HttpService:%s Start Listen Port %d\n",h.name,h.port)
	http.HandleFunc("/",h.handler)
	err := http.ListenAndServe("127.0.0.1:"+strconv.Itoa(h.port), nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
		return err
	}
	fmt.Printf("HttpService:%s end...",h.name)
	return nil
}

func (h *HttpService)handler(writer http.ResponseWriter, request *http.Request)  {
	t := time.Now()
	timeStr := fmt.Sprintf("Hello this is {\"time\": \"%s\"} from "+h.name, t)
	writer.Write([]byte(timeStr))
}




