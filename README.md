# mardown-blog
这是一个实验性的，基于 `Golang` 开发的静态 **Markdown博客** 程序
> 如果你和我一样，平时喜欢使用`markdown`来记录自己的工作与生活中的点滴，又希望把这些记录生成个人博客，那`mardown-blog`再适合不过了，它部署简单，可以把md文件快速变为个人博客。该程序的优点：简洁、轻快，安全，等你体验。

## 示例
> https://blog.gaowei.tech

## 支持平台
> Windows、Linux、Mac OS

## 安装
1. 下载 [release](https://github.com/gaowei-space/markdown-blog/releases/)

2. 解压
    ```shell
    tar zxf markdown-blog-v0.0.1-linux-arm64.tar.gz
    ```

3. 创建 markdown 文件目录
    ```shell
    cd markdown-blog-linux-arm64
    mkdir md
    ```

4. 运行
    ```shell
    ./markdown-blog web
    ```

5. 访问 http://127.0.0.1:5006，查看效果

## 使用
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

### 导航排序
> 博客导航默认按照`字典`排序，可以通过 `@` 前面的数字来自定义顺序

##### 个人博客目录示例：

##### 博客导航展示如图：


## 开发
1. 安装 `Golang` 开发环境
2. Fork [源码](https://github.com/gaowei-space/gocron)
3. 启动 web服务
    > 运行之后访问地址 http://localhost:5006，API请求会转发给 markdown-blog
    ```shell
    make run
    ```

4. 编译
    ```shell
    make
    ```

5. 打包
    > 在 markdown-blog-package 生成当前系统的压缩包 markdown-blog-v0.0.1-darwin-arm64.tar
    ```shell
    make package
    ```

6. 生成 Windows、Linux、Mac 的压缩包
    > 在 markdown-blog-package 生成压缩包 markdown-blog-v0.0.1-darwin-arm64.tar markdown-blog-v0.0.1-linux-arm64.tar.gz markdown-blog-v0.0.1-windows-arm64.zip
    ```shell
    make package-all
    ```
