package main

import "fmt"

// code for question 1
func exp(x float64, n int) float64 {
	var sum, term float64 = 1, 1
	for i := 1; i <= n; i++ {
		term *= x / float64(i)
		sum += term
	}
	return sum
}

func question1() {
	var x float64 = 2.0
	var n int = 10

	fmt.Println(exp(x, n))
}

// code for question 2

func getLastDigit(n int64) int {
	if n == 0 {
		return 0
	}
	value1, value2, value3 := 1, 1, 0
	for _, rec := range fmt.Sprintf("%b", n)[1:] {
		c := (value2 * value2) % 10
		value1, value2, value3 = (value1*value1+c)%10, ((value1+value3)*value2)%10, (c+value3*value3)%10
		if rec == '1' {
			value1, value2, value3 = (value1+value2)%10, value1, value2
		}
	}
	return value2
}

func question2() {
	var n int64 = 440
	fmt.Println("last digit of nth fibonacci number - ", getLastDigit(n))
}

func main() {
	question1()
	question2()
}
