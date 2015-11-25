package main

import (
	"fmt"
	"my_go_module/modules/config"
)

func main() {
	c, err := config.CreateConfig()

	if err != nil {
		fmt.Println(err)
	}

	err = c.SetConfig("db")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(c.Get("key"))
}
