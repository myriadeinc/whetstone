T = $(TAG)
up:
	docker-compose up

build:
	docker build -f Dockerfile -t myriadeinc/whetstone:$(T) .

dev:
	docker build -f Dockerfile.dev -t myriadeinc/whetstone:dev .