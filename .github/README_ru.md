<h1 align="center">
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/gb.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_ua.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/ua.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_ru.md">
        <img height="25px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/ru.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_zh.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/cn.svg">
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
    <br>
    Goxygen
    <a href="https://github.com/Shpota/goxygen/actions?query=workflow%3Abuild">
        <img src="https://github.com/Shpota/goxygen/workflows/build/badge.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/releases">
        <img src="https://img.shields.io/badge/version-v0.2.1-green">
    </a>
    <a href="https://gitter.im/goxygen/community">
        <img src="https://badges.gitter.im/goxygen/community.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/pulls">
        <img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square">
    </a>
</h1>

<img src="../templates/react.webapp/src/logo.svg" align="right" width="230px" alt="goxygen logo">

**Goxygen генерирует Web проекты на Go, Angular/React/Vue и MongoDB.**

Goxygen сохранит ваше время при создании новых приложений. Он
генерирует базовую структуру Web проекта и позволяет сразу же перейти
к реализации бизнес логики без забот о конфигурации. Goxygen
создает back end код на Go, связывает его с front end компонентами,
добавляет `Dockerfile` и `docker-compose` для удобного запуска
локально и в production среде.

## Использование
Вам нужно иметь Go 1.11 или новее.
```go
go get -u github.com/shpota/goxygen
go run github.com/shpota/goxygen init my-app
```
В результате будет сгенерирован проект в директории `my-app`.

React используется по умолчанию. Вы можете выбрать Angular, React
или Vue, передав `angular`, `react` или `vue` как значение для 
флага`--frontend`. Например:

```go
go run github.com/shpota/goxygen init --frontend vue my-app
```

Сгенерированный проект готов к запуску через `docker-compose`:
```sh
cd my-app
docker-compose up
```
После завершения сборки, приложение будет доступно на
http://localhost:8080.

Больше деталей о приложении и как с ним работать можно найти
в его README.

![Showcase](showcase.gif)

## Структура сгенерированного проекта (на примере React)

    my-app
    ├── server                   # серверная часть приложения (Go)
    │   ├── db                   # коммуникации с MongoDB
    │   ├── model                # доменные объекты
    │   ├── web                  # REST API и Web сервер
    │   ├── server.go            # входная точка серверного кода
    │   └── go.mod               # описание Go модуля и зависимости
    ├── webapp                    
    │   ├── public               # иконки, статические файлы и index.html
    │   ├── src                       
    │   │   ├── App.js           # главный React компонент
    │   │   ├── App.css          # стили главного компонента
    │   │   ├── index.js         # входная точка front end приложения          
    │   │   └── index.css        # глобальные стили
    │   ├── package.json         # front end зависимости
    │   ├── .env.development     # API URL для запуска в development среде
    │   └── .env.production      # API URL для запуска в production среде
    ├── Dockerfile               # собирает front end и back end вместе
    ├── docker-compose.yml       # настройки для запуска в production среде
    ├── docker-compose-dev.yml   # настройки для запуска локальной базы данных
    ├── init-db.js               # наполняет базу данных начальными данными
    ├── .dockerignore            # включает файлы, которые игнорируются при Docker сборке
    ├── .gitignore
    └── README.md                # инструкция по использованию проекта

Юнит тесты и демонстрационные компоненты не включены в структуру
для простоты.

## Зависимости

Goxygen генерирует только базовою структуру проекта и не навязывает
вам использование специфических библиотек или утилит. Сгенерированный
проект имеет только две сторонние зависимости: драйвер для базы
данных и библиотеку для асинхронных REST запросов,
[axios](https://github.com/axios/axios), в React и Vue проектах.

## Как поучаствовать в разработке Goxygen

Если вы нашли ошибку, или хотите предложить улучшения,
[откройте issue](https://github.com/Shpota/goxygen/issues) и мы им
займемся. Также можете предложить свои изменения
[в нашем Gitter чате](https://gitter.im/goxygen/community), или
создав Pull Request. 

### Благодарности
Лого Goxygen создал [Egon Elbre](https://twitter.com/egonelbre).
