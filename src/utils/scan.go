package utils

import (
	"bufio"
	"os"
)

func Scanln(a *string) {
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	*a = string(data)
}
