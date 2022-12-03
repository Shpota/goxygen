FROM node:16.18.1-alpine3.16 AS JS_BUILD
COPY webapp /webapp
WORKDIR webapp
RUN npm install && npm run build

FROM golang:1.19.0-alpine3.16 AS GO_BUILD
RUN apk update && apk add build-base
COPY server /server
WORKDIR /server
RUN go build -o /go/bin/server

FROM alpine:3.16.3
COPY --from=JS_BUILD /webapp/build* ./webapp/
COPY --from=GO_BUILD /go/bin/server ./
CMD ./server
