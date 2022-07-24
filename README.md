# 🍭 Markdown-Blog
[![GitHub branch checks state](https://img.shields.io/github/checks-status/gaowei-space/meituan-pub-union/main)](https://github.com/gaowei-space/markdown-blog/tree/main)
[![GitHub issues](https://img.shields.io/github/issues/gaowei-space/markdown-blog?color=blueviolet)](https://github.com/gaowei-space/markdown-blog/issues)
[![StyleCI](https://github.styleci.io/repos/494669204/shield?branch=main&style=flat)](https://github.styleci.io/repos/494669204?branch=main)
[![Latest Release](https://img.shields.io/github/v/release/gaowei-space/markdown-blog)](https://github.com/gaowei-space/markdown-blog/releases)
[![GitHub license](https://img.shields.io/github/license/gaowei-space/markdown-blog)](https://github.com/gaowei-space/markdown-blog/blob/main/LICENSE)

[Markdown-Blog](https://github.com/gaowei-space/markdown-blog) 是一款小而美的**Markdown静态博客**程序
> 如果你和我一样，平时喜欢使用`markdown`文件来记录自己的工作与生活中的点滴，又希望把这些记录生成个人博客，那[Markdown-Blog](https://github.com/gaowei-space/markdown-blog)再适合不过了。它简洁、轻快，部署简单，可以把markdown文件快速变为个人博客，它不需要管理后台，无需进行文章的二次发布。

## 案例
> [TechMan'Blog](https://blog.gaowei.tech)

- PC端
<img alt="PC端" src="https://user-images.githubusercontent.com/10205742/176992945-6016193f-e27e-4b19-bf5d-27ff4dfe1fdc.png">

- 移动端

<img alt="移动端" src="https://user-images.githubusercontent.com/10205742/180644851-e3760085-9668-4675-9bab-65c361dd5195.jpeg">



## 支持平台
> Windows、Linux、Mac OS

## 安装
1. 下载 [release](https://github.com/gaowei-space/markdown-blog/releases/)

2. 解压
    ```shell
    tar zxf markdown-blog-v0.0.2-linux-arm64.tar.gz
    ```

3. 创建 markdown 文件目录
    ```shell
    cd markdown-blog-linux-arm64
    mkdir md
    echo "### Hello World" > ./md/主页.md
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
   - --index value, -i value  设置默认首页的文件名称, 默认为空
   - -h                       查看版本

### 关于默认首页
> 如果启动是未指定`index`，程序默认以导航中的第一个文件作为首页

### 导航排序
> 博客导航默认按照`字典`排序，可以通过 `@` 前面的数字来自定义顺序

#### 个人博客目录如下图
<img width="390" alt="image" src="https://user-images.githubusercontent.com/10205742/176992908-affe01b6-0a50-488b-bb67-216a75f2a02c.png">


#### 博客导航展示如下图
<img width="407" alt="image" src="https://user-images.githubusercontent.com/10205742/176992913-148a5ba5-bce0-42ed-b09a-9f914556723a.png">

### 部署
> Nginx 反向代理配置文件参考

#### HTTP协议
```
server {
    listen      80;
    listen [::]:80;
    server_name yourhost.com;

    location / {
         proxy_pass  http://127.0.0.1:5006;
         proxy_set_header   Host             $host;
         proxy_set_header   X-Real-IP        $remote_addr;
         proxy_set_header   X-Forwarded-For  $proxy_add_x_forwarded_for;
     }
}
```
#### HTTPS 协议（80端口自动跳转至443）
```
server {
    listen      80;
    listen [::]:80;
    server_name yourhost.com;

    location / {
        rewrite ^ https://$host$request_uri? permanent;
    }
}

server {
    listen          443 ssl;
    server_name     yourhost.com;
    access_log      /var/log/nginx/markdown-blog.access.log main;


    #证书文件名称
    ssl_certificate /etc/nginx/certs/yourhost.com_bundle.crt;
    #私钥文件名称
    ssl_certificate_key /etc/nginx/certs/yourhost.com.key;
    ssl_session_timeout 5m;
    ssl_protocols TLSv1.2 TLSv1.3;
    ssl_ciphers ECDHE-RSA-AES128-GCM-SHA256:HIGH:!aNULL:!MD5:!RC4:!DHE;
    ssl_prefer_server_ciphers on;

    location / {
         proxy_pass  http://127.0.0.1:5006;
         proxy_set_header   Host             $host;
         proxy_set_header   X-Real-IP        $remote_addr;
         proxy_set_header   X-Forwarded-For  $proxy_add_x_forwarded_for;
     }
 }
 ```

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
    > 在 markdown-blog-package 生成当前系统的压缩包 markdown-blog-v0.0.2-darwin-arm64.tar
    ```shell
    make package
    ```

6. 生成 Windows、Linux、Mac 的压缩包
    > 在 markdown-blog-package 生成压缩包 markdown-blog-v0.0.2-darwin-arm64.tar markdown-blog-v0.0.2-linux-arm64.tar.gz markdown-blog-v0.0.2-windows-arm64.zip
    ```shell
    make package-all
    ```
## 授权许可
本项目采用 MIT 开源授权许可证，完整的授权说明已放置在 [LICENSE](https://github.com/gaowei-space/markdown-blog/blob/main/LICENSE) 文件中。
