package wrap

import (
	"strconv"
	"strings"
)

type String string

func (s String) String() string {
	return string(s)
}

// count byte size.
func (s String) Len() int {
	return len(s)
}

// count chars.
func (s String) RuneLen() int {
	return len([]rune(s))
}

func (s String) Trim(cutset string) String {
	trimed := strings.Trim(string(s), cutset)
	return String(trimed)
}

func (s String) TrimLeft(cutset string) String {
	trimed := strings.TrimLeft(string(s), cutset)
	return String(trimed)
}

func (s String) TrimRight(cutset string) String {
	trimed := strings.TrimRight(string(s), cutset)
	return String(trimed)
}

func (s String) Atoi() (int, error) {
	return strconv.Atoi(string(s))
}
