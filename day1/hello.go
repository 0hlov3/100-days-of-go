package main

import "fmt"

func Hello() string {
	return "Hello, World"
}

func HelloName(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println("Hello, world")
	fmt.Println(Hello())
	fmt.Println(HelloName("0hlov3"))
}
