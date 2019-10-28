package main

import (
	"fmt"
	"goworkspace/main/stringreverse"
	"goworkspace/testpkg"
)

func main() {
	fmt.Println("Sample Program Git test...")
	testpkg.Otherfunc()
	stringreverse.Reverse("test")

}
