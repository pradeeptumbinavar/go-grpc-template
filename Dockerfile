FROM golang:1.21-alpine

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN make proto build

EXPOSE 50051
CMD ["./bin/server"]
