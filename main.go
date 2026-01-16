package main

import (
	"fmt"
	"test/feature1"
	"test/feature2"
)

func main() {
	fmt.Println("Main runing")
	feature1.Feature1()
	feature2.Feature2()
	fmt.Println("Main end")
}
