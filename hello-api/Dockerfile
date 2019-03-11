FROM golang:1.12-alpine as build
RUN apk add --no-cache git
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 go build

FROM alpine:3.9
RUN apk --update add ca-certificates
COPY --from=build /app/hello-api /
ENTRYPOINT [ "/hello-api" ]