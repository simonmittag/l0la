package l0la

import "testing"

func TestLPad(t *testing.T) {
	padded := LPad(20, "blah")
	if len(padded) != 20 {
		t.Error("should be padded to 20")
	}
}

func TestFGroup(t *testing.T) {
	got := FGroup(1234567)
	want := "1,234,567"
	if got != want {
		t.Errorf("wanted %v got %v", want, got)
	}
}
