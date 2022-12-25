// usage
// import . "github.com/oguri-souhei/goutil/pkg/wrap"
package wrap

import "strings"

type String string

func (s String) Len() int {
	return len(s)
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
