package googleAuth

import (
	"testing"
)

const (
	key = "RXOOV3HI4KGVTEST"
)

func TestGetCode(t *testing.T) {
	ga := GoogleAuth{}
	code, _ := ga.GetCode(key)

	t.Log(code)
	t.Log(ga.VerifyCode(key, code))
}
