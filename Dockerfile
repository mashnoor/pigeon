FROM golang:1.17.8-alpine3.15 as build

WORKDIR app/
RUN apk add git
COPY go.mod .
COPY go.sum .
RUN ["go", "mod", "tidy"]
COPY . .

RUN ["go", "build", "main/runner.go"]



FROM alpine:latest
WORKDIR /app
COPY --from=build /go/app/runner .
CMD ["./runner"]