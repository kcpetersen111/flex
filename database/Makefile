all: start

build:
	docker build -t flexdb .

start:
	docker run -d --name flexDB -p 5432:5432 flexdb

bs: | build start

clean:
	docker container rm flexDB

kill: 
	docker kill flexDB
