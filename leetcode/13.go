package main

import (
	"fmt"
)

func romanToInt(s string) int {
	charMaps := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	//
	//I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
	//X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
	//C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。

	// IV
	// IX
	//
	// LVIII
	// MCMXCIV

	prev := byte(0)
	sum := 0
	for _, current := range []byte(s) {
		if prev == 0 {
			prev = current
			continue
		}

		if current == 'V' {
			if prev == 'I' {
				prev = 0
				sum += 4
			} else {
				sum += charMaps[prev]
				prev = current
			}
		} else if current == 'X' {
			if prev == 'I' {
				prev = 0
				sum += 9
			} else {
				sum += charMaps[prev]
				prev = current
			}
		} else if current == 'L' {
			if prev == 'X' {
				prev = 0
				sum += 40
			} else {
				sum += charMaps[prev]
				prev = current
			}
		} else if current == 'C' {
			if prev == 'X' {
				prev = 0
				sum += 90
			} else {
				sum += charMaps[prev]
				prev = current
			}
		} else if current == 'D' {
			if prev == 'C' {
				prev = 0
				sum += 400
			} else {
				sum += charMaps[prev]
				prev = current
			}
		} else if current == 'M' {
			if prev == 'C' {
				prev = 0
				sum += 900
			} else {
				sum += charMaps[prev]
				prev = current
			}
		} else {
			// 不是以上六中情况 直接把 prev 的数据加进去
			sum += charMaps[prev]
			prev = current
		}
	}
	// 不要遗漏这个数据
	sum += charMaps[prev]
	return sum
}

func main() {
	fmt.Println(romanToInt("VIII"))
}
