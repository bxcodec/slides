build:
	docker build . -t devfest
run:
	@docker-compose build --no-cache
	@docker-compose up -d --force-recreate
stop:
	@docker-compose down

