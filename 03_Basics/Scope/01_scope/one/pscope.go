package main

import (
    "fmt"
    "github.com/aaarslan/goPractice/03_Basics/Scope/01_scope/two"
  )

var x = 42

func main(){
  fmt.Println(x)
  foo()
  fmt.Println(two.Testing())
}

func foo(){
  fmt.Println(x)
}
