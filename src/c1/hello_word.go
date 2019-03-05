package main

import (
	"fmt"
	"os"
)

func main() {
  fmt.Println("hello Jack GO")
  fmt.Println(os.Args)
  //os.Exit(0)
  os.Exit(-1)
}