/**
 * 接口开发
 * @params args 参数对象
 * @params db  数据库操作
 * @params http
 */
module.exports = function (z) {
    let args = z.args
    let db = z.db


    return z.then(p=>{
        return p.db.select('')
    })
}
