package main

import (
	"fmt"
	"strconv"
)

func strExtract(s string) string {
	k := 0
	var tmp []rune
	isSlash := false
	for i, r := range s {
		if r == 92 { // "\"
			isSlash = true
			continue
		}

		if r < 97 && false == isSlash {
			if i == 0 {
				return "Invalid string"
			}
			number, _ := strconv.Atoi(string(r))
			if k == 0 {
				k = number
			} else {
				k = k*10 + number
			}
			if i < len(s)-1 {
				continue
			}
		}

		isSlash = false

		for j := 0; j < k-1; j++ {
			tmp = append(tmp, tmp[len(tmp)-1])
		}

		if i == len(s)-1 && k > 0 {
			break
		}

		tmp = append(tmp, r)
		k = 0

	}
	return string(tmp)
}

func main() {
	fmt.Printf("res: %s", strExtract(`a1b1n\4`))
}
