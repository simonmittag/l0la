package l0la

import (
	"os"
	"testing"
)

func TestWatch(t *testing.T) {
	Watch(os.Getpid(), 4)
}
