package l0la

import (
	"os"
	"testing"
)

func TestPidGrepActive(t *testing.T) {
	pg := PidgrepActive(os.Getpid())
	if pg == false {
		t.Errorf("curren pid can't be inactive")
	}
}

func TestPidGrepStatus(t *testing.T) {
	ps := Pidgrepstatus(os.Getpid())
	if ps != "[ACTIVE]" {
		t.Errorf("current pid has to be active")
	}
}
