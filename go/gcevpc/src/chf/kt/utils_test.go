package kt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleLineElide(t *testing.T) {
	assert.Equal(t, "hello", SingleLineElide("hello", -1))
	assert.Equal(t, "hello", SingleLineElide("hello\r\n", -1))

	assert.Equal(t, "hello", SingleLineElide("hello", 100))
	assert.Equal(t, "hello", SingleLineElide("hello", 5))
	assert.Equal(t, "h...", SingleLineElide("hello", 4))
	assert.Equal(t, "...", SingleLineElide("hello", 3))
	assert.Equal(t, "hello", SingleLineElide("hello", 2))

	assert.Equal(t, "hello", SingleLineElide("\r\n\r\nhello\r\n\n", 5))
	assert.Equal(t, "h...", SingleLineElide("\r\n\r\nhello", 4))
}
