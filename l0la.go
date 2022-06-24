package l0la

import (
	"fmt"
	"github.com/gosuri/uilive"
	. "github.com/logrusorgru/aurora"
	"time"
)

const Version = "v0.1.1"

func Watch(pid int) {
	l01 := uilive.New()
	l01.RefreshInterval = time.Hour * 24 * 30
	l01.Start()
	l02 := l01.Newline()
	l03 := l01.Newline()
	l04 := l01.Newline()
	l05 := l01.Newline()
	l06 := l01.Newline()
	l07 := l01.Newline()
	l08 := l01.Newline()
	l09 := l01.Newline()
	l10 := l01.Newline()
	l11 := l01.Newline()
	l12 := l01.Newline()
	l13 := l01.Newline()
	defer l01.Stop()

	pad := 26
	for {
		fmt.Fprint(l01, Logo()+"\n")
		fmt.Fprintf(l02, "watching pid: %v\n", pid)
		fmt.Fprintf(l03, "┏━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓\n")
		fmt.Fprintf(l04, "┃ TCP conns   ┃ %s ┃\n", Red(LPad(pad, FGroup(Conngrep(pid)))))
		fmt.Fprintf(l05, "┣━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫\n")
		fmt.Fprintf(l06, "┃ Open files  ┃ %s ┃\n", Yellow(LPad(pad, FGroup(Filegrep(pid)))))
		fmt.Fprintf(l07, "┣━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫\n")
		fmt.Fprintf(l08, "┃ Threads     ┃ %s ┃\n", Blue(LPad(pad, FGroup(Threadgrep(pid)))))
		fmt.Fprintf(l09, "┣━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫\n")
		fmt.Fprintf(l10, "┃ RSS bytes   ┃ %s ┃\n", Gray(8, LPad(pad, FGroup(Memgrep(pid)))))
		fmt.Fprintf(l11, "┣━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫\n")
		fmt.Fprintf(l12, "┃ CPU percent ┃ %s ┃\n", Magenta(LPad(pad, Cpugrep(pid))))
		fmt.Fprintf(l13, "┗━━━━━━━━━━━━━┻━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛\n")
		l01.Flush()
		time.Sleep(time.Millisecond * 500)
	}
}

func Logo() string {
	return fmt.Sprintf(Red("<").String()+"l"+Blue("0").String()+"la"+Red(">").String()+" %s", Version)
}
