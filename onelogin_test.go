package geetest

import (
	"testing"
)

const appID = "zoekwui1hnmg49x5fwzf5la0ml5dziwn"
const appKey = "gywzffojtnzl0vd6kcut8fcgyud5wg49"

func TestOneLogin_CheckPhone(t *testing.T) {
	o := NewOneLogin(appID, appKey)
	s, err := o.CheckPhone("cc", "cc", "cc")
	if err != nil {
		panic(err)
	}
	t.Log(s)
}

func TestOneLogin_DecodePhone(t *testing.T) {
	o := NewOneLogin(appID, appKey)
	phone, err := o.DecodePhone("be28dea08ee543320b1ef9e1bceb51e4")
	if err != nil {
		panic(err)
	}
	t.Log(phone)
}

func TestHmacSha256(t *testing.T) {
	o := NewOneLogin(appID, appKey)
	s := o.GetSign(1542355862990)
	t.Log(s)
}
