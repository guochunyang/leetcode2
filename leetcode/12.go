package main

import (
	"bytes"
	"fmt"
)

func intToRoman(num int) string {

	buffer := bytes.NewBufferString("")

	// < 3999
	m := num / 1000
	if m > 0 {
		switch m {
		case 1:
			fallthrough
		case 2:
			fallthrough
		case 3:
			for i := 0; i < m; i++ {
				//fmt.Print("M")
				buffer.WriteByte('M')
			}
		default:
			panic("error")
		}
	}
	num = num % 1000

	c := num / 100
	if c > 0 {
		// C 100
		// D 500
		// M 1000
		switch c {
		case 1:
			fallthrough
		case 2:
			fallthrough
		case 3:
			for i := 0; i < c; i++ {
				//fmt.Print("C")
				buffer.WriteByte('C')
			}
		case 4:
			//fmt.Print("CD")
			buffer.WriteString("CD")
		case 5:
			//fmt.Print("D")
			buffer.WriteByte('D')
		case 6:
			fallthrough
		case 7:
			fallthrough
		case 8:
			//fmt.Print("D")
			buffer.WriteByte('D')
			for i := 0; i < c-5; i++ {
				//fmt.Print("C")
				buffer.WriteByte('C')
			}
		case 9:
			//fmt.Print("CM")
			buffer.WriteString("CM")
		}

	}
	num = num % 100

	x := num / 10
	// X 10
	// L 50
	// C 100
	if x > 0 {
		switch x {
		case 1:
			fallthrough
		case 2:
			fallthrough
		case 3:
			for i := 0; i < x; i++ {
				//fmt.Print("X")
				buffer.WriteByte('X')
			}
		case 4:
			//fmt.Print("XL")
			buffer.WriteString("XL")
		case 5:
			//fmt.Print("L")
			buffer.WriteByte('L')
		case 6:
			fallthrough
		case 7:
			fallthrough
		case 8:
			//fmt.Print("L")
			buffer.WriteByte('L')
			for i := 0; i < x-5; i++ {
				fmt.Print("X")
				buffer.WriteByte('X')
			}
		case 9:
			//fmt.Print("XC")
			buffer.WriteString("XC")
		}
	}

	num = num % 10
	if num > 0 {
		switch num {
		case 1:
			fallthrough
		case 2:
			fallthrough
		case 3:
			for i := 0; i < num; i++ {
				//fmt.Print("I")
				buffer.WriteByte('I')
			}
		case 4:
			//fmt.Print("IV")
			buffer.WriteString("IV")
		case 5:
			//fmt.Print("V")
			buffer.WriteByte('V')
		case 6:
			fallthrough
		case 7:
			fallthrough
		case 8:
			//fmt.Print("V")
			buffer.WriteByte('V')
			for i := 0; i < num-5; i++ {
				//fmt.Print("I")
				buffer.WriteByte('I')
			}
		case 9:
			//fmt.Print("IX")
			buffer.WriteString("IX")
		}
		//fmt.Println()
	}

	return buffer.String()
}

func main() {
	fmt.Println(intToRoman(1994))
}
