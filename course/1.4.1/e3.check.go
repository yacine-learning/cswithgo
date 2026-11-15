package main

import "fmt"

func main() {
	g := Graph{
		"web": {"api"}, "api": {"db", "cache", "queue"}, "worker": {"db", "queue"},
		"admin": {"db"}, "db": nil, "cache": nil, "queue": nil,
	}
	try("deploy first (20 calls, always sorted)", func() string {
		for i := 0; i < 20; i++ {
			if got := fmt.Sprint(DeployFirst(g)); got != "[cache db queue]" {
				return "DeployFirst = " + got + ", expected [cache db queue]"
			}
		}
		return ""
	})
	try("safe to remove (20 calls, always sorted)", func() string {
		for i := 0; i < 20; i++ {
			if got := fmt.Sprint(SafeToRemove(g)); got != "[admin web worker]" {
				return "SafeToRemove = " + got + ", expected [admin web worker]"
			}
		}
		return ""
	})
	try("a lone service is both", func() string {
		g := Graph{"solo": nil}
		if fmt.Sprint(DeployFirst(g)) != "[solo]" || fmt.Sprint(SafeToRemove(g)) != "[solo]" {
			return fmt.Sprintf("DeployFirst = %v, SafeToRemove = %v, expected [solo] for both", DeployFirst(g), SafeToRemove(g))
		}
		return ""
	})
	try("a chain a -> b -> c", func() string {
		g := Graph{"a": {"b"}, "b": {"c"}, "c": nil}
		if fmt.Sprint(DeployFirst(g)) != "[c]" || fmt.Sprint(SafeToRemove(g)) != "[a]" {
			return fmt.Sprintf("DeployFirst = %v, SafeToRemove = %v; expected [c] and [a]", DeployFirst(g), SafeToRemove(g))
		}
		return ""
	})
}

func try(name string, f func() string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("FAIL %s: the code panicked: %v\n", name, r)
		}
	}()
	if msg := f(); msg != "" {
		fmt.Printf("FAIL %s: %s\n", name, msg)
	} else {
		fmt.Printf("PASS %s\n", name)
	}
}
