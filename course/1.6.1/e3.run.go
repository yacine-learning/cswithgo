package main

import "fmt"

func main() {
	booked := []Window{{120, 180}, {0, 30}, {150, 200}, {30, 45}, {300, 360}, {170, 190}}
	fmt.Println("booked:", booked)
	fmt.Println("merged:", MergeWindows(booked))
}
