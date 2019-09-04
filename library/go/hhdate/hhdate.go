// 日期时间处理模块
package hhdate

import (
	"time"
)

// 时间转换的模板，golang里面只能是 "2006-01-02 15:04:05" （go的诞生时间）
const (
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	// Handy time stamps.
	Stamp      = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"

	// 上面是go内置time包中的定义
	// ------------------------------------------
	// 下面是常用格式定义
	YMDHMS = "2006-01-02 15:04:05"
	YMD    = "2006-01-02"
	HMS    = "15:04:05"
)

// 获取当前时间戳
func Time() int64 {
	return time.Now().Unix()
}

// 时间戳（秒）转time
/*func GetTime(v time.Duration) *time.Timer  {
	return time.NewTimer(v*time.Second)
}*/

// 时间戳（秒）转指定格式时间字符串
func GetYMDHMS() {

}
