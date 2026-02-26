// main.go
package main

import (
	"otus-go/internal/repository"
	"otus-go/internal/service"
)

func main() {
	storage := repository.NewStorage()
	generator := service.NewFileGenerator(storage, "./static")
	generator.StartGenerating()
}
