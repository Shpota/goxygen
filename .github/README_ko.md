<h1 align="center">
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/gb.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_zh.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/cn.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_ua.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/ua.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_ru.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/ru.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_ko.md">
        <img height="25px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/kr.svg">
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
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_he.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/il.svg">
    </a>
    <br>
    Goxygen
    <a href="https://github.com/Shpota/goxygen/actions?query=workflow%3Abuild">
        <img src="https://github.com/Shpota/goxygen/workflows/build/badge.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/releases">
        <img src="https://img.shields.io/github/v/tag/shpota/goxygen?color=green&label=version">
    </a>
    <a href="https://gitter.im/goxygen/community">
        <img src="https://badges.gitter.im/goxygen/community.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/pulls">
        <img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg">
    </a>
</h1>

<img src="../templates/react.webapp/src/logo.svg" align="right" width="230px" alt="goxygen logo">

**순식간에 GO 그리고 Angular/React/Vue 를 활용한 웹 프로젝트를 생성해보세요**
Goxygen은 사용자가 새로운 프로젝트를 설정하는데 필요한 시간을 아끼는데 중점을 두고있습니다.
사용자를 위한 모든 설정을 포함한 어플리케이션의 뼈대를 생성합니다. 사용자는 생각했던
비지니스 로직을 즉시 행동으로 옮길 수 있습니다. Goxygen 은 Go 코드로 벡엔드를 생성하고, 
그것을 프론트엔드 요소들에 연결 시킨 후, 어플리케이션을 위한 Dockerfile 을 제공하면서
개발과 제작 환경에서 간단한 실행을 위한 docker-compose 파일을 만듭니다.

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

## 사용하는 방법
사용자는 Go 1.11 또는 그 이상의 버젼이 필요합니다.
```go
go get -u github.com/shpota/goxygen
go run github.com/shpota/goxygen init my-app
```
`my-app` 폴더에 프로젝트를 생성합니다.

기본설정으로 React 와 MongoDB를 사용합니다.
다른 프런트엔드 프레임워크와 데이터베이스를 선택하고 싶다면 `--frontend`와 `--db` 플래그를 사용합니다. 예를 들어 아래의 명령은 Vue와 PostgreSQL을 사용하여 프로젝트를 생성합니다:

```go
go run github.com/shpota/goxygen init --frontend vue --db postgres my-app
```

The `--frontend` flag accepts `angular`, `react` and `vue`.
The `--db` flag accepts `mongo`, `mysql` and `postgres`.

생성된 프로젝트는 `docker-compose`를 통해 실행 될 수 있습니다.
```sh
cd my-app
docker-compose up
```
빌드가 끝나면, 어플리케이션은 http://localhost:8080에서 연결이 
가능합니다. 

사용자는 프로젝트의 README 파일에서 더 자세한 사항을 알아볼 수 있습니다

![Showcase](showcase.gif)

## 생성된 프로젝트의 구조 (React/MongoDB 기반 앱)

    my-app
    ├── server                   # Go 프로젝트 파일
    │   ├── db                   # MongoDB 연결
    │   ├── model                # domain objects
    │   ├── web                  # REST APIs, web server
    │   ├── server.go            # 서버의 시작점
    │   └── go.mod               # server 필수 요소
    ├── webapp                    
    │   ├── public               # icons, static 파일, 그리고 index.html
    │   ├── src                       
    │   │   ├── App.js           # 주 React 요소
    │   │   ├── App.css          # App 요소 중점의 스타일(styles)
    │   │   ├── index.js         # application의 시작 기점
    │   │   └── index.css        # global 스타일(styles)
    │   ├── package.json         # front end 필수 요소
    │   ├── .env.development     # 개발환경의 API endpoint
    │   └── .env.production      # 제작환경의 API endpoint
    ├── Dockerfile               # 벡엔드와 프론트엔드를 함께 빌드
    ├── docker-compose.yml       # 제작환경 배포 설명서
    ├── docker-compose-dev.yml   # 개발에 필요한 로컬 MongoDB를 실행
    ├── init-db.js               # 테스트 data 로 MongoDB 컬렉션 생성
    ├── .dockerignore            # Docker 빌드에서 무시할 파일 설정
    ├── .gitignore
    └── README.md                # 생성된 레포지토리 사용 가이드

간략함을 위해 유닛 테스트나 샘플요소는 추가되지 않았습니다.

## 필수요소(dependencies)
Goxygen 프로젝트의 기본적인 구조만을 생성하며 사용자가 구체적인 도구(tool)를 
사용하라고 강요하지 않습니다. 프로젝트에 필요없는 필수요소(dependency)를 
가져오지 않는 이유가 여기에 있습니다. 벡엔드에는 오직
database driver을 사용하며
React나 Vue 프로젝트에선 [axios](https://github.com/axios/axios) 만을 사용합니다.
Angular 프로젝트는 오직 Angular 특화된 라이브러리만을 사용합니다.


## 기여하는 방법

만약 버그를 찾았더라던가 이 프로젝트를 더 나은 방향으로 인도할 수 있는
아이디어를 갖고 있다면
[이슈를 열어주세요](https://github.com/Shpota/goxygen/issues)
그리고 빠른 시일내에 조치하겠습니다. Pull Request를 통해 변경사항을 제안할 수
있습니다. 레포지토리를 Fork하고 더 좋은 방향으로 변경하고 Pull Request 를 
저희에게 보내주세요. 짧은 시일내로 반영하겠습니다. 또한 저희는
[Gitter chat](https://gitter.im/goxygen/community) 에서 모든 사항을 
공유 중 입니다.

### 크레딧
Goxygen 로고는 [Egon Elbre](https://twitter.com/egonelbre) 에 의해 만들어 졌습니다.
