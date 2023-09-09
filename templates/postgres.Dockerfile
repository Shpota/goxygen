FROM node:16.20-alpine3.18 AS JS_BUILD
COPY webapp /webapp
WORKDIR /webapp
RUN npm install && npm run build

FROM golang:1.21.1-alpine3.18 AS GO_BUILD
COPY server /server
WORKDIR /server
RUN go build -o /go/bin/server

FROM alpine:3.18.3
COPY --from=JS_BUILD /webapp/build* ./webapp/
COPY --from=GO_BUILD /go/bin/server ./
CMD ./server
