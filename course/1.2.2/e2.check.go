package main

import "fmt"

func main() {
	try("put then get", func() string {
		m := NewHashMap()
		m.Put("web-1", 80)
		if v, ok := m.Get("web-1"); !ok || v != 80 {
			return fmt.Sprintf("Get(\"web-1\") = %d, %v after Put(\"web-1\", 80)", v, ok)
		}
		return ""
	})
	try("a missing key", func() string {
		m := NewHashMap()
		m.Put("web-1", 80)
		if v, ok := m.Get("db-1"); ok {
			return fmt.Sprintf("Get(\"db-1\") = %d, true; it was never put", v)
		}
		return ""
	})
	try("putting a key again replaces its value", func() string {
		m := NewHashMap()
		m.Put("web-1", 80)
		m.Put("web-1", 8080)
		if v, _ := m.Get("web-1"); v != 8080 {
			return fmt.Sprintf("Get(\"web-1\") = %d after Put 80 then 8080", v)
		}
		if m.Len() != 1 {
			return fmt.Sprintf("Len() = %d after putting the same key twice, expected 1", m.Len())
		}
		return ""
	})
	try("keys in the same bucket", func() string {
		m := NewHashMap()
		// with 4 buckets these keys collide: find two that do
		var a, b string
		for i := 0; a == "" || b == ""; i++ {
			k := fmt.Sprintf("host-%d", i)
			if m.bucketOf(k) == 0 {
				if a == "" {
					a = k
				} else {
					b = k
				}
			}
		}
		m.Put(a, 1)
		m.Put(b, 2)
		va, _ := m.Get(a)
		vb, _ := m.Get(b)
		if va != 1 || vb != 2 {
			return fmt.Sprintf("%q and %q share bucket 0: Get gave %d and %d, expected 1 and 2", a, b, va, vb)
		}
		return ""
	})
	try("1 000 keys through several growths", func() string {
		m := NewHashMap()
		for i := 0; i < 1000; i++ {
			m.Put(fmt.Sprintf("k%d", i), i)
		}
		if m.Len() != 1000 {
			return fmt.Sprintf("Len() = %d, expected 1000", m.Len())
		}
		for i := 0; i < 1000; i++ {
			if v, ok := m.Get(fmt.Sprintf("k%d", i)); !ok || v != i {
				return fmt.Sprintf("Get(\"k%d\") = %d, %v", i, v, ok)
			}
		}
		if len(m.buckets) < 500/2 {
			return fmt.Sprintf("only %d buckets for 1 000 keys: is count updated?", len(m.buckets))
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
