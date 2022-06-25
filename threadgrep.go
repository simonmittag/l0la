package l0la

import (
	"fmt"
	"strconv"
	"strings"
)

const threadgrepcmd = "ps -M %d | grep -v '^USER' | wc -l"

func Threadgrep(pid int) int {
	cmd := fmt.Sprintf(threadgrepcmd, pid)
	stdout, _, _ := Shellout(cmd)
	r, _ := strconv.Atoi(strings.TrimSpace(stdout))
	return int(r)
}
