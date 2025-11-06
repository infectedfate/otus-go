package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func generateChessboard(size int) string {
	var builder strings.Builder
	for i := range size {
		for j := range size {
			if (i+j)%2 == 0 {
				builder.WriteString(" ")
			} else {
				builder.WriteString("#")
			}
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func main() {
	size := 8 // Значение по умолчанию
	if len(os.Args) > 1 {
		n, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка: размер должен быть целым числом. Используется значение по умолчанию: 8\n")
			size = 8
		} else if n <= 0 {
			fmt.Fprintf(os.Stderr, "Ошибка: размер должен быть положительным. Используется значение по умолчанию: 8\n")
			size = 8
		} else {
			size = n
		}
	}

	chessboard := generateChessboard(size)
	fmt.Print(chessboard)
}
