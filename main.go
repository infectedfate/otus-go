package main

import (
	"fmt"
    "otus-go/internal/model"
)

func main() {
	config := model.ServerConfig{
        Port:      8080,
        StaticDir: "./static",
	}

	fmt.Println(config)
}
