package main

func main() {
	solve(123)
	solve("hello")
	solve(true)
	solve(make(chan int))
	solve(make(chan string))
	solve(make(chan bool))
	solve(12.34)

}

func solve(perem interface{}) {
	switch perem.(type) {
	case int:
		println("integer")
	case string:
		println("string")
	case bool:
		println("boolean")
	case chan int:
		println("channel of int")
	case chan string:
		println("channel of string")
	case chan bool:
		println("channel of bool")
	default:
		println("unknown type")
	}
}
