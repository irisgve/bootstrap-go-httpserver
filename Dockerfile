FROM golang:1.16

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o bootstrap-go-httpserver ./cmd/bootstrap-go-httpserver/main.go
CMD ["./bootstrap-go-httpserver"]