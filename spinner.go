package l0la

import "strings"

type SpinnerAnim struct {
	chars []string
	index int
}

func NewSpinnerAnim() *SpinnerAnim {
	return &SpinnerAnim{
		chars: strings.Split("⣾⣽⣻⢿⡿⣟⣯⣷", ""),
		index: 0,
	}
}

func (b *SpinnerAnim) Next() string {
	c := b.chars[b.index]
	b.index = (b.index + 1) % len(b.chars)
	return c
}
