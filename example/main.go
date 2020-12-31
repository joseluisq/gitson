package main

import (
	"fmt"
	"gitson"
	"os"
)

func main() {
	jsonb, err := gitson.Log(os.Args, os.Getenv("REPO"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(jsonb))
}
