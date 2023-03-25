package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTop(t *testing.T) {
	in := `hello one! how are you? Hi    winter super, hi morning super random.    hi yes. Tower less one less one 
			more text`
	out := "1: hi - 3\n2: one - 3\n3: less - 2\n4: super - 2\n"
	require.Equal(t, top(in, 4), out, "test top")
}
