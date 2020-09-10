package main

import (
	"fmt"
	"github.com/kkmmttdd/my-playground-golang/error/utils"
)
func main() {
	_, err := utils.Render()
	fmt.Println(err)
	_, err = utils.Render2()
	fmt.Println(err)
}