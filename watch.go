package l0la

import "fmt"

const Version = "v0.1.0"

func Watch(pid uint) {
	fmt.Printf("watching pid %d\n", pid)
}
