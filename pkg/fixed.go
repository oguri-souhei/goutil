package pkg

import (
	"reflect"
	"strconv"
	"strings"

	"golang.org/x/text/transform"
)

// decode bytes to structure.
func Unmarshal(buf []byte, v interface{}, t transform.Transformer) (interface{}, error) {
	rv, rt := reflectFromPtr(v)
	for i := 0; i < rv.Elem().NumField(); i++ {
		begin, end := parseFixedTag(rt.Field(i))
		col, err := decode(buf[begin:end], t)
		if err != nil {
			return nil, err
		}
		field := rv.Elem().Field(i)
		field.SetString(string(col))
	}
	return v, nil
}

func reflectFromPtr(v interface{}) (reflect.Value, reflect.Type) {
	rv := reflect.ValueOf(v)
	rt := reflect.Indirect(rv).Type()
	if rt.Kind() != reflect.Struct {
		panic("non-structure value.")
	}
	return rv, rt
}

func decode(buf []byte, t transform.Transformer) ([]byte, error) {
	if t == nil {
		return buf, nil
	}
	rslt, _, err := transform.Bytes(t, buf)
	return rslt, err
}

// get struct-tag fixed.
func parseFixedTag(field reflect.StructField) (int, int) {
	// get struct tag.
	fixed := field.Tag.Get("fixed")
	// parse struct tag.
	splited := strings.Split(fixed, ",")
	begin, err := strconv.Atoi(splited[0])
	if err != nil {
		panic(err)
	}
	end, err := strconv.Atoi(splited[1])
	if err != nil {
		panic(err)
	}
	return begin, end
}
