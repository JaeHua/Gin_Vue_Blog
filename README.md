# Gin_Vue_Blog



<p align="center">
  <img src="https://github.com/user-attachments/assets/0b249525-d8e5-47c6-af8f-43f5e856470e" style="max-width=250px max-height=250px" alt="blog (1)">


</p>

<p align="center">
  <img src="https://img.shields.io/badge/Language-Go-blue" alt="Go">
  <img src="https://img.shields.io/badge/Frontend-Vue-green" alt="Vue">
  <img src="https://img.shields.io/badge/Backend-Gin-orange" alt="Gin">
  <img src="https://img.shields.io/badge/Tool-Docker-blue" alt="Docker">
  <img src="https://img.shields.io/badge/Database-MySQL-pink" alt="MySQL">
</p>

## 项目简介

Gin_Vue_Blog 是一个博客项目，采用前后端分离架构。后端使用 Go 语言的 Gin 框架，前端采用 Vue.js。该项目旨在提供一个简单而强大的博客系统，支持基本的博客功能，如文章发布、编辑、删除以及评论功能，现在还在开发中......

## 项目结构
```
GinVueBlog
│  
├─backend
│  │  backend.exe
│  │  go.mod
│  │  go.sum
│  │  main.go
│  │  
│  ├─api
│  │  └─v1
│  │          article.go
│  │          category.go
│  │          email.go
│  │          login.go
│  │          profile.go
│  │          upload.go
│  │          user.go
│  │          
│  ├─config
│  │      config.ini
│  │      
│  ├─middleware
│  │      cors.go
│  │      jwt.go
│  │      log.go
│  │      
│  ├─model
│  │      Article.go
│  │      Category.go
│  │      db.go
│  │      Email.go
│  │      Profile.go
│  │      redis.go
│  │      Upload.go
│  │      User.go
│  │      
│  ├─routes
│  │      route.go
│  │      
│  └─utils
│      │  setting.go
│      │  
│      ├─email
│      │      email.go
│      │      
│      ├─errmsg
│      │      errmsg.go
│      │      
│      └─validator
│              validator.go
│              
└─frontend
    ├─admin
    │  ├─public
    │  │      favicon.ico
    │  │      index.html
    │  └─src
    │      │  App.vue
    │      │  main.js
    │      │  
    │      ├─assets        
    │      ├─components
    │      │  ├─admin
    │      │  ├─article
    │      │  ├─category  
    │      │  ├─editor   
    │      │  └─user      
    │      ├─plugin    
    │      ├─router
    │      │      index.js  
    │      └─views
    │              Admin.vue
    │              Login.vue
    │              
    └─front
        ├─public
        │  │  favicon.ico
        │  │  index.html
        │  │  
        │  └─markdown
        │          markdown.css
        │          
        └─src
            │  App.vue
            │  main.js
            │  
            ├─assets
            │  └─css
            │
            ├─store
            │      index.js
            ├─components
            │      404.vue
            │      Account.vue
            │      ArtList.vue
            │      Category.vue
            │      Details.vue
            │      Footer.vue
            │      LikedArticles.vue
            │      Nav.vue
            │      Register.vue
            │      Result.vue
            │      SavedArticles.vue
            │      TopBar.vue
            │      UserSettings.vue
            ├─plugins
            │      catalog.js
            │      http.js
            │      vuetify.js
            │      
            ├─router  
            └─views
                    

```
## 后端技术
- Golang 
- Gin Web 框架
- Gorm ORM
- JWT 鉴权
- Mysql 数据库
- Redis 缓存
- Validator 验证
- 云存储服务（图床）
## 前端技术 

- Vue3
- Vuetify UI
- Ant Design Vue
- Axios 网络请求
- Vue Router 路由
- Vuex 状态管理
- so on

## 教程

### Docker安装
- 更新软件包列表
```bash
sudo apt update
sudo apt upgrade
```
- 安装依赖
```bash
apt-get install ca-certificates curl gnupg lsb-release
```
- 添加官方秘钥
```
curl -fsSL http://mirrors.aliyun.com/docker-ce/linux/ubuntu/gpg | sudo apt-key add -
```
- 添加docker软件源
```
sudo add-apt-repository "deb [arch=amd64] http://mirrors.aliyun.com/docker-ce/linux/ubuntu $(lsb_release -cs) stable"
```
- 安装docker
```
apt-get install docker-ce docker-ce-cli containerd.io
```

### Pointer配置
- 拉取

```
docker pull portainer/portainer
```
- 运行
```
docker run -p 9000:9000 -p 8000:8000 --name portainer \--restart=always \-v /var/run/docker.sock:/var/run/docker.sock \-v /mydata/portainer/data:/data \-d portainer/portainer
```

### Mysql配置

- 拉取镜像（版本5.7）

```Bash
docker pull mysql:5.7
```

- 创建对应文件夹

```Bash
mkdir -p /home/mysql/conf
mkdir -p /home/mysql/data
```

- 授权文件夹

```Bash
chmod 777 /home/mysql/conf
chmod 777 /home/mysql/data
```

- 运行容器

```Bash
docker run --name mysql \
  -v /home/mysql/conf:/etc/mysql/conf.d \
  -v /home/mysql/data:/var/lib/mysql \
  -e MYSQL_ROOT_PASSWORD=123456 \
  -p 3306:3306 \
  --restart=always \
  -d mysql:5.7
```

## Redis部署

- 拉取镜像(版本7.0)

```Bash
docker pull redis:7.0
```

- 创建对应文件夹

```Bash
mkdir -p /usr/local/docker/redis/conf
mkdir -p /usr/local/docker/redis/data
```

- 授权文件夹

```Bash
chmod 777 /usr/local/docker/redis/conf
chmod 777 /usr/local/docker/redis/data
```
- 复制一份配置文件
  根据对应版本复制到redis.conf
- 运行容器

```Bash
docker run \
    -d \
    -p 6379:6379 \
    --name Redis \
    --restart=always \
    -v /usr/local/docker/redis/data:/data \
    -v /usr/local/docker/redis/conf/redis.conf:/etc/redis \
    redis
```
### Gin部署

- 创建对应文件夹(代码打包放在这)

```Bash
mkdir /home/blog/
```

- 构建镜像

```Bash
docker build -t goblog:1.0 .
```

- 运行容器

```Bash
docker run -d -p 80:80 --name MyBlog goblog:1.0
```

