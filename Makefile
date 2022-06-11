dev:
	docker-compose -f appOnly.docker-compose.yml up --build

start:
	docker-compose up --build --detach

stop:
	docker-compose down

restart: stop start