package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CountByUsercd(t *testing.T) {

	addTests := []struct {
		usercd   string
		expected int
	}{
		{"a0006802", 1},
		{"a0006801", 0},
	}

	for _, test := range addTests {
		s := CountByUsercd(test.usercd)
		assert.Equal(t, test.expected, s)
	}
}
