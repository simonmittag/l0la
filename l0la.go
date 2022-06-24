package l0la

import (
	"fmt"
	"github.com/gosuri/uilive"
	. "github.com/logrusorgru/aurora"
	"io"
	"time"
)

const Version = "v0.1.1"

func Watch(pid int) {
	lines := make([]io.Writer, 0)
	l0 := uilive.New()
	l0.RefreshInterval = time.Hour * 24 * 30
	l0.Start()
	defer l0.Stop()
	lines = append(lines, l0)

	for i := 0; i < 13; i++ {
		lines = append(lines, lines[0].(*uilive.Writer).Newline())
	}

	pad := 26
	for {
		fmt.Fprint(lines[0], Logo()+"\n")
		fmt.Fprintf(lines[0], "watching pid: %v\n", pid)
		fmt.Fprintf(lines[0], "┏━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓\n")
		fmt.Fprintf(lines[0], "┃ TCP conns   ┃ %s ┃\n", Red(LPad(pad, FGroup(Conngrep(pid)))))
		fmt.Fprintf(lines[0], "┣━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫\n")
		fmt.Fprintf(lines[0], "┃ Open files  ┃ %s ┃\n", Yellow(LPad(pad, FGroup(Filegrep(pid)))))
		fmt.Fprintf(lines[0], "┣━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫\n")
		fmt.Fprintf(lines[0], "┃ Threads     ┃ %s ┃\n", Blue(LPad(pad, FGroup(Threadgrep(pid)))))
		fmt.Fprintf(lines[0], "┣━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫\n")
		fmt.Fprintf(lines[0], "┃ RSS bytes   ┃ %s ┃\n", Gray(8, LPad(pad, FGroup(Memgrep(pid)))))
		fmt.Fprintf(lines[0], "┣━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫\n")
		fmt.Fprintf(lines[0], "┃ CPU percent ┃ %s ┃\n", Magenta(LPad(pad, Cpugrep(pid))))
		fmt.Fprintf(lines[0], "┗━━━━━━━━━━━━━┻━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛\n")
		l0.Flush()
		time.Sleep(time.Millisecond * 500)
	}
}

func Logo() string {
	return fmt.Sprintf(Red("<").String()+"l"+Blue("0").String()+"la"+Red(">").String()+" %s", Version)
}
