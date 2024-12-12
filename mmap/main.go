package main

import (
	"fmt"

	"golang.org/x/exp/mmap"
)

func main(){
    // f, _ := os.OpenFile("tmp.txt", os.O_CREATE | os.O_RDWR, 0644);
    at, _ := mmap.Open("tmp.txt");
    buff := make([]byte, 2)
    // _, _ = f.ReadAt(buff, 4)
    _, _ = at.ReadAt(buff, 4)
    _ = at.Close()
    fmt.Println(string(buff))
}
