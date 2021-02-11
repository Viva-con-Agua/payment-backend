.PHONY: all build db execDB

build:
	docker-compose build --force-rm --no-cache

push:
	docker push vivaconagua/payment-backend

db:
	docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d db

execDB:
	docker-compose exec payment-db mongo mail
