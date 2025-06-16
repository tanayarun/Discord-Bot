package main

import "fmt"

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
	}
}
