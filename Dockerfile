FROM golang:1.23-alpine3.20

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN go build -o ltf  cmd/main.go

EXPOSE 5000

CMD [ "./ltf" ]