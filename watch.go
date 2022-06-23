package l0la

import (
	"fmt"
	"github.com/gosuri/uilive"
	"time"
)

const Version = "v0.1.0"

func Watch(pid int) {
	fmt.Printf("watching pid %s\n", FGroup(pid))

	l1 := uilive.New()
	l1.RefreshInterval = time.Hour * 24 * 30
	l1.Start()
	l2 := l1.Newline()
	defer l1.Stop()

	for {
		fmt.Fprintf(l1, "TCP conns: %s\n", FGroup(Conngrep(pid)))
		fmt.Fprintf(l2, "Open files: %s\n", FGroup(Filegrep(pid)))
		l1.Flush()
		time.Sleep(time.Millisecond * 500)
	}
}
