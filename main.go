package main

import (
	"fmt"
	"test/modules/config"
)

func main() {
	c, err := config.CreateConfig(config.DEFAULT)

	if err != nil {
		fmt.Println(err)
	}

	err = c.SetConfig("db.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(c.Get("key"))
}
