target: run

run:
	go run ./main.go

build:
	go build -o flex *.go

docker:
	docker build -t flex:latest .

up:
	docker run --name flex -dt -p 8080:8080 flex

down:
	docker rm -f flex

logs:
	docker logs -f flex
