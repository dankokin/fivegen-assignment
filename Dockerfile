FROM golang:latest

WORKDIR /file_hosting

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 8000

RUN go build -o main .

CMD ./main