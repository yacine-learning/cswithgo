package main

import "fmt"

func main() {
	var v Vec
	for i := 1; i <= 10; i++ {
		v.Push(i * 10)
		fmt.Printf("after Push(%d): len %d, cap %d, copies %d\n", i*10, v.Len(), v.Cap(), v.Copies)
	}
}
