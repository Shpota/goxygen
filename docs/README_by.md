<h1 align="center">
    <a href="https://github.com/Shpota/goxygen/tree/main/docs/README.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/gb.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/docs/README_zh.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/cn.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/docs/README_ua.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/ua.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/docs/README_ko.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/kr.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/docs/README_pt-br.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/br.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/docs/README_by.md">
        <img height="25px" src="https://raw.githubusercontent.com/Shpota/goxygen/main/docs/flag-by.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/docs/README_fr.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/fr.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/docs/README_es.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/es.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/docs/README_jp.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/jp.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/docs/README_id.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/id.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/docs/README_he.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/il.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/docs/README_tr.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/tr.svg">
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

<img src="../templates/vue.webapp/src/assets/logo.svg" align="right" width="230px" alt="goxygen logo">

**Goxygen генеруе Web праекты на Go і сучасных SPA фреймворках.**

Goxygen захавае ваш час пры стварэнні новых дадаткаў. Ён генеруе 
базавую структуру Web праекта і дазваляе адразу ж перайсці да 
рэалізацыі бізнес логікі без клопатаў аб канфігурацыі. Goxygen
стварае back end код на Go, звязвае яго з front end кампанентамі,
дадае `Dockerfile` і `docker-compose` для зручнага запуску
лакальна і ў production асяроддзі.

<table>
    <thead>
    <tr align="center">
        <td colspan=4><b>Падтрымліваюцца тэхналогіі</b></td>
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
        <td>База даных</td>
        <td>MongoDB</td>
        <td>MySQL</td>
        <td>PostgreSQL</td>
    </tr>
    </tbody>
</table>

## Requirements
Вам трэба мець Go 1.16 або навей вэрсію.

The `GO111MODULE` environment variable has to be set to `auto`
for the generation logic to work. It is a default for Go
versions up to 1.15. For Go 1.16, you need to set it explicitly:
```
export GO111MODULE=auto
```

## Выкарыстанне

```go
go run github.com/shpota/goxygen@latest init my-app
```

У выніку будзе згенераваны праект у дырэкторыі `my-app`.

React і MongoDB  выкарыстоўваецца па змаўчанню. Вы 
можаце выбраць розныя front end фреймворк і розныя 
базу даных з дапамогай `--frontend` і `--db` сцягі. 
Напрыклад, гэтая каманда створыць праект з Vue і PostgreSQL:

```go
go run github.com/shpota/goxygen@latest init --frontend vue --db postgres my-app
```

Сцяг `--frontend` прымае `angular`, `react` або ` vue`.
Сцяг `--db` прымае `mongo`, `mysql` або `postgres`.

Сфарміраваны праект гатовы да запуску з `docker-compose`: 
```sh
cd my-app
docker compose up
```
Пасля завяршэння зборкі, дадатак будзе даступны на
http://localhost:8080.

Больш дэталяў пра дадатак і як з ім працаваць можна знайсці ў 
яго README.

![Showcase](showcase.gif)

## Структура генераванага праекта (прыкладанне на аснове React/MongoDB)

    my-app
    ├── server                   # серверная частка дадатку (Go)
    │   ├── db                   # камунікацыі з MongoDB
    │   ├── model                # даменныя аб'екты
    │   ├── web                  # REST API і Web сервер
    │   ├── server.go            # ўваходная кропка сервернага кода
    │   └── go.mod               # апісанне Go модуля і залежнасці
    ├── webapp                    
    │   ├── public               # іконкі, статычныя файлы і index.html
    │   ├── src                       
    │   │   ├── App.js           # галоўны React кампанент
    │   │   ├── App.css          # стылі галоўнага кампанента
    │   │   ├── index.js         # ўходная кропка front end дадатку          
    │   │   └── index.css        # глабальныя стылі
    │   ├── package.json         # front end залежнасці
    │   ├── .env.development     # API URL для запуску ў development асяроддзі
    │   └── .env.production      # API URL для запуску ў production асяроддзі
    ├── Dockerfile               # збірае front end і back end разам
    ├── docker-compose.yml       # настройкі для запуску ў production асяроддзі
    ├── docker-compose-dev.yml   # настройкі для запуску лакальнай базы даных
    ├── init-db.js               # напаўняе базу даных пачатковымі данымі
    ├── .dockerignore            # вызначае файлы, якія ігнаруюцца ў Docker зборцы
    ├── .gitignore
    └── README.md                # інструкцыя па выкарыстанні праекта

Юніт тэсты і дэманстрацыйныя кампаненты не ўключаны ў структуру для прастаты.

## Залежнасці

Goxygen генеруе толькі базавая структуру праекта і не навязвае
вам выкарыстанне спецыфічных бібліятэк або утыліт. Згенераваны
праект будзе мець толькі два старонніх залежнасці: драйвер для базы
даных і бібліятэку для асінхронных REST запытаў,
[axios](https://github.com/axios/axios), для праектаў React і Vue.

## Як далучыцца да праекта

Калі вы знайшлі памылку, ці хочаце прапанаваць паляпшэння,
[адкрыйце issue](https://github.com/Shpota/goxygen/issues) і мы з гэтым 
разбярэмся. Таксама можаце прапанаваць свае змены
[у нашым Gitter чаце](https://gitter.im/goxygen/community), або
стварыўшы Pull Request. 

### Падзякі

Лагатып Goxygen стварыў [Egon Elbre](https://twitter.com/egonelbre).
