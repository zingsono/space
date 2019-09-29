// nodejs script
// 清理Elasticsearch日志历史数据
const axios = require('axios');
const moment = require('moment');
console.log('******************************');
(async function () {
    // 清理120天前的数据
    let date = moment().add(-120,'days').format("YYYY.MM.DD");
    let index = `logs-logback-${date}`;
    console.log(`>准备清理ES索引 ${index}`);

    //let url = `http://elastic-sc0:9200/${index}`;
    let url = `http://elastic.sc0:9200/logs-logback-2019.05.01`;
    try {
        console.log(`>请求URL=${url}`);
        let response = await axios.delete(url);
        if (response.status == 200) {
            console.log(`>成功：清理ES索引 ${index}`);
        }
        console.log(`>httpStatus=${response.status} 清理数据请求响应 ${response.data}`);
    } catch (e) {
        console.log(`>异常：请求URL=${url} 异常信息 ${e.message}`);
    }
})();
