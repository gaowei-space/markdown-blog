<img src="https://user-images.githubusercontent.com/10205742/204092260-007e94c8-f4de-4fdc-94cd-7e1c7f945a91.png">

[![GitHub Checks State](https://img.shields.io/github/checks-status/gaowei-space/meituan-pub-union/main)](https://github.com/gaowei-space/markdown-blog/tree/main)
[![GitHub Issues](https://img.shields.io/github/issues/gaowei-space/markdown-blog?color=yellow)](https://github.com/gaowei-space/markdown-blog/issues)
[![StyleCI](https://github.styleci.io/repos/494669204/shield?branch=main&style=flat)](https://github.styleci.io/repos/494669204?branch=main)
[![GitHub Latest Release](https://img.shields.io/github/v/release/gaowei-space/markdown-blog)](https://github.com/gaowei-space/markdown-blog/releases)
[![GitHub Last Commit](https://img.shields.io/github/last-commit/gaowei-space/markdown-blog?color=blueviolet)](https://github.com/gaowei-space/markdown-blog)
[![Go Report](https://goreportcard.com/badge/github.com/gaowei-space/markdown-blog)](https://goreportcard.com/report/github.com/gaowei-space/markdown-blog)
[![Docker Image Size](https://img.shields.io/docker/image-size/willgao/markdown-blog/latest?color=green)](https://hub.docker.com/repository/docker/willgao/markdown-blog)

[中文](./README.md) | [English](./README_EN.md)

[Markdown-Blog](https://github.com/gaowei-space/markdown-blog) is incredibly fast, easy to use, and converts Markdown formatted text files into beautifully rendered HTML pages.

> Designed with simplicity and ease-of-use in mind, this application allows users to create and publish blog content without the need for complex web development skills. anyone can use this software to create professional-looking blogs with minimal hassle. Whether you're a blogger, writer, or developer, our static blog generator is an excellent choice for creating fast, stylish, and aesthetically pleasing website.

## Case
> [TechMan'Blog](https://blog.willgao.net)

### Web side
Style | Preveiw
--------|------
Dark | <img max-width="600" alt="pc dark" src="https://user-images.githubusercontent.com/10205742/200173152-ca9fa52c-3590-4528-910a-ad42cb278f06.png">
Light | <img max-width="600" alt="pc white" src="https://user-images.githubusercontent.com/10205742/200173231-90f02b72-9e12-4a8b-8dd2-91e1ec4b4ff8.png">

### Mobile
Dark | Light
--------|------
<img max-width="400" alt="mobile dark" src="https://user-images.githubusercontent.com/10205742/201472561-7cd1222e-da0a-4d8c-be11-9f7e9d0851e0.png"> | <img max-width="400" alt="mobile white" src="https://user-images.githubusercontent.com/10205742/201472579-458b902a-dcae-4340-a305-3b54f1679aba.png">

## Supported Platforms
> Windows, Linux, Mac OS

## Updates
* `[v1.1.1]` 2023-05-20
  - Support setting record number
  - Fix cache parameter invalidation
  - Support reading local files, including but not limited to pictures
  - windows environment parsing problem
  - MD folder, read only markdonw files
  - Other known issues fixed

* `[v1.1.0]` 2022-11-26
  - Support for comments
  - Parameter setting, support reading from local file (config.yml)
  - Support for loading favicon.ico
  - Other known issues fixed

* `[v1.0.0]` 2022-11-20
  - Support for **Docker** deployments
  - Packaged static files, optimized for a single application, no more external web directories
  - Other known issues fixed

* `[v0.1.1]` 2022-11-12
  - New third-party analysis statistics configuration, including: Baidu, Google
  - Support configuration page cache time
  - Mobile style optimization
  - Other known issues fixed

* `[v0.0.5]` 2022-11-06
  - Support TOC syntax, when the first line of the file content using `[toc]` will be automatically parsed
  - New bright 🔆 theme, support light and dark switch
  - Other known issues fixed

## Install
### Binary
1. Download [release](https://github.com/gaowei-space/markdown-blog/releases/)

2. Decompress
```sh
tar zxf markdown-blog-v0.0.5-linux-amd64.tar.gz
```

3. create markdown file directory
```sh
cd markdown-blog-linux-amd64
mkdir md
echo "## Hello World" > . /md/home.md
```

4. run
```sh
. /markdown-blog web
```

5. Visit http://127.0.0.1:5006 to see the results

### Docker
1. Download
```sh
docker pull willgao/markdown-blog:latest
```

2. Start
   - online environment
    ```sh
    docker run -dit --rm --name=markdown-blog \
    -p 5006:5006 \
    -v $(pwd)/md:/md -v $(pwd)/cache:/cache \
    willgao/markdown-blog:latest
    ```

   - Development environment
    ```sh
    docker run -dit --rm --name=markdown-blog \
    -p 5006:5006 \
    -v $(pwd)/md:/md -v $(pwd)/cache:/cache \
    willgao/markdown-blog:latest \
    -e dev
    ```

3. Visit http://127.0.0.1:5006 to see the results

4. Other usage
```sh
# View help
docker run -dit --rm --name=markdown-blog \
    -p 5006:5006 \
    -v $(pwd)/md:/md -v $(pwd)/cache:/cache \
    willgao/markdown-blog:latest -h


# Set the title
docker run -dit --rm --name=markdown-blog \
    -p 5006:5006 \
    -v $(pwd)/md:/md -v $(pwd)/cache:/cache \
    willgao/markdown-blog:latest \
    -t "TechMan'Blog"


# Set up Google stats
docker run -dit --rm --name=markdown-blog \
    -p 5006:5006 \
    -v $(pwd)/md:/md -v $(pwd)/cache:/cache \
    willgao/markdown-blog:latest \
    -t "TechMan'Blog" \
    --analyzer-google "De44AJSLDdda"
```

## Use the

### command
- markdown-blog
    - h to view the version
    - web to run the blog service
- markdown-blog web
   - -config FILE                   Load configuration file, default is empty
   - -dir value, -d value           Specify the markdown folder, default: . /md/
   - -title value, -t value         Web service title, default: "Blog"
   - -port value, -p value          Web service port, default: 5006
   - -env value, -e value           Runtime environment, optional: dev,test,prod, default: "prod"
   - -index value, -i value         Set the default home page file name, default is empty
   - -cache value, -c value         Set the page cache time, in minutes, default is 3 minutes
   - --icp value                    ICP record number, default is empty
   - --copyright value              Copyright year, default current year, such as: 2023
   - --fdir value                   The name of the static resource directory under the markdown directory, such as pictures, etc., the default is "public"
   - -analyzer-baidu value          Set Baidu analyzer statistics
   - -analyzer-google value         Set Google analyzer statistics
   - -gitalk.client-id value        Set Gitalk ClientId, default is null
   - -gitalk.client-secret value    Set Gitalk ClientSecret, default is null
   - -gitalk.repo value             Set Gitalk Repo, default is null
   - -gitalk.owner value            Set Gitalk Owner, default is null
   - -gitalk.admin value            Set Gitalk Admin, default is array [gitalk.owner]
   - -gitalk.labels value           Set Gitalk Admin, default is array ["gitalk"].
   - -ignore-file value             Set ignore file, eg: demo.md
   - -ignore-path value             Set ignore folders, eg: demo
   - -h Help

### Run parameters
> Support reading configuration items from configuration file, but specify parameters to take precedence over configuration file at runtime, refer to `config/config.yml.tmp` for configuration content

### Configuration file

1. create a new configuration file `config/config.yml` 2.

2. load the configuration file at startup
- Binary file
```sh
. /markdown-blog web --config . /config/config.yml
```

- Docker
```sh
docker run -dit --rm --name=markdown-blog \
-p 5006:5006 \
-v $(pwd)/md:/md -v $(pwd)/cache:/cache -v $(pwd)/config:/config \
willgao/markdown-blog:latest --config . /config/config.yml
```

### Default home page
> If `index` is not specified at startup, the program defaults to the first file in the navigation as the home page

### Comment plugin
> The comment plugin uses **Gitalk**, please read the plugin instructions before using it [English](https://github.com/gitalk/gitalk/blob/master/readme.md) | [Chinese](https://github.com/gitalk/) gitalk/blob/master/readme-cn.md)

#### Add a new `gitalk` configuration file to be loaded at startup

```yaml
gitalk:
    client-id: "Your github oauth app client-id, required. e.g.: ad549a9d085d7f5736d3"
    client-secret: "Your github oauth app
    client-secret, required. e.g.: 510d1a6bb875fd5031f0d613cd606b1d"
    repo: "The name of the project you intend to use for comments, required. e.g.: blog-issue"
    owner: "Your Github account, required."
    admin:
        - "Your Github account"
    labels:
        - "Custom issue labels, e.g.: gitalk"
```

### Analysis stats
#### Baidu
##### 1. Visit https://tongji.baidu.com to create a site and get the parameter `0952befd5b7da358ad12fae3437515b1` in the official code
```html
<script>
var _hmt = _hmt || [];
(function() {
var hm = document. createElement("script");
hm.src = "https://hm.baidu.com/hm.js?0952befd5b7da358ad12fae3437515b1";
var s = document. getElementsByTagName("script")[0];
s.parentNode.insertBefore(hm, s);
})();
</script>
```
##### 2. Configuration
```sh
./markdown-blog web --analyzer-baidu 0952befd5b7da358ad12fae3437515b1
```

#### Google
##### 1. Visit https://analytics.google.com to create a site and get the parameter `G-MYSMYSMYS` in the official code
```html
<script async="" src="https://www.googletagmanager.com/gtag/js?id=G-MYSMYSMYS"></script>
<script>
     window.dataLayer = window.dataLayer || [];
     function gtag(){dataLayer. push(arguments);}
     gtag('js', new Date());

     gtag('config', 'G-MYSMYSMYS');
</script>
```
##### 2. Configuration
```sh
./markdown-blog web --analyzer-google G-MYSMYSMYS
```

### Title Bar Icon
> By default, read the **favicon.ico** file in the same directory as the program is running

### Navigation Sorting
> The blog navigation is sorted by `dictionary` by default, you can customize the order by the number in front of `@`

#### Personal blog directory as shown below
<img width="390" alt="image" src="https://user-images.githubusercontent.com/10205742/176992908-affe01b6-0a50-488b-bb67-216a75f2a02c.png">

#### Blog navigation display as shown below
<img width="407" alt="image" src="https://user-images.githubusercontent.com/10205742/176992913-148a5ba5-bce0-42ed-b09a-9f914556723a.png">

### deployment
> Nginx reverse proxy configuration file reference

#### HTTP protocol
```nginx
server {
     listen 80;
     listen [::]:80;
     server_name yourhost.com;

     location / {
          proxy_pass http://127.0.0.1:5006;
          proxy_set_header Host $host;
          proxy_set_header X-Real-IP $remote_addr;
          proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
      }
}
```
#### HTTPS protocol (port 80 automatically jumps to 443)
```nginx
server {
     listen 80;
     listen [::]:80;
     server_name yourhost.com;

     location / {
         rewrite ^ https://$host$request_uri? permanent;
     }
}

server {
     listen 443 ssl;
     server_name yourhost.com;
     access_log /var/log/nginx/markdown-blog.access.log main;


     #Certificate file name
     ssl_certificate /etc/nginx/certs/yourhost.com_bundle.crt;
     #Private key file name
     ssl_certificate_key /etc/nginx/certs/yourhost.com.key;
     ssl_session_timeout 5m;
     ssl_protocols TLSv1.2 TLSv1.3;
     ssl_ciphers ECDHE-RSA-AES128-GCM-SHA256:HIGH:!aNULL:!MD5:!RC4:!DHE;
     ssl_prefer_server_ciphers on;

     location / {
          proxy_pass http://127.0.0.1:5006;
          proxy_set_header Host $host;
          proxy_set_header X-Real-IP $remote_addr;
          proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
      }
  }
  ```

## upgrade
1. Download the latest version [release](https://github.com/gaowei-space/markdown-blog/releases/)

2. Stop the program, decompress and replace `markdown-blog`

3. Restart the program

## development
1. Install `Golang` development environment

2. Fork [source code](https://github.com/gaowei-space/gocron)

3. Start the web service

     After running, visit the address [http://localhost:5006](http://localhost:5006), and the API request will be forwarded to `markdown-blog` program
     ```
     make run
     ```

4. Compile

     Generate the compressed package of the current system in the **bin** directory, such as: markdown-blog-v1.1.0-darwin-amd64.tar
     ```
     make
     ```

5. Pack

     Generate the compressed package of the current system in the **package** directory, such as: markdown-blog-v1.1.0-darwin-amd64.tar
     ```
     make package
     ```

6. Generate compressed packages for Windows, Linux, and Mac

     Generate a compressed package in **package**, such as: markdown-blog-v1.1.0-darwin-amd64.tar markdown-blog-v1.1.0-linux-amd64.tar.gz markdown-blog-v1.1.0-windows- amd64.zip
     ```
     make package-all
     ```

## Feedback
- If you encounter problems in the project, you can find answers in [issues](https://github.com/gaowei-space/markdown-blog/issues) or ask questions directly
- If you have any suggestions and ideas, you can start a discussion in [discussions](https://github.com/gaowei-space/markdown-blog/discussions)


## Sponsors
- Thanks to [JetBrains](https://www.jetbrains.com/?from=gaowei-space/markdown-blog) for supporting this project!

<a href="https://www.jetbrains.com/?from=gaowei-space/markdown-blog" target="_blank">
     <img src="https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.png" width="100" height="100">
</a>

## License
This project adopts the MIT open source license, and the complete authorization description has been placed in the [LICENSE](https://github.com/gaowei-space/markdown-blog/blob/main/LICENSE) file.
