# TypeScript 

tsconfig.json  
```
{
    "compilerOptions": {
        "outDir": "./built",
        "allowJs": true,
        "target": "es5"
    },
    "include": [
        "./src/**/*"
    ]
}
这里我们为TypeScript设置了一些东西:

读取所有可识别的src目录下的文件（通过include）。
接受JavaScript做为输入（通过allowJs）。
生成的所有文件放在built目录下（通过outDir）。
将JavaScript代码降级到低版本比如ECMAScript 5（通过target）。
现在，如果你在工程根目录下运行tsc，就可以在built目录下看到生成的文件。 
built下的文件应该与src下的文件相同。 现在你的工程里的TypeScript已经可以工作了。
```
