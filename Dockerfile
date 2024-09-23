FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o clean-architecture

EXPOSE 50051 9030 9031

CMD ["./clean-architecture"]