package wrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	testCases := []struct {
		desc string
		r    String
		exp  string
	}{
		{
			desc: "should return string",
			r:    String("test"),
			exp:  "test",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			act := tC.r.String()
			assert.Equal(t, tC.exp, act)
		})
	}
}

func TestLen(t *testing.T) {
	testCases := []struct {
		desc string
		r    String
		exp  int
	}{
		{
			desc: "should count byte size",
			r:    String("test"),
			exp:  4,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			act := tC.r.Len()
			assert.Equal(t, tC.exp, act)
		})
	}
}

func TestRuneLen(t *testing.T) {
	testCases := []struct {
		desc string
		r    String
		exp  int
	}{
		{
			desc: "should count chars",
			r:    String("test"),
			exp:  4,
		},
		{
			desc: "should count chars(zenkaku)",
			r:    "てすと",
			exp:  3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			act := tC.r.RuneLen()
			assert.Equal(t, tC.exp, act)
		})
	}
}

func TestTrim(t *testing.T) {
	testCases := []struct {
		desc string
		r    String
		arg  string
		exp  String
	}{
		{
			desc: "should trim cutset",
			r:    String("   test   "),
			arg:  " ",
			exp:  "test",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			act := tC.r.Trim(tC.arg)
			assert.Equal(t, tC.exp, act)
		})
	}
}

func TestTrimLeft(t *testing.T) {
	testCases := []struct {
		desc string
		r    String
		arg  string
		exp  String
	}{
		{
			desc: "should trim cutset",
			r:    String("   test   "),
			arg:  " ",
			exp:  "test   ",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			act := tC.r.TrimLeft(tC.arg)
			assert.Equal(t, tC.exp, act)
		})
	}
}

func TestRight(t *testing.T) {
	testCases := []struct {
		desc string
		r    String
		arg  string
		exp  String
	}{
		{
			desc: "should trim cutset",
			r:    String("   test   "),
			arg:  " ",
			exp:  "   test",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			act := tC.r.TrimRight(tC.arg)
			assert.Equal(t, tC.exp, act)
		})
	}
}

func TestAtoi(t *testing.T) {
	type exp struct {
		value int
	}
	testCases := []struct {
		desc string
		r    String
		exp  exp
		err  bool
	}{
		{
			desc: "should return int",
			r:    String("1"),
			exp: exp{
				value: 1,
			},
			err: false,
		},
		{
			desc: "should return error",
			r:    String("a"),
			exp: exp{
				value: 0,
			},
			err: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			act, err := tC.r.Atoi()
			assert.Equal(t, tC.exp.value, act)
			if tC.err {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
