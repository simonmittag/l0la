package l0la

import (
	"fmt"
	"time"
)

const Version = "v0.1.0"

func Watch(pid uint) {
	fmt.Printf("watching pid %d\n", pid)
	for {
		fmt.Printf("\nconns: %d", Conngrep(pid))
		time.Sleep(time.Millisecond * 500)
	}
}
