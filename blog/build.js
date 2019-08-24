/**
 * 编译.md文件为html文件，输出到dist目录，且保留目录结构
 */
const marked = require('marked');

// 配置选项 https://marked.js.org/#/USING_ADVANCED.md#options
marked.setOptions({
    baseUrl:'',
    headerIds: false,
});

console.log(marked('# 文件+'));

/**
 * 编译步骤：
 * 1.  拷贝资源文件assets目录到dist目录；
 * 2.  循环读取note目录所有.md文件；
 * 3.  编译.md文件为.html文件，输出到dist，输出目录结构同note目录;
 * 3.1 获取一级文件夹下所有文件夹名、文件名作为侧边栏导航目录；
 */
