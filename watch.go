package l0la

import (
	"fmt"
	"github.com/gosuri/uilive"
	. "github.com/logrusorgru/aurora"
	"time"
)

const Version = "v0.1.0"

func Watch(pid int) {
	fmt.Printf("watching pid %s\n", FGroup(pid))

	l1 := uilive.New()
	l1.RefreshInterval = time.Hour * 24 * 30
	l1.Start()
	l2 := l1.Newline()
	l3 := l1.Newline()
	l4 := l1.Newline()
	defer l1.Stop()

	for {
		fmt.Fprintf(l1, "TCP conns: %s\n", Red(FGroup(Conngrep(pid))))
		fmt.Fprintf(l2, "Open files: %s\n", Yellow(FGroup(Filegrep(pid))))
		fmt.Fprintf(l3, "Threads: %s\n", Blue(FGroup(Threadgrep(pid))))
		fmt.Fprintf(l4, "RSS bytes: %s\n", Gray(8, FGroup(Memgrep(pid))))
		l1.Flush()
		time.Sleep(time.Millisecond * 500)
	}
}
