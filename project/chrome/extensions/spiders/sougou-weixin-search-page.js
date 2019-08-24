// 搜狗关键词搜索微信公众号文章采集

// 输入关键词
let inputKeyword = async (value) => {
    await MS.sleep(3);
    $('.query').val(value);
    console.log('输入关键词：', value);
};

// 点击搜索按钮
let clickSearch = async () => {
    await MS.sleep(3);
    $('.swz')[0].click();
    console.log('点击搜文章按钮');
};

// 点击目标文章链接，打开标签页
let _blankClick = async (seconds = 120) => {
    await MS.sleep(3);
    let wxli = $('.news-list li');
    console.log('匹配数据列表数量: ', wxli.length);
    for (let item of wxli) {
        let a = $(item).children('.txt-box').children('h3').children('a');
        await MS.sleep(seconds);
        console.log('执行点击事件 ');
        a[0].click();
    }
};

// 点击下一页按钮
let npClick = async () => {
    await MS.sleep(10);
    $('.np') && $('.np')[0].click();
}

(async (is) => {
    // 非微信搜索界面，不做任何处理
    if (!is) {
        return
    }
    console.log('#Init Chrome-extensions:sougou-weixin-search.js')
    console.log('#当前URL：', decodeURIComponent(location.href))

    // 关键词，可以改为根据接口获取关键词
    let keyword = '流行品牌';

    // 当前链接中不存在关键词时，输入关键词点击搜索
    if (location.href.indexOf(encodeURIComponent(keyword)) == -1) {
        // 输入关键词
        await inputKeyword(keyword)
        // 点击搜索
        await clickSearch();
    }

    // 读取搜索结果列表，自动打开新标签页
    // await _blankClick(60);

    // 点击下一页
    await npClick();

})(location.href.indexOf('weixin.sogou.com/weixin') > 0);
