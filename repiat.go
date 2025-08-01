package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func repiat() bool {
	for {
		fmt.Print("\nПовторить операцию?\n  (y -да)\n  (n - нет)")
		work, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		work = strings.TrimSpace(work)

		switch work {
		case "y":
			return true
		case "n":
			return false
		default:
			fmt.Println(`Введите "y" или "n"`)
		}
	}
}
