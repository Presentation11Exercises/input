package main

import (
	"fmt"
	)

func main() {
  fmt.Println("Enter some data and press enter: ")
  var userData string
  fmt.Scanln(&userData)
  fmt.Println(userData)
  
}
