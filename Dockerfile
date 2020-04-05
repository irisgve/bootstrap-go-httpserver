FROM golang:1.14

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o bootstrap-go-httpserver ./cmd/socialkittly_streamer/main.go
CMD ["./bootstrap-go-httpserver"]