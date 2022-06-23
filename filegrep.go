package l0la

import (
	"fmt"
	"strconv"
	"strings"
)

const filegrepcmd = "lsof -p %d | wc -l"

func Filegrep(pid int) int {
	cmd := fmt.Sprintf(filegrepcmd, pid)
	stdout, _, _ := Shellout(cmd)
	r, _ := strconv.Atoi(strings.TrimSpace(stdout))
	return r
}
