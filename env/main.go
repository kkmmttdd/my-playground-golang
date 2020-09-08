package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
	"os"
)

type Config struct {
	Url string `default:"http://localhost"`
	User string `default:"waaaay"`
}


func main() {
	config := Config{}
	fmt.Println(os.Getenv("URL"))
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	if err := envconfig.Process("", &config); err != nil {
		log.Fatal(err)
	}
	fmt.Println(config.Url)
	fmt.Println(config.User)
}