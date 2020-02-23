package main

import (
	"fmt"
	"strings"
)

func letterCombinations(digits string) []string {
	result := make([]string, 0)

	chMaps := map[int32][]int32{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	tmp := make([]*strings.Builder, 0)
	for _, digit := range digits {
		chs := chMaps[digit]
		tmp2 := make([]*strings.Builder, 0)
		for _, ch := range chs {
			if len(tmp) == 0 {
				n := &strings.Builder{}
				n.WriteByte(byte(ch))
				tmp2 = append(tmp2, n)
			} else {
				// 将 tmp 里面的元素全部加上 ch
				for _, t := range tmp {
					n := &strings.Builder{}
					n.WriteString(t.String())
					n.WriteByte(byte(ch))
					tmp2 = append(tmp2, n)
				}
			}

		}
		tmp = tmp2
	}

	for _, t := range tmp {
		result = append(result, t.String())
	}
	return result
}

func main() {
	fmt.Println(letterCombinations("23"))
}
