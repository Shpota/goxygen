# project-name

## Environment setup

You need to have [Go](https://go.dev/),
[Node.js](https://nodejs.org/) and
[Docker](https://www.docker.com/)
installed on your computer.

Verify the tools by running the following commands:

```sh
go version
npm --version
docker --version
```

## Start in development mode

In the project directory run the command (you might
need to prepend it with `sudo` depending on your setup):
```sh
docker compose -f docker-compose-dev.yml up
```

This starts a local PostgreSQL database on `localhost:5432`.
The database will be populated with test records from the
[init-db.sql](init-db.sql) file.

Navigate to the `server` folder and start the back end:

```sh
cd server
go run server.go
```
The back end will serve on http://localhost:8080.

Navigate to the `webapp` folder, install dependencies,
and start the front end development server by running:

```sh
cd webapp
npm install
npm start
```
The application will be available on http://localhost:3000.
 
## Start in production mode

Perform:
```sh
docker compose up
```
This will build the application and start it together with
its database. Access the application on http://localhost:8080.
