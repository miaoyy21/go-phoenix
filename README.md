# go-phoenix

## 在线演示
[项目在线演示地址](http://www.nsmei.com:8090/)

## 本地开发模式运行

1. 下载本项目至本地;
2. 使用根目录提供的mysql.sql脚本初始化数据库，目前没有对其他类型数据库做适配;
3. 执行 go mod init;
4. 修改根目录下的config.json文件，编译后即可运行;

## 打包发布
将 [前端项目](https://github.com/miaoyy21/webix-phoenix)使用webpack打包后，拷贝dist目录下的所有文件和src目录下的assets文件夹至本项目的www目录下即可。
