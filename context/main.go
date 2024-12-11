package main

import (
	"context"
	"fmt"
	"time"
)

func req(ctx context.Context, msg string){
    for{
        select{
            case <-ctx.Done():
                fmt.Println("stop")
                return
            default:
                fmt.Println(msg)
                time.Sleep(time.Second)
        }
    }
}

func main(){
    ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))
    go req(ctx, "work1")
    go req(ctx, "work2")
    time.Sleep(3 * time.Second);
    cancel()
    time.Sleep(3 * time.Second);
}

