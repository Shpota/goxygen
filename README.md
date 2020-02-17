# goxygen [![build](https://github.com/Shpota/goxygen/workflows/build/badge.svg)](https://github.com/Shpota/goxygen/actions?query=workflow%3Abuild) [![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](https://github.com/Shpota/goxygen/pulls)

<img src="./templates/webapp/src/logo.svg" align="right" width="230px" alt="goxygen logo">

**Generate a Full Stack Web project with Go, React, and MongoDB in seconds.**

Goxygen aims at saving your time while setting up a new project. It
creates a skeleton of an application with all configuration done for
you. You can start implementing your business logic strait away.
Goxygen generates back end Go code, connects it with front end React
components, provides a Dockerfile for the application and creates
docker-compose files for convenient run in development and production
environments.

## How to use
```go
go get -u github.com/shpota/goxygen
go run github.com/shpota/goxygen init my-app
```
This creates a new `my-app` project folder. The generated project
is ready to run with `docker-compose`:
```sh
cd my-app
docker-compose up
```
After the build is completed, the application is accessible
on http://localhost:8080.

![Showcase](.github/showcase.gif)

## Structure of a generated project

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

Files such as unit tests or sample components are not reflected here
for simplicity.

## Toolchain

You need to have [Go](https://golang.org/) installed to generate
a project with Goxygen.

In order to work with the generated project you need to have
[Go](https://golang.org/), [Node.js](https://nodejs.org/),
[Docker](https://www.docker.com/),
and [Docker Compose](https://docs.docker.com/compose/)
(comes pre-installed with Docker on Mac and Windows).

Verify the toolchain by running the following commands:

```sh
go version
npm --version
docker --version
docker-compose --version
```

## How to contribute

If you found a bug or have an idea on how to improve the project
[open an issue](https://github.com/Shpota/goxygen/issues)
and we will fix it as soon as possible.

You can also propose your changes via a Pull Request. Fork the
repository, make changes, send us a pull request and we'll
review it shortly.

### Logo
Goxygen's logo was created by [Egon Elbre](https://twitter.com/egonelbre).

