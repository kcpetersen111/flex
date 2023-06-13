from golang:1.20

workdir /app

copy . .

run go build *.go

expose 8080

cmd ["./main"]
