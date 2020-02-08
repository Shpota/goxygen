# goxygen [![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](https://github.com/Shpota/goxygen/pulls)

<img src="./templates/webapp/src/logo.svg" align="right" height="150px" alt="goxygen logo">

**Generate a Full Stack Web project with Go, React, and MongoDB in seconds.**

Goxygen aims at saving your time while setting up a new project. It
creates a skeleton of an application with all configuration done for
you: the back end Go code with MongoDB, the front end React components
that communicate the back end via REST and Docker scripts to glue all
together.

## How to use
```go
go get -u github.com/shpota/goxygen
go run github.com/shpota/goxygen init my-dream-app
```
This creates a new `my-dream-app` project folder. The generated project
is ready to run with `docker-compose`:
```sh
cd my-dream-app
docker-compose up
```
After the build is completed, the application is accessible
on http://localhost:8080.

## How to contribute

If you found a bug or have an idea on how to improve the project
[open an issue](https://github.com/Shpota/goxygen/issues)
and we will fix it as soon as possible.

You can also propose your changes via a Pull Request. Fork the
repository, make changes, send us a pull request and we'll
review it shortly.

## Logo
Goxygen's logo was created by [Egon Elbre](https://twitter.com/egonelbre).
