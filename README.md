# goxygen [![build](https://github.com/Shpota/goxygen/workflows/build/badge.svg)](https://github.com/Shpota/goxygen/actions?query=workflow%3Abuild) [![version](https://img.shields.io/badge/version-v0.1.0-green)](https://github.com/Shpota/goxygen/releases) [![](https://badges.gitter.im/goxygen/community.svg)](https://gitter.im/goxygen/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge) [![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](https://github.com/Shpota/goxygen/pulls)

<img src="./templates/webapp/src/logo.svg" align="right" width="230px" alt="goxygen logo">

**Generate a Full Stack Web project with Go, React, and MongoDB in seconds.**

Goxygen aims at saving your time while setting up a new project. It
creates a skeleton of an application with all configuration done for
you. You can start implementing your business logic straight away.
Goxygen generates back end Go code, connects it with front end React
components, provides a Dockerfile for the application and creates
docker-compose files for convenient run in development and production
environments.

## How to use
You need to have Go 1.11 or newer on your machine.
```go
git clone https://github.com/mattb2401/goxygen.git
cd goxygen/
go install
goxygen init -n my-app
```
This generates a project in `my-app` folder. 

The generated project is ready to run with `docker-compose`:
```sh
cd my-app
docker-compose up
```
After the build is completed, the application is accessible
on http://localhost:8080.

You can find more details on how to work with the generated
project in its readme file. 

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

Files such as unit tests or sample components are not included here
for simplicity.

## Dependencies

Goxygen generates a basic structure of a project and doesn't force you
to use a specific set of tools. That's why it doesn't bring unneeded
dependencies to your project. The only two dependencies are
[mongo-go-driver](https://github.com/mongodb/mongo-go-driver) on the
back end side and [axios](https://github.com/axios/axios) on the front
end side.

## How to contribute

If you found a bug or have an idea on how to improve the project
[open an issue](https://github.com/Shpota/goxygen/issues)
and we will fix it as soon as possible.

You can also propose your changes via a Pull Request. Fork the
repository, make changes, send us a pull request and we'll
review it shortly.

### Credits
Goxygen's logo was created by [Egon Elbre](https://twitter.com/egonelbre).

