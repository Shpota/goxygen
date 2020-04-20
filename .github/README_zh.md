<h1 align="center">
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/gb.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_zh.md">
        <img height="25px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/cn.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_ua.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/ua.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_ru.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/ru.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_ko.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/kr.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_pt-br.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/br.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_by.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/by.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_fr.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/fr.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_es.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/es.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_jp.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/jp.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_id.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/id.svg">
    </a>
    <br>
    Goxygen
    <a href="https://github.com/Shpota/goxygen/actions?query=workflow%3Abuild">
        <img src="https://github.com/Shpota/goxygen/workflows/build/badge.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/releases">
        <img src="https://img.shields.io/badge/version-v0.3.0-green">
    </a>
    <a href="https://gitter.im/goxygen/community">
        <img src="https://badges.gitter.im/goxygen/community.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/pulls">
        <img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg">
    </a>
</h1>
<img src="../templates/react.webapp/src/logo.svg" align="right" width="230px" alt="goxygen logo">

**分分钟生成一个全栈Web项目(Go，Angular/React/Vue)。**  

Goxygen致力于节省你搭建一个项目的时间。它自动生成一个完全配置好的项目骨架，以方便你可以立即着手实现你的业务逻辑。Goxygen生产后端的Go代码，并将其与前端组件相连，并且为生成的项目提供Dockerfile和docker-compose文件，以便项目在开发环境和生产环境快速运行起来。  

<table>
    <thead>
    <tr align="center">
        <td colspan=4><b>Supported Technologies</b></td>
    </tr>
    </thead>
    <tbody>
    <tr align="center">
        <td align="center">Front End</td>
        <td>Angular</td>
        <td>React</td>
        <td>Vue</td>
    </tr>
    <tr align="center">
        <td>Back End</td>
        <td colspan=3>Go</td>
    </tr>
    <tr align="center">
        <td>Database</td>
        <td>MongoDB</td>
        <td>MySQL</td>
        <td>PostgreSQL</td>
    </tr>
    </tbody>
</table>

# 如何使用
你需要安装Go 1.11或者更新的版本。  
```go
go get -u github.com/shpota/goxygen
go run github.com/shpota/goxygen init my-app
```
这会在你的 `my-app` 目录下生成一个项目。  

By default, it will use React and MongoDB. You can select
a different front end framework and a database using
`--frontend` and `--db` flags. For instance, this command
will create a project with Vue and PostgreSQL:

```go
go run github.com/shpota/goxygen init --frontend vue --db postgres my-app
```

The `--frontend` flag accepts `angular`, `react` and `vue`.
The `--db` flag accepts `mongo`, `mysql` and `postgres`.


这个项目已经可以用`docker-compose`来运行了：  
```sh
cd my-app
docker-compose up
```
build完成之后，就可以在 http://localhost:8080 查看你的Web项目了。    

你可以在生成的项目里的README里查看更多细节信息。  

![Showcase](showcase.gif)

## 生成的项目的结构 (基于React/MongoDB的项目)
    my-app
    ├── server                   # Go项目文件
    │   ├── db                   # MongoDB通信 
    │   ├── model                # 领域对象
    │   ├── web                  # REST APIs, web server
    │   ├── server.go            # 后端入口
    │   └── go.mod               # 后端依赖
    ├── webapp                    
    │   ├── public               # icons, static files, 和 index.html
    │   ├── src                       
    │   │   ├── App.js           # React main组件
    │   │   ├── App.css          # App组件样式
    │   │   ├── index.js         # 前端应用入口
    │   │   └── index.css        # 全局样式
    │   ├── package.json         # 前端依赖
    │   ├── .env.development     # 包含开发环境的API endpoint  
    │   └── .env.production      # 包含生产环境的API endpoint  
    ├── Dockerfile               # 前后端build Dockerfile
    ├── docker-compose.yml       # 生产环境的docker-compose
    ├── docker-compose-dev.yml   # 开发使用的docker-compose
    ├── init-db.js               # 创建一个 MongoDB collection，并写入测试数据
    ├── .dockerignore            # 指明Docker build的时候要忽略的文件
    ├── .gitignore
    └── README.md                # 如何使用生成repo的教程

为了简洁性，诸如单测和样例组件的文件没有显示在这里。   

## 依赖

Goxygen只是为你生成一个项目的基本框架，但不强迫你使用任何特定的工具。它不会为你的项目添加任何不必要的依赖，仅有的依赖只有后端的database driver和前端的[axios](https://github.com/axios/axios)

## 如何贡献

如果你发现了bug或者对于如果改进项目有新的想法，[请提交一个issue](https://github.com/Shpota/goxygen/issues)，我们会尽快解决。  

你也可以直接Pull Request来提交你的修改。Fork这个项目，作出修改之后，发起一个Pull Request，我们会很快review它。    

### 感谢
Goxygen的logo是由[Egon Elbre](https://twitter.com/egonelbre)创作的。  
