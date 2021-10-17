package main

import (
	"strconv"
	"strings"
)

//0-1
func stringAdd(a, b string) string {
	s1 := strings.Split(a, ".")
	s2 := strings.Split(b, ".")

	a1, b1 := s1[1], s2[1]

	if len(a1) < len(b1) {
		a1, b1 = b1, a1
	}
	res := strings.Split(a1, "")

	high := 0
	for i := len(b1)-1; i >= 0; i-- {
		atoi1, _ := strconv.Atoi(string(b1[i]))
		atoi2, _ := strconv.Atoi(string(a1[i]))

		v := atoi1 + atoi2 + high

		high = v/10
		res[i] = strconv.Itoa(v%10)
	}

	return strconv.Itoa(high) + "." + strings.Join(res, "")
}

func main() {
	stringAdd("0.99", "0.01")
}
