package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello() string {
	return "Hello, World"
}

func HelloName(name string) string {
	return "Hello, " + name
}

func HelloNameConst(name string) string {
	return englishHelloPrefix + name
}

func HelloNameFallBack(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println("Hello, world")
	fmt.Println(Hello())
	fmt.Println(HelloName("0hlov3"))
	fmt.Println(HelloNameConst("0hlov3"))
	fmt.Println(HelloNameFallBack("0hlov3"))
}
