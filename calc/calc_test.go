package calc

import (
	"fmt"
	"os"
	"testing"
)

func setup(){
    fmt.Println("before all tests")
}

func teardown(){
    fmt.Println("after all tests")
}

func TestAdd(t *testing.T) {
    if sum := Add(1,2); sum != 3 {
		t.Error("add(1, 2) should be equal to 3")
	}
}

func TestMul(t *testing.T) {
    cases := []struct{
        Name string
        A, B, Expected int
    }{
        {"pos", 2, 3, 6},
        {"neg", -2, 3, -6},
        {"zero", 0, 5, 0},
        // {"wrong", 1, 2, 0}, // wrong case
    }
    for _, c := range cases{
        t.Run(c.Name, func(t * testing.T){
            if rc := Mul(c.A, c.B); rc != c.Expected{
                t.Fatalf("Mul(%d, %d) should be equal to %d, not %d",
                            c.A, c.B, c.Expected, rc)
            }
        }) 
    }
}

func TestDiv(t * testing.T){
    if quo, rem := Div(7, 2); quo != 3 || rem != 1 {
		t.Error("div(7, 2) shoul be equal to 3, 1")
	}
}

func BenchmarkHello(b * testing.B){
    for i := 0; i < b.N; i++{
        fmt.Sprintf("abd%d", 1);
    }
}

func TestMain(m * testing.M){
    setup()
    rc := m.Run()
    teardown()
    os.Exit(rc);
}
