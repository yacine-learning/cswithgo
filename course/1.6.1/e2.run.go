package main

import "fmt"

func main() {
	deploys := []Rec{{3, "api-1"}, {1, "web-1"}, {3, "api-2"}, {2, "db-1"}, {1, "web-2"}, {3, "api-3"}}
	fmt.Println("by priority:", MergeSort(deploys))
}
