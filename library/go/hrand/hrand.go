package hrand

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

// 生成32位随机字符串
func Id32() string {
	s := time.Now().String() + fmt.Sprintf("%d", rand.Int()) + fmt.Sprintf("%d", rand.Int())
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(s))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}
