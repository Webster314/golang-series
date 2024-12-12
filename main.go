package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"reflect"
	"strconv"
	"sync"
	"time"

	"rsc.io/quote"
)

type Result struct {
	Num, Ans int
}

var wg sync.WaitGroup

func download_wg(url string) {
	fmt.Println("Start to download", url)
	time.Sleep(time.Second)
	wg.Done()
}

type Person interface {
	getName() string
}

type Student struct {
	name string
	age  int
}

func (stu *Student) getName() string {
	return stu.name
}

type Worker struct {
	name string
}

func (w *Worker) getName() string {
	return w.name
}

func (stu *Student) hello(person string) string {
	return fmt.Sprintf("hello %s, I'm %s", person, stu.name)
}
func get(index int) (ret int) {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println(r)
			ret = -1
		}
	}()
	arr := [3]int{2, 3, 4}
	return arr[index]
}

func hello(name string) error {
	if len(name) == 0 {
		return errors.New("error: name is nil")
	}
	fmt.Println("hello, ", name)
	return nil
}

func main() {
	str1 := "Golang"
	var str2 = "Go语言"
	fmt.Println(str1, str2)
	fmt.Println(reflect.TypeOf((str2[2])).Kind())
	fmt.Println(str1[1], string(str1[1]))
	fmt.Printf("%d %c\n", str1[0], str1[0])
	fmt.Println("len(str2):", len(str2))

	runeArr := []rune(str2)
	fmt.Println(reflect.TypeOf((runeArr[2])).Kind())
	fmt.Println(runeArr[2], string(runeArr[2]))
	fmt.Println("len(runeArr):", len(runeArr))

	var arr = []int{1}
	arr[0] = 2

	slice1 := make([]float32, 3)
	slice2 := make([]float32, 3, 5)
	for i := 0; i < len(slice2); i++ {
		println(slice2[i])
	}
	println(len(slice1), cap(slice2))
	slice2 = append(slice2, 1, 2, 3, 4)
	for i := 0; i < len(slice2); i++ {
		println(slice2[i])
	}
	println(len(slice2), cap(slice2))
	sub1 := slice2[3:]
	sub2 := slice2[:3]
	combined := append(sub1, sub2...)
	println(cap(combined))

	m1 := make(map[string]int)
	m2 := map[string]string{
		"Sam":   "Male",
		"Alice": "Female",
	}
	m2["Tom"] = "Male"
	m1["Tom"] = 18
	for key, value := range m2 {
		fmt.Println(key, value)
	}

	str := "Golang"
	var p *string = &str
	*p = "Hello"
	fmt.Println(str)

	// if for switch
	if age := 18; age < 18 {
		fmt.Printf("Kid")
	} else {
		fmt.Printf("Adult")
	}

	type Gender int8
	const (
		MALE   Gender = 1
		FEMALE Gender = 2
	)

	gender := MALE

	switch gender {
	case FEMALE:
		fmt.Println("female")
		// fallthrough
	case MALE:
		fmt.Println("male")
	default:
		fmt.Println("unknown")
	}

	nums := []int{10, 20, 30, 40}
	for i, num := range nums {
		fmt.Println(i, num)
	}

	_, err := os.Open("nonexist.file")
	if err != nil {
		fmt.Println(err)
	}

	// error panic
	err = hello("")
	fmt.Println(err)
	fmt.Println(get(5))
	fmt.Println("finished")

	// structure method interface
	stu := &Student{
		name: "Tom",
		age:  12,
	}
	msg := stu.hello("Tommy")
	fmt.Println(msg)

	var stu1 *Student = new(Student)
	fmt.Println(stu1.hello(stu.name))

	var person Person = &Student{
		name: "Tom",
		age:  18,
	}
	fmt.Println(person.getName())

	var _ Person = (*Student)(nil)
	var _ Person = (*Worker)(nil)

	m := make(map[string]interface{})
	m["name"] = "Tom"
	m["age"] = 1
	m["scores"] = [3]int{98, 99, 85}
	fmt.Println(m)

	// sync goroutine
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go download_wg("a.com/" + strconv.Itoa(i+'0'))
	}
	wg.Wait()
	fmt.Println("Finished")

	for i := 0; i < 3; i++ {
		go download_chan("a.com/" + strconv.Itoa(i+'0'))
	}
	for i := 0; i < 3; i++ {
		msg = <-ch
		fmt.Println(msg)
	}
	fmt.Println("Done")

	//
	fmt.Println(quote.Go())

	c := [5]int{}
	fmt.Println(c)

	config := &tls.Config{
		InsecureSkipVerify: true,
	}
	conn, _ := tls.Dial("tcp", "localhost:1234", config)
	defer conn.Close()
	client := rpc.NewClient(conn)
	var res Result
	if err = client.Call("Cal.Square", 12, &res); err != nil {
		log.Fatal("Call failed")
	}
	log.Printf("%d^2 == %d", res.Num, res.Ans)

	asyncCall := client.Go("Cal.Square", 20, &res, nil)
	log.Printf("%d^2 == %d", res.Num, res.Ans)

	<-asyncCall.Done
	log.Printf("%d^2 == %d", res.Num, res.Ans)
}
