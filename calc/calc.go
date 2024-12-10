package calc

func Add(x int, y int) int {
	return x + y
}

func Div(num1 int, num2 int) (quo int, rem int) {
	quo = num1 / num2
	rem = num1 % num2
	return
}
