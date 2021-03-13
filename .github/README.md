<h1 align="center">
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README.md">
        <img height="25px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/gb.svg">
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

**Generate a Web project with Go and Angular, React or Vue.**

Goxygen aims at saving your time while setting up a new project. It
creates a skeleton of an application with all configuration done for
you. You can start implementing your business logic straight away.
Goxygen generates back end Go code, connects it with front end
components, provides a Dockerfile for the application and creates
docker-compose files for convenient run in development and production
environments.

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

## How to use
You need to have Go 1.11 or newer on your machine.
```go
go get -u github.com/shpota/goxygen
go run github.com/shpota/goxygen init my-app
```
This generates a project in `my-app` folder. 

By default, it will use React and MongoDB. You can select
a different front end framework and a database using
`--frontend` and `--db` flags. For instance, this command
will create a project with Vue and PostgreSQL:

```go
go run github.com/shpota/goxygen init --frontend vue --db postgres my-app
```

The `--frontend` flag accepts `angular`, `react` and `vue`.
The `--db` flag accepts `mongo`, `mysql` and `postgres`.

The generated project is ready to run with `docker-compose`:
```sh
cd my-app
docker-compose up
```
After the build is completed, the application is accessible
on http://localhost:8080.

You can find more details on how to work with the generated
project in its README file. 

![Showcase](showcase.gif)

## Structure of a generated project (React/MongoDB example)

    my-app
    ├── server                   # Go project files
    │   ├── db                   # MongoDB communications
    │   ├── model                # domain objects
    │   ├── web                  # REST APIs, web server
    │   ├── server.go            # the starting point of the server
    │   └── go.mod               # server dependencies
    ├── webapp                    
    │   ├── public               # icons, static files, and index.html
    │   ├── src                       
    │   │   ├── App.js           # the main React component
    │   │   ├── App.css          # App component-specific styles
    │   │   ├── index.js         # the entry point of the application          
    │   │   └── index.css        # global styles
    │   ├── package.json         # front end dependencies
    │   ├── .env.development     # holds API endpoint for dev environment
    │   └── .env.production      # API endpoint for prod environment
    ├── Dockerfile               # builds back end and front end together
    ├── docker-compose.yml       # prod environment deployment descriptor
    ├── docker-compose-dev.yml   # runs local MongoDB for development needs
    ├── init-db.js               # creates a MongoDB collection with test data
    ├── .dockerignore            # specifies files ignored in Docker builds
    ├── .gitignore
    └── README.md                # guide on how to use the generated repo

Files such as unit tests or sample components are not included here
for simplicity.

## Dependencies

Goxygen generates a basic structure of a project and doesn't force you
to use a specific set of tools. That's why it doesn't bring unneeded
dependencies to your project. It uses only a database driver on the
back end side and [axios](https://github.com/axios/axios) in React
and Vue projects. Angular projects use only Angular specific libraries.

## How to contribute

If you found a bug or have an idea on how to improve the project
[open an issue](https://github.com/Shpota/goxygen/issues)
and we will fix it as soon as possible. You can also propose your
changes via a Pull Request. Fork the repository, make changes, send
us a pull request and we'll review it shortly. We also have a
[Gitter chat](https://gitter.im/goxygen/community) where we discuss
all the changes.

## Troubleshooting

### I'm getting an error
If when running the `init` command you are getting the following error:
```
$ go run github.com/shpota/goxygen init my-app
no required module provides package github.com/shpota/goxygen: working directory is not part of a module
```
This means you are using Go version 1.13 or later. To make this command work you need to set your `GO111MODULE` environment variable to `auto`, like this:
```
export GO111MODULE=auto
```

## Credits
Goxygen's logo was created by [Egon Elbre](https://twitter.com/egonelbre).
