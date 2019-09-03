// 模块错误码定义
package libmongo

import (
	"log"
)

// 异常使用 log.Panicf() ，输出日志并抛出
// 外层函数使用recover()  处理
func Throw(err error, args ...interface{}) {
	if err != nil {
		log.Panicf(err.Error(), args)
	}
}
