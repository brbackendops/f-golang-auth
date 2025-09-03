FROM golang:1.25-alpine

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o app .

EXPOSE 3000

CMD ["sh","-c","./app"]