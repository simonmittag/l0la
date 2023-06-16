package l0la

import "testing"

func TestSpinnerAnim(t *testing.T) {
	wants := []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}

	sa := NewSpinnerAnim()

	for i := 0; i < len(wants); i++ {
		n := sa.Next()
		if n != wants[i] {
			t.Errorf("got %v want %v index %v", n, wants[i], i)
		} else {
			t.Logf("next: %v", wants[i])
		}
	}
	sa.Next()
}
