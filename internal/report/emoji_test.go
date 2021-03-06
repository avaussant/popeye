package report

import (
	"testing"
	"unicode/utf8"

	"github.com/derailed/popeye/internal/issues"
	"github.com/stretchr/testify/assert"
)

func TestEmojiForLevel(t *testing.T) {
	s := new(Sanitizer)
	for k, v := range map[int]int{0: 1, 1: 1, 2: 1, 3: 1, 4: 1, 100: 1} {
		assert.Equal(t, v, utf8.RuneCountInString(s.EmojiForLevel(issues.Level(k))))
	}
}

func TestEmojiUgry(t *testing.T) {
	s := new(Sanitizer)
	s.jurassicMode = true
	for k, v := range map[int]string{0: "OK", 1: "I", 2: "W", 3: "E", 100: "C"} {
		assert.Equal(t, v, s.EmojiForLevel(issues.Level(k)))
	}
}
