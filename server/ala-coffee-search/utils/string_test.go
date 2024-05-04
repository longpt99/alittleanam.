package utils

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	length := uint(6)
	numStr := RandomNumStr(length)
	assert.NotEmpty(t, numStr)
	assert.Len(t, numStr, int(length))

	randStr := GenerateString(length)
	assert.NotEmpty(t, randStr)
	assert.Regexp(t, regexp.MustCompile("[#$@!%&*?]"), randStr)
	assert.NotRegexp(t, regexp.MustCompile("[#$@!%&*?]"), numStr)
}
