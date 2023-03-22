package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCount(t *testing.T) {
	var s []string
	s = append(s, `q12aa3b1\46`, "qqqqqqqqqqqqaaaab444444")
	require.Equal(t, strExtract(s[0]), s[1], "test 1")
}
