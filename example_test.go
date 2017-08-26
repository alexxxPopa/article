package article

import (
	"testing"
	"strings"
	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	actual := strings.ToUpper("hello")
	expected := "HELLO"

	assert.Equal(t, expected, actual)
}