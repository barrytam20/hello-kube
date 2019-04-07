FROM node:10-alpine as build-web
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN npm install && npm run build

FROM golang:1.12-alpine as build-webserver
RUN apk add --no-cache git
RUN mkdir /app
COPY /webserver /app
WORKDIR /app
RUN CGO_ENABLED=0 go build

FROM alpine:3.9
RUN apk --update add ca-certificates
RUN mkdir /app
RUN mkdir /app/dist
WORKDIR /app
COPY --from=build-web /app/webserver/dist /app/dist
COPY --from=build-webserver /app/webserver /app
ENTRYPOINT [ "/app/webserver" ]