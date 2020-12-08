package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func HttpRun(port int)error{
	err := http.ListenAndServe("127.0.0.1:"+strconv.Itoa(port), nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
		return err
	}
	return nil
}

func HttpRegister()error{
	http.HandleFunc("/",handler)
	return nil
}
func handler(writer http.ResponseWriter, request *http.Request)  {
	t := time.Now()
	timeStr := fmt.Sprintf("Hello this time is %s\n",t.String())
	writer.Write([]byte(timeStr))
}

