package l0la

import (
	"os"
	"testing"
)

func TestMemgrep(t *testing.T) {
	mem := Memgrep(os.Getpid())
	if mem < 1 {
		t.Errorf("memory cannot be <1 bytes, was: %v", mem)
	}
}
