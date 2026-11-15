package main

import (
	"fmt"
	"strings"
)

var files = []string{
	"compiler/go.mod 2",
	"compiler/lexer/lex.go 12",
	"compiler/lexer/testdata/input.txt 20",
	"compiler/parser/parse.go 1",
}

func expect(lines []string, depth int, want string) func() string {
	return func() string {
		out, err := Sizes(lines, depth)
		if err != nil {
			return "unexpected error: " + err.Error()
		}
		got := strings.Join(out, " | ")
		if got != want {
			return fmt.Sprintf("Sizes(depth %d) = %q\n      expected %q", depth, got, want)
		}
		return ""
	}
}

func main() {
	try("depth 0: the top folder only", expect(files, 0, "35\tcompiler"))
	try("depth 1: post-order, names sorted", expect(files, 1, "32\tcompiler/lexer | 1\tcompiler/parser | 35\tcompiler"))
	try("depth 2: deeper folders first", expect(files, 2, "20\tcompiler/lexer/testdata | 32\tcompiler/lexer | 1\tcompiler/parser | 35\tcompiler"))
	try("several top-level folders, sorted", expect([]string{"web/a 5", "api/b 3", "web/x/c 1"}, 1,
		"3\tapi | 1\tweb/x | 6\tweb"))
	try("a top-level file is not listed", expect([]string{"notes.txt 50", "compiler/a 1"}, 0, "1\tcompiler"))
	try("a folder with no files of its own", expect([]string{"a/b/c/d 7"}, 3, "7\ta/b/c | 7\ta/b | 7\ta"))
	try("no files", expect(nil, 2, ""))
	try("a line without a size", func() string {
		_, err := Sizes([]string{"compiler/a 1", "compiler/b"}, 1)
		if err == nil || err.Error() != `line 2: cannot read "compiler/b"` {
			return fmt.Sprintf("error %v, expected line 2: cannot read \"compiler/b\"", err)
		}
		return ""
	})
	try("a size that is not a number", func() string {
		_, err := Sizes([]string{"compiler/a big"}, 1)
		if err == nil || err.Error() != `line 1: cannot read "compiler/a big"` {
			return fmt.Sprintf("error %v, expected line 1: cannot read \"compiler/a big\"", err)
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
