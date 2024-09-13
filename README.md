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
卷 Data 的文件夹 PATH 列表
卷序列号为 8A81-E686
GinVueBlog
│  
├─backend
│  │  backend.exe
│  │  go.mod
│  │  go.sum
│  │  main.go
│  │  
│  ├─.idea
│  │      .gitignore
│  │      backend.iml
│  │      modules.xml
│  │      vcs.xml
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
│  ├─static
│  │  ├─admin
│  │  │  │  favicon.ico
│  │  │  │  index.html
│  │  │  │  
│  │  │  └─static
│  │  │      ├─css   
│  │  │      ├─fonts  
│  │  │      ├─img    
│  │  │      └─js       
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
    │  │      
    │  └─src
    │      │  App.vue
    │      │  main.js
    │      │  
    │      ├─assets
    │      │  │  logo.png
    │      │  │  
    │      │  └─css
    │      │          style.css
    │      │          
    │      ├─components
    │      │  ├─admin
    │      │  │      Footer.vue
    │      │  │      Header.vue
    │      │  │      Index.vue
    │      │  │      Nav.vue
    │      │  │      
    │      │  ├─article
    │      │  │      AddArt.vue
    │      │  │      ArtList.vue
    │      │  │      
    │      │  ├─category
    │      │  │      CateList.vue
    │      │  │      
    │      │  ├─editor
    │      │  │      index.vue
    │      │  │      
    │      │  └─user
    │      │          Profile.vue
    │      │          UserList.vue
    │      │          
    │      ├─plugin
    │      │      antui.js
    │      │      axios.js
    │      │      
    │      ├─router
    │      │      index.js
    │      │      
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
            │  │  404.jpg
            │  │  biaozhi1.png
            │  │  imgBg.jpg
            │  │  logo.png
            │  │  logo.svg
            │  │  
            │  └─css
            │          style.css
            │          
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
            │      
            ├─plugins
            │      catalog.js
            │      http.js
            │      vuetify.js
            │      
            ├─router
            │      index.js
            │      
            └─views
                    Home.vue
                    

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
- so on
