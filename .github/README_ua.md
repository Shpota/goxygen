<h1 align="center">
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/gb.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_zh.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/cn.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_ua.md">
        <img height="25px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/ua.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_ko.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/kr.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_pt-br.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/br.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_by.md">
        <img height="20px" src="https://raw.githubusercontent.com/Shpota/goxygen/main/.github/flag-by.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_fr.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/fr.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_es.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/es.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_jp.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/jp.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_id.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/id.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_he.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/il.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/main/.github/README_tr.md">
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

**Згенеруйте Web проєкт з використанням Go та сучасних SPA фреймворків.**

Goxygen зберігає ваш час при створенні нових проєктів. Він
генерує базову структуру проєкту і дозволяє вам відразу ж перейти до
реалізації бізнес логіки без турбот про налаштування. Goxygen створює
back end код на Go, зв'язує його з front end компонентами, додає
`Dockerfile` та `docker-compose` для зручного запуску локально та в
production середовищі.

<table>
    <thead>
    <tr align="center">
        <td colspan=4><b>Доступні технології</b></td>
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
        <td>База даних</td>
        <td>MongoDB</td>
        <td>MySQL</td>
        <td>PostgreSQL</td>
    </tr>
    </tbody>
</table>

## Вимоги
Ви повинні мати Go 1.11 або новішу версію на вашому комп'ютері.

Змінна оточення `GO111MODULE` має бути встановлена в `auto`
для того, щоб згенерувати проєкт. Це значення за замовчанням
для версій Go до 1.15. Для 1.16 її треба встановити явно:
```
export GO111MODULE=auto
```

## Як користуватися
```go
go get -u github.com/shpota/goxygen
go run github.com/shpota/goxygen init my-app
```
В результаті буде згенеровано проєкт в папці `my-app`. 

React і MongoDB використовуються за замовчанням. Ви можете
обрати інший front end фреймворк та іншу базу даних
використовуючи параметри `--frontend` та `--db` відповідно. 
Наступна команда генерує проєкт з Vue та PostgreSQL:

```go
go run github.com/shpota/goxygen init --frontend vue --db postgres my-app
```

Параметер `--frontend` може бути встановленим в `angular`, `react` та `vue`.
Параметер `--db` може бути встановленим в  `mongo`, `mysql` та `postgres`.

Згенерований проєкт готовий до запуску з `docker-compose`:
```sh
cd my-app
docker compose up
```
Після завершення збірки, застосунок буде доступний на 
http://localhost:8080.

Більше деталей про те, як працювати зі згенерованим проєктом,
можна знайти в його README. 

![Showcase](showcase.gif)

## Структура згенерованого проєкту (на пиркладі React/MongoDB)

    my-app
    ├── server                   # серверна частина застосунку (Go)
    │   ├── db                   # комунікації MongoDB
    │   ├── model                # доменні об'єкти
    │   ├── web                  # REST API та Web сервер
    │   ├── server.go            # вхідна точка серверного коду
    │   └── go.mod               # опис Go модуля та залежності
    ├── webapp                    
    │   ├── public               # іконки, статичні файли та index.html
    │   ├── src                       
    │   │   ├── App.js           # головний React компонент
    │   │   ├── App.css          # стилі головного компоненту
    │   │   ├── index.js         # вхідна точка front end застосунку          
    │   │   └── index.css        # глобальні стилі
    │   ├── package.json         # front end залежності
    │   ├── .env.development     # API URL для запуску в development середовищі
    │   └── .env.production      # API URL для запуску в production середовищі
    ├── Dockerfile               # збирає front end та back end разом
    ├── docker-compose.yml       # налаштування для запуску в production середовищі
    ├── docker-compose-dev.yml   # налаштування запуску локальної MongoDB
    ├── init-db.js               # наповнює базу даних початковими даними
    ├── .dockerignore            # перерахунок файлів, які ігноруються при збірці Docker
    ├── .gitignore
    └── README.md                # інструкція по використанню проєкту

Юніт тести та демонстраційні компоненти не відображені в
структурі задля простоти.

## Залежності

Goxygen генерує лише базову структуру проєкту і не нав'язує вам
використання специфічних бібліотек чи утиліт. Згенерований проєкт
використовує лише драйвер бази даних в back end частині та
[axios](https://github.com/axios/axios) для асинхронних REST запитів
(для React i Vue проєктів).

## Як долучитися до проєкту

Якщо ви знайшли помилку або маєте ідеї щодо покращення проєкту,
[відкрийте issue](https://github.com/Shpota/goxygen/issues)
і ми ним займемся. Якщо хочете запропонувати зміни - можете 
обговорити їх [в нашому Gitter чаті](https://gitter.im/goxygen/community)
або створити Pull Request. 

### Подяки
Лого Goxygen створив [Egon Elbre](https://twitter.com/egonelbre).
