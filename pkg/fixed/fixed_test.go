package pkg

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func TestUnmarshal(t *testing.T) {
	type args struct {
		buf []byte
		v   interface{}
		d   transform.Transformer
	}
	type exp struct {
		v   interface{}
		err error
	}
	type testA struct {
		X string `fixed:"0,1"`
		Y string `fixed:"1,3"`
	}
	type testB struct {
		X string `fixed:"0,2"`
		Y string `fixed:"2,6"`
	}
	testCases := []struct {
		desc string
		args args
		exp  exp
	}{
		{
			desc: "decode byte (utf-8) data to structure",
			args: args{
				buf: []byte{97, 98, 99}, // abc
				v:   &testA{},
				d:   nil,
			},
			exp: exp{
				v: &testA{
					X: "a",
					Y: "bc",
				},
				err: nil,
			},
		},
		{
			desc: "decode byte data (sjis) to structure",
			args: args{
				buf: []byte{130, 160, 130, 162, 130, 164}, // あいう(sjis)
				v:   &testB{},
				d:   japanese.ShiftJIS.NewDecoder(),
			},
			exp: exp{
				v: &testB{
					X: "あ",
					Y: "いう",
				},
				err: nil,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			act, err := Unmarshal(tC.args.buf, tC.args.v, tC.args.d)
			assert.Equal(t, tC.exp.v, act)
			assert.Equal(t, tC.exp.err, err)
		})
	}
}

func TestDecode(t *testing.T) {
	type args struct {
		buf []byte
		t   transform.Transformer
	}
	type exp struct {
		rslt []byte
		err  error
	}
	testCases := []struct {
		desc string
		args args
		exp  exp
	}{
		{
			desc: "decode sjis to utf-8",
			args: args{
				buf: []byte{130, 160, 130, 162, 130, 164}, // あいう(sjis)
				t:   japanese.ShiftJIS.NewDecoder(),
			},
			exp: exp{
				rslt: []byte("あいう"),
				err:  nil,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			rslt, err := transfer(tC.args.buf, tC.args.t)
			assert.Equal(t, tC.exp.rslt, rslt)
			assert.Equal(t, tC.exp.err, err)
		})
	}
}

func TestParseFixedTag(t *testing.T) {
	type exp struct {
		begin, end int
	}
	type testA struct {
		A string `fixed:"0,1"`
	}
	type invalid struct {
		A string `fixed:"foo"`
	}
	testCases := []struct {
		desc  string
		arg   reflect.StructField
		exp   exp
		panic bool
	}{
		{
			desc: "parse fixed-struct tag",
			arg:  reflect.TypeOf(testA{}).Field(0),
			exp: exp{
				begin: 0,
				end:   1,
			},
		},
		{
			desc:  "panic when invalid fixed-tag",
			arg:   reflect.TypeOf(invalid{}).Field(0),
			panic: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.panic {
				assert.Panics(t, func() { parseFixedTag(tC.arg) })
			} else {
				begin, end := parseFixedTag(tC.arg)
				assert.Equal(t, tC.exp.begin, begin)
				assert.Equal(t, tC.exp.end, end)
			}
		})
	}
}
