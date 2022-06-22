package l0la

import "fmt"

const Version = "v0.1.0"

func Watch(pid uint) {
	fmt.Println("watching pid %d", pid)
}
