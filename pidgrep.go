package l0la

import (
	"fmt"
	"strconv"
	"strings"
)

const pidgrepcmd = "ps -p %d | grep -v PID | wc -l"

func PidgrepActive(pid int) bool {
	cmd := fmt.Sprintf(pidgrepcmd, pid)
	stdout, _, _ := Shellout(cmd)
	r, _ := strconv.Atoi(strings.TrimSpace(stdout))
	return r == 1
}

func Pidgrepstatus(pid int) string {
	if PidgrepActive(pid) {
		return "[ACTIVE]"
	} else {
		return "[DEAD?]"
	}
}
