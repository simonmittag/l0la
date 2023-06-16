package l0la

import (
	"fmt"
	"strconv"
	"strings"
)

const memgrepcmd = "/bin/ps -m -o rss,pid | grep %d | head -1"

func Memgrep(pid int) int {
	cmd := fmt.Sprintf(memgrepcmd, pid)
	stdout, _, _ := Shellout(cmd)
	rss := strings.Split(strings.TrimSpace(stdout), " ")[0]
	r, _ := strconv.Atoi(rss)
	return int(r)
}
