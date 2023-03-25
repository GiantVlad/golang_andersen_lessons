package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Top struct {
	Word string
	Rate int
}

func top(s string, amount int) string {
	words := strings.Split(s, " ")
	top := map[string]int{}
	for _, word := range words {
		word = strings.ToLower(strings.Trim(word, " ,?.!\n\t"))
		if len(word) > 0 {
			top[word] = top[word] + 1
		}
	}
	res := make([]Top, 0, len(top))
	for key, val := range top {
		res = append(res, Top{key, val})
	}

	sort.Slice(res, func(i, j int) bool {
		if res[j].Rate == res[i].Rate {
			return res[j].Word > res[i].Word
		}
		return res[j].Rate < res[i].Rate
	})
	if amount > len(res) {
		amount = len(res)
	}
	res = res[:amount]

	resStr := ""
	for idx, itemTop := range res {
		resStr += strconv.Itoa(idx+1) + ": " + itemTop.Word + " - " + strconv.Itoa(itemTop.Rate) + "\n"
	}

	return resStr
}

func main() {
	input := `hello one! how are you? Hi    winter super, hi morning super random.    hi yes. Tower less one less one`
	fmt.Printf("%s", top(input, 20))
}
