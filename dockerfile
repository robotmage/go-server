FROM golang
COPY . /app
WORKDIR /app
CMD go run /app/main.go
