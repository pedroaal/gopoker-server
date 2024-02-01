FROM golang:1.21

LABEL maintainer="Pedro Altamirano <pedroaal@hotmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 3000

CMD ["./main"]