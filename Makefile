.DEFAULT_GOAL := usage

build:
	docker-compose build

up:
	docker-compose up

stop:
	docker-compose stop

down:
	docker-compose down

clean-all:
	docker-compose down --rmi all

connect-db:
	docker exec -it golang-crud-sample-db bash

usage:
	@echo usage: [build, up, stop, down, clean-all, connect-db]