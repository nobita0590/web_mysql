package main

import "flag"
import "fmt"

func main()  {
	wordPtr := flag.String("env", "foo", "a string")
	flag.Parse()
	fmt.Println("env:", *wordPtr)
}
