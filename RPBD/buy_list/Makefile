OS=linux
ARCH=amd64

run:
	docker-compose up -d

run-build:
	docker-compose up --build -d

db-run:
	docker-compose --profile db up -d

app-run:
	docker-compose --profile app up -d

app-run-build:
	docker-compose --profile app up --build -d

stop:
	docker-compose down -d