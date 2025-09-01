package main

import "fmt"

func main() {
	Hello("World")
}

func Hello(name string) string {
	message := fmt.Sprintf("Hello, %v!", name)
	fmt.Println(message)
	return message
}
