package main

import (
	"fmt"
	"os"

	"github.com/crazykun/googleAuth"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("please input key")
		return
	}
	key := os.Args[1]
	ga := googleAuth.NewGoogleAuth()
	code, _ := ga.GetCode(key)
	fmt.Println(code)
}
