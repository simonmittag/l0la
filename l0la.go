package l0la

import (
	"fmt"
	"github.com/gosuri/uilive"
	. "github.com/logrusorgru/aurora"
	"io"
	"time"
)

const Version = "v0.1.4"

func Watch(pid int, seconds ...int) {
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
	s := NewSpinnerAnim()

	timer := time.NewTimer(time.Hour * 24)
	if len(seconds) > 0 && seconds[0] >= 1 {
		timer = time.NewTimer(time.Second * time.Duration(seconds[0]))
	}

	if PidgrepActive(pid) {
	Watch:
		for {
			select {
			case <-timer.C:
				break Watch
			default:

				fmt.Fprint(lines[0], Logo()+"\n")
				fmt.Fprintf(lines[0], "%s watching pid: %v %s \n", s.Next(), pid, Pidgrepstatus(pid))
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
	} else {
		fmt.Printf(Logo())
		fmt.Printf("\n%d not a pid, insufficient privileges or dead?\n", pid)
	}
}

func Logo() string {
	return fmt.Sprintf(Red("<").String()+"l"+Blue("0").String()+"la"+Red(">").String()+" %s", Version)
}
