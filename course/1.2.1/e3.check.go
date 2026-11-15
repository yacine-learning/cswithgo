package main

import "fmt"

func main() {
	try("two results from the same base keep their own flag", func() string {
		base := make([]string, 0, 10)
		base = append(base, "deploy", "--env=prod")
		a := WithFlag(base, "--canary")
		b := WithFlag(base, "--all-regions")
		if len(a) != 3 || a[2] != "--canary" {
			return fmt.Sprintf("first result is %v after building the second one; expected [deploy --env=prod --canary]", a)
		}
		if len(b) != 3 || b[2] != "--all-regions" {
			return fmt.Sprintf("second result is %v, expected [deploy --env=prod --all-regions]", b)
		}
		return ""
	})
	try("base is unchanged", func() string {
		base := []string{"deploy", "--env=prod"}
		WithFlag(base, "--canary")
		if len(base) != 2 || base[0] != "deploy" || base[1] != "--env=prod" {
			return fmt.Sprintf("base became %v", base)
		}
		return ""
	})
	try("changing a result does not change base", func() string {
		base := make([]string, 2, 10)
		base[0], base[1] = "deploy", "--env=prod"
		r := WithFlag(base, "--canary")
		r[0] = "rollback"
		if base[0] != "deploy" {
			return "setting result[0] also changed base[0]: they share one array"
		}
		return ""
	})
	try("an empty base", func() string {
		r := WithFlag(nil, "--help")
		if len(r) != 1 || r[0] != "--help" {
			return fmt.Sprintf("WithFlag(nil, \"--help\") = %v", r)
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
