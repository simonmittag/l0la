package l0la

import "testing"

func TestShellout(t *testing.T) {
	so, se, err := Shellout("/bin/ls")
	t.Logf("stdout: %v", so)
	t.Logf("stderr: %v", se)

	if err != nil {
		t.Errorf("command failed, cause: %v", err)
	}
}
