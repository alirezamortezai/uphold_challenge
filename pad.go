package main

import (
	"fmt"
	"os"
	"strconv"
)

func advanceCursor(in []byte, pad int) (apd []byte, rem []byte) {
	ch := int(in[0]) - int('0')
	if ch < 0 || ch > 9 {
		return []byte{in[0]}, in[1:]
	}

	apd = []byte{}

	for _, c := range in {
		ch := int(c) - int('0')
		if ch < 0 || ch > 9 {
			break
		}
		apd = append(apd, c)
	}

	apdLen := len(apd)

	rem = in[apdLen:]

	for i := 0; i < pad-apdLen; i++ {
		apd = append([]byte{'0'}, apd...)
	}

	//fmt.Printf("in= %s, apd=%s, rem=%s\n", string(in), string(apd), string(rem))

	return apd, rem

}

func PadWholeNumbers(s string, pad int) string {
	bArr := []byte(s)
	result := []byte{}
	apd := []byte{}
	for len(bArr) > 0 {
		apd, bArr = advanceCursor(bArr, pad)
		result = append(result, apd...)
	}

	return string(result)
}

func Run() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./pad <string to be padded> <length of the pad>")
		os.Exit(-1)
	}

	s := os.Args[1]

	if len(s) < 1 {
		fmt.Println("Usage: ./pad <string to be padded> <length of the pad>")
	}
	p, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Usage: ./pad <string to be padded> <length of the pad>")
		os.Exit(-1)
	}

	res := PadWholeNumbers(s, p)
	fmt.Println(res)
	os.Exit(0)

}
