package l0la

import (
	"os"
	"testing"
)

func TestConngrep(t *testing.T) {
	con := Conngrep(os.Getpid())
	if con != 0 {
		t.Errorf("test does not have network conns, was: %v", con)
	}
}
