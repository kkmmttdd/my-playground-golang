package main

import (
	"fmt"
	"github.com/kkmmttdd/go-playground/error_sample"
)
func main() {
	_, err := error_sample.Render()
	fmt.Println(err)
	_, err = error_sample.Render2()
	fmt.Println(err)
	//fmt.Println(errors.Is(err, &error_sample.QueryError{}))
	//fmt.Println(innerInner == inner)
	//fmt.Println(innerInner == innerInner2)
}