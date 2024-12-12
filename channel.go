package main

import (
	"fmt"
	"time"
)

var ch = make(chan string, 10)

func download_chan(url string) {
	fmt.Println("downloading", url)
	time.Sleep(time.Second)
	ch <- url
}
