package main

import "fmt"

func main() {
	for out := range FizzBuzz(100) {
		fmt.Println(out)
	}
}

func FizzBuzz(amount int) <-chan string {

	out := make(chan string, amount)

	go func() {
		for i := 1; i <= amount; i++ {
			result := ""
			if i%3 == 0 { result += "Fizz" }
			if i%5 == 0 { result += "Buzz" }
			if result == "" { result = fmt.Sprintf("%v", i) }
			out <- result
		}
		close(out)
	}()

	return out
}