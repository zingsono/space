#!/usr/bin/env node
console.log("xxl-job: es-clean nodejs")
var arguments = process.argv

console.log("脚本位置: " + arguments[1])
console.log("任务参数: " + arguments[2])
console.log("分片序号: " + arguments[3])
console.log("分片总数: " + arguments[4])

// nodejs script
// 清理Elasticsearch日志历史数据

const axios = require('axios');
const moment = require('moment');

(async function () {
    console.log('BEG******************************');
    // 清理120天前的数据
    let date = moment().add(-120,'days').format("YYYY.MM.DD");

    console.log(`logs-logback*******************************`)
    let index = `logs-logback-${date}`;
    console.log(`>准备清理ES索引 ${index}`);

    let url = `http://elastic.sc0:9200/${index}`;
    //let url = `http://elastic.sc0:9200/logs-logback-2019.05.01`;
    console.log(`>请求URL=${url}`);
    try {
        let response = await axios.delete(url);
        if (response.status == 200) {
            console.log(`>成功：清理ES索引 ${index}`);
        }
        console.log(`>httpStatus=${response.status} 清理数据请求响应 ${JSON.stringify(response.data)}`);
    } catch (e) {
        console.log(`>异常：请求URL=${url} 异常信息 ${e.message}`);
    }

    console.log(`logs*******************************`)
    index = `logs-${date}`;
    console.log(`>准备清理ES索引 ${index}`);

    url = `http://elastic.sc0:9200/${index}`;
    //let url = `http://elastic.sc0:9200/logs-logback-2019.05.01`;
    console.log(`>请求URL=${url}`);
    try {
        let response = await axios.delete(url);
        if (response.status == 200) {
            console.log(`>成功：清理ES索引 ${index}`);
        }
        console.log(`>httpStatus=${response.status} 清理数据请求响应 ${JSON.stringify(response.data)}`);
    } catch (e) {
        console.log(`>异常：请求URL=${url} 异常信息 ${e.message}`);
    }
    console.log('END******************************');
    process.exit(0)
})();
