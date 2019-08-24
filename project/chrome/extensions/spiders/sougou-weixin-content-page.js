// 微信公众号内容页面

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

// 点击目标文章链接


// 点击下一页按钮


(async (is) => {
    // 非微信搜索界面，不做任何处理
    if (!is) {
        return
    }


})(location.href.indexOf('weixin.sogou.com/weixin') > 0);
