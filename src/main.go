package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	var name string = "Shivaraju"
	fmt.Printf("Hello World %s \n", name)
	fmt.Println(time.Since(start))
}
