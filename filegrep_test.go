package l0la

import (
	"os"
	"testing"
)

func TestFilegrep(t *testing.T) {
	fil := Filegrep(os.Getpid())
	if fil < 1 {
		t.Errorf("open files be <1 bytes, was: %v", fil)
	}
}
