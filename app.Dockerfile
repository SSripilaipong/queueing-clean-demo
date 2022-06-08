FROM golang:1.18-buster

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY app app
COPY base base
COPY domain domain
COPY implementation implementation
COPY rest rest
COPY outbox outbox
COPY worker worker
COPY main.go .
RUN go build -o /service

EXPOSE 8080

CMD [ "/service" ]