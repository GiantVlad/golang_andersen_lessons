package main

import (
	"fmt"
	"go_andr_less/12_envdir/runner"
	"os"
)

func main() {
	pwd, _ := os.Getwd()
	eVals, err := runner.ReadDir(pwd + "/12_envdir/envdir")
	if err != nil {
		panic(err)
	}
	cmd := []string{pwd + "/12_envdir/script.sh"}
	runner.RunCmd(cmd, eVals)
	fmt.Printf("env %v, %v", eVals, err)
}
