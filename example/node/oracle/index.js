
let db = require('./oracle_template')
db.setConfig({
    user: 'unionlive',
    password: 'unionlive',
    connectString: 'proxy.unionlive.com:1521/orcl'
})

db.result('select * from ULTAB_SC_USER').catch(err=>{
    console.error(err)
})



