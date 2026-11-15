package main

import (
	"fmt"
	"strings"
)

var files = []string{
	"app/README.md 2",
	"app/src/main.go 12",
	"app/src/api/api.go 20",
	"app/config/app.yml 1",
}

func expect(lines []string, depth int, want string) func() string {
	return func() string {
		out, err := Du(lines, depth)
		if err != nil {
			return "unexpected error: " + err.Error()
		}
		got := strings.Join(out, " | ")
		if got != want {
			return fmt.Sprintf("Du(depth %d) = %q\n      expected %q", depth, got, want)
		}
		return ""
	}
}

func main() {
	try("depth 0: the top directory only", expect(files, 0, "35\tapp"))
	try("depth 1: post-order, names sorted", expect(files, 1, "1\tapp/config | 32\tapp/src | 35\tapp"))
	try("depth 2: deeper directories first", expect(files, 2, "1\tapp/config | 20\tapp/src/api | 32\tapp/src | 35\tapp"))
	try("several top-level directories, sorted", expect([]string{"web/a 5", "api/b 3", "web/x/c 1"}, 1,
		"3\tapi | 1\tweb/x | 6\tweb"))
	try("a top-level file is not listed", expect([]string{"notes.txt 50", "app/a 1"}, 0, "1\tapp"))
	try("a directory with no files of its own", expect([]string{"a/b/c/d 7"}, 3, "7\ta/b/c | 7\ta/b | 7\ta"))
	try("no files", expect(nil, 2, ""))
	try("a line without a size", func() string {
		_, err := Du([]string{"app/a 1", "app/b"}, 1)
		if err == nil || err.Error() != `line 2: cannot read "app/b"` {
			return fmt.Sprintf("error %v, expected line 2: cannot read \"app/b\"", err)
		}
		return ""
	})
	try("a size that is not a number", func() string {
		_, err := Du([]string{"app/a big"}, 1)
		if err == nil || err.Error() != `line 1: cannot read "app/a big"` {
			return fmt.Sprintf("error %v, expected line 1: cannot read \"app/a big\"", err)
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
