package main

import "fmt"

func main() {
	ok := Graph{"web": {"api"}, "api": {"db", "cache"}, "worker": {"db"}, "db": nil, "cache": nil}
	bad := Graph{"web": {"api"}, "api": {"auth", "db"}, "auth": {"users"}, "users": {"api"}, "db": nil}
	fmt.Println("no cycle:  ", FindCycle(ok))
	fmt.Println("with cycle:", FindCycle(bad))
}
