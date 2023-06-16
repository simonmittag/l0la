package l0la

import (
	"os"
	"testing"
)

func TestThreadGrep(t *testing.T) {
	ts := Threadgrep(os.Getpid())
	if ts < 1 {
		t.Errorf("thread count for proc can't be 0 or negative, was: %v", ts)
	} else {
		t.Logf("thread count for pid was: %v", ts)
	}
}
