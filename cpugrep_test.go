package l0la

import (
	"os"
	"strings"
	"testing"
)

func TestCpuGrep(t *testing.T) {
	cpu := Cpugrep(os.Getpid())
	t.Logf("cpu percent was %v", cpu)
	if !strings.HasSuffix(cpu, "%") {
		t.Errorf("invalid cpu pc, got: %v", cpu)
	}
}
