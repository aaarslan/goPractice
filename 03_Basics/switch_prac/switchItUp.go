package main

import (
	"fmt"
	"runtime"
)

func main() {
  fmt.Print("Go runs on ")
  switch os := runtime.GOOS; os {
  case "darwin":
    fmt.Println("os X")
  case "linux":
    fmt.Println("won't work on this one")
  default:
  fmt.Println("this ain't windows sooon")
  }
}
