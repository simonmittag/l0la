package l0la

import (
	"fmt"
	"strconv"
	"strings"
)

const cpugrepcmd = "ps -m -o %%cpu,pid | grep %d | head -1"

func Cpugrep(pid int) string {
	cmd := fmt.Sprintf(cpugrepcmd, pid)
	stdout, _, _ := Shellout(cmd)
	cpu := strings.Split(stdout, " ")[0]
	r, _ := strconv.ParseFloat(cpu, 8)
	return fmt.Sprintf("%.2f%%", r)
}
