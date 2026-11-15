package main

import "fmt"

func main() {
	fine := Graph{"puppy": {"young", "dog"}, "dog": {"animal", "domestic"}, "young": nil, "animal": nil, "domestic": nil}
	circular := Graph{"big": {"large"}, "large": {"great"}, "great": {"big", "size"}, "size": nil, "huge": {"big"}}
	fmt.Println("no cycle:  ", FindCycle(fine))
	fmt.Println("with cycle:", FindCycle(circular))
}
