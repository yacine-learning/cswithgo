package main

import "fmt"

func main() {
	try("Back returns to the previous page", func() string {
		var b Browser
		b.Visit("home")
		b.Visit("docs")
		if !b.Back() || b.Current != "home" {
			return fmt.Sprintf("after Visit home, Visit docs, Back: current is %q, expected \"home\"", b.Current)
		}
		return ""
	})
	try("Back twice, then Forward", func() string {
		var b Browser
		b.Visit("home")
		b.Visit("docs")
		b.Visit("api")
		b.Back()
		b.Back()
		if !b.Forward() || b.Current != "docs" {
			return fmt.Sprintf("home, docs, api, Back, Back, Forward: current is %q, expected \"docs\"", b.Current)
		}
		if !b.Forward() || b.Current != "api" {
			return fmt.Sprintf("a second Forward: current is %q, expected \"api\"", b.Current)
		}
		return ""
	})
	try("nowhere to go back to", func() string {
		var b Browser
		b.Visit("home")
		if b.Back() {
			return "Back with no previous page returned true"
		}
		if b.Current != "home" {
			return fmt.Sprintf("a refused Back changed the current page to %q", b.Current)
		}
		return ""
	})
	try("nowhere to go forward to", func() string {
		var b Browser
		b.Visit("home")
		b.Visit("docs")
		if b.Forward() {
			return "Forward with nothing ahead returned true"
		}
		return ""
	})
	try("visiting a page clears what was ahead", func() string {
		var b Browser
		b.Visit("home")
		b.Visit("docs")
		b.Back()
		b.Visit("blog")
		if b.Forward() {
			return fmt.Sprintf("Forward after Back then Visit blog went to %q", b.Current)
		}
		if !b.Back() || b.Current != "home" {
			return fmt.Sprintf("Back from blog: current is %q, expected \"home\"", b.Current)
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
