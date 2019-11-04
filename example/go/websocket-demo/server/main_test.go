package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestSid(t *testing.T) {
	t.Log(Sid())
	t.Log(Sid())
	t.Log(Sid())
	t.Log(Sid())
	t.Log(Sid())

}

func TestA(t *testing.T) {

	md5Ctx := md5.New()
	md5Ctx.Write([]byte("test md5 encrypto"))
	cipherStr := md5Ctx.Sum(nil)
	fmt.Print(cipherStr)
	fmt.Print("\n")
	fmt.Println(hex.EncodeToString(cipherStr))
}
