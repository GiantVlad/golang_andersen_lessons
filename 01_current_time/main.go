package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Errorf("error: %v", err)
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("Current time %s", time)
}
