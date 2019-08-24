/**
 * @Auth Zengs
 * @Email zingsono@gmail.com
 * @Date 2019年1月5日
 * @Description Oracle数据库操作，依赖 Oracle client
 */
let oracledb = require('oracledb')
// 不设置自动提交
oracledb.autoCommit = false
// 查询超时时间设置 10 seconds, 默认60 seconds
oracledb.queueTimeout = 10000

console.log('Oracle client library version is ' + oracledb.oracleClientVersionString)

/**
 * 默认配置
 * @param config
 * @constructor
 */
let DB = function (config){
    this.config = config||{
        user: 'appadmin',
        password: 'appadmin',
        connectString: '127.0.0.1:1521/orcl'
    }
}

/**
 * 配置连接
 * @param config
 */
DB.prototype.setConfig = function(config){
    this.config = config
}

/**
 * 打开连接
 * @returns {Promise<any>}
 */
DB.prototype.open = function(){
    return new Promise((resolve, reject)=>{
        oracledb.getConnection(this.config, function (err, connection) {
            if(err){
                reject(err)
            }else {
                resolve(connection)
            }
        })
    })
}

/**
 * 执行SQL命令
 * @param sql
 * @returns {Promise<any | never>}
 */
DB.prototype.execute = function(sql){
    return this.open().then((conn)=>{
        return conn.execute(sql)
    })
}

DB.prototype.result = function(sql){
    return this.execute(sql).then(result=>{
        console.log('Column metadata: ', result.metaData)
        console.log('Query results: ')
        console.log(result.rows)
    })
}

/**
 * insert 操作，返回成功记录数
 * @param sql
 * @returns {Promise<any | never | never>}
 */
DB.prototype.insert = function(sql){
    return this.execute(sql).then(result=>{
        console.log('Column metadata: ', result.metaData)
        console.log('Query results: ')
        console.log(result.rows)
    })
}

/**
 * update 操作，返回成功记录数
 * @param sql
 * @returns {Promise<any | never | never>}
 */
DB.prototype.update = function(sql){
    return this.execute(sql).then(result=>{
        console.log('Column metadata: ', result.metaData)
        console.log('Query results: ')
        console.log(result.rows)
    })
}

/**
 * delete 操作，返回成功记录数
 * @param sql
 * @returns {Promise<any | never | never>}
 */
DB.prototype.delete = function(sql){
    return this.execute(sql).then(result=>{
        console.log('Column metadata: ', result.metaData)
        console.log('Query results: ')
        console.log(result.rows)
    })
}

/**
 * select 操作，返回结果集List<Object>
 * @param sql
 * @returns {Promise<any | never | never>}
 */
DB.prototype.select = function(sql){
    return this.execute(sql).then(result=>{
        console.log('Column metadata: ', result.metaData)
        console.log('Query results: ')
        console.log(result.rows)
    })
}

/**
 * select 分页查询，返回分页对象：{ total:100,pageSize:1,pageNum:10,data:[{}] }
 * @param sql
 * @returns {Promise<any | never | never>}
 */
DB.prototype.selectPage = function(sql){
    return this.execute(sql).then(result=>{
        console.log('Column metadata: ', result.metaData)
        console.log('Query results: ')
        console.log(result.rows)
    })
}

/**
 * commit 提交事务
 * @param bool 值为true提交，为false回滚
 */
DB.prototype.commit = function(bool){
    console.log(bool)
}

module.exports = new DB()
