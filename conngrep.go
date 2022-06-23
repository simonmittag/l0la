package l0la

import (
	"fmt"
	"strconv"
	"strings"
)

const conngrepcmd = "lsof -ai -p %d | grep ESTABLISHED | wc -l"

func Conngrep(pid uint) uint {
	cmd := fmt.Sprintf(conngrepcmd, pid)
	stdout, _, _ := Shellout(cmd)
	r, _ := strconv.Atoi(strings.TrimSpace(stdout))
	return uint(r)
}
