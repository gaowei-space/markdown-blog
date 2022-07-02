# mardown-blog
这是一个实验性的，基于`Golang`开发的静态 **Markdown博客** 程序

## 示例
> https://blog.gaowei.tech

## 支持平台
> Windows、Linux、Mac OS

## 安装
1. 下载 [release](https://github.com/gaowei-space/markdown-blog/releases/)

2. 解压
    ```
    tar zxf markdown-blog-v0.0.1-linux-arm64.tar.gz
    ```

3. 创建 markdown 文件目录
    ```
    cd markdown-blog-linux-arm64
    mkdir md
    ```

4. 运行
    ```
    ./markdown-blog web
    ```

5. 访问 http://127.0.0.1:5006，查看效果


### 命令
- markdown-blog
    - -h 查看版本
    - web 运行博客服务
- markdown-blog web
   - --dir value, -d value    指定markdown文件夹，默认：./md/
   - --title value, -t value  web服务标题，默认："Blog"
   - --port value, -p value   web服务端口，默认：5006
   - --env value, -e value    运行环境, 可选：dev,test,prod，默认："prod"
   - -h                       查看版本

## 开发
1. 安装 Go1.9+
1. fork [源码](https://github.com/gaowei-space/gocron)
2. 启动web服务
    > 运行之后访问地址 http://localhost:5006，API请求会转发给 markdown-blog
    ```
    make run
    ```

3. 编译
    ```
    make
    ```

4. 打包
    > 在 markdown-blog-package 生成当前系统的压缩包 markdown-blog-v0.0.1-darwin-arm64.tar
    ```
    make package
    ```

5. 生成 Windows、Linux、Mac 的压缩包
    > 在 markdown-blog-package 生成压缩包 markdown-blog-v0.0.1-darwin-arm64.tar markdown-blog-v0.0.1-linux-arm64.tar.gz markdown-blog-v0.0.1-windows-arm64.zip
    ```
    make package-all
    ```
