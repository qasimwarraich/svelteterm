
FROM golang:1.18

WORKDIR /fiber-app

COPY . ./

RUN go mod download && go mod verify

RUN go build -v -o ./server ./src/main.go

EXPOSE 8080

CMD ./server
