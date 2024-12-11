package main

import (
	"fmt"
	"log"
	"os"
    
	"google.golang.org/protobuf/proto"
)

func main(){
    test := &Student{
        Name : "wbx",
        Male : true,
        Score : []int32{80, 90, 70},
    }

    var data []byte;
    // data, err := proto.Marshal(test);
    // if err != nil{
    //     log.Fatal("marshalling error", err);
    // }
    //
    // file ,err := os.OpenFile("struct", os.O_RDWR, 0);
    // n, err := file.Write(data);
    // if err != nil{
    //     log.Fatal("write error", err);
    // }
    // fmt.Println(n, "bytes write");

    data, err := os.ReadFile("struct");
    if err != nil{
        log.Fatal("read error", err);
    }
    fmt.Println(len(data), "bytes read");

    newTest := &Student{}
    err = proto.Unmarshal(data, newTest);
    if err != nil{
        log.Fatal("unmarshalling error", err);
    }
    if test.GetName() != newTest.GetName(){
        log.Fatalf("data mistach! %q != %q", test.GetName(), newTest.GetName());
    }
}
