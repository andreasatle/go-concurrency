package main

import "fmt"

type ByteCounter int

// Write returns the accumulated byte-count of all calls
func (bc *ByteCounter) Write(p []byte) (int, error) {
	*bc += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var b ByteCounter

	str1 := "Hello, World!"
	fmt.Fprintf(&b, str1)
	fmt.Printf("String: \"%s\"\nlength: %d\nByteCount: %d\n\n", str1, len(str1), b)

	str2 := "Goodbye, World!"
	fmt.Fprintf(&b, str2)
	fmt.Printf("String: \"%s\"\nlength: %d\nByteCount: %d\n\n", str2, len(str2), b)

	str3 := "Foo, bar!"
	fmt.Fprintf(&b, str3)
	fmt.Printf("String: \"%s\"\nlength: %d\nByteCount: %d\n\n", str3, len(str3), b)
}
