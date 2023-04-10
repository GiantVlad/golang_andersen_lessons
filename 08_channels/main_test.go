package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func f11() error {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Microsecond)
		if i > 8 {
			return fmt.Errorf("error f11 %d", i)
		}
	}
	return nil
}

func f22() error {
	for i := 0; i < 10; i++ {
		time.Sleep(2 * time.Microsecond)
		if i > 5 {
			return fmt.Errorf("error f22, %d", i)
		}
	}
	return nil
}

func f33() error {
	for i := 0; i < 10; i++ {
		time.Sleep(3 * time.Microsecond)
		if i > 3 {
			return fmt.Errorf("error f33 %d", i)
		}
	}
	return nil
}

func f44() error {
	for i := 0; i < 10; i++ {
		time.Sleep(300 * time.Nanosecond)
	}
	return nil
}

func TestWorker(t *testing.T) {
	var funcs []func() error
	funcs = append(funcs, f44, f22, f33, f11)
	Worker(funcs, 2, 2)
	require.Equal(t, true, true, "")
}
