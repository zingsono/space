package hdate

import (
	"time"
)

type HTime time.Time

// 对任意struct增加 MarshalJSON ,UnmarshalJSON , String 方法，实现自定义json输出格式与打印方式。
// 时间格式化
func (HTime) MarshalJSON() ([]byte, error) {

}
