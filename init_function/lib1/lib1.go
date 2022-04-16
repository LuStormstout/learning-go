package lib1

import "fmt"

// TestLib1 当前 lib1 包提供的 API
func TestLib1() {
	fmt.Println("TestLib1() ...")
}

func init() {
	fmt.Println("lib1 init() ...")
}
