package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFunc(t *testing.T) {

	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str" + s)
	// t.Log(10 + strconv.Atoi("10"))   // 返回两个值
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
	// for i, err :=   .Atoi("10"); err == nil {
	//     t.Log(10 + i)
	// }
}
