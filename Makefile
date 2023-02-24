# start the environment of FreeCar
.PHONY: start
start:
	docker-compose up -d

# stop the environment of FreeCar
.PHONY: stop
stop:
	docker-compose down