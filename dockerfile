FROM golang:1.17

WORKDIR /go/src/app
COPY . .
ENV GIN_MODE=release
EXPOSE 8080
RUN go get -d -v ./...
RUN go install -v ./...

CMD [ "go", "run", "main.go" ]