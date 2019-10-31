docker:
	@docker build . -t asia.gcr.io/bxcodec/slides:latest
run:
	@docker-compose build --no-cache
	@docker-compose up -d --force-recreate
stop:
	@docker-compose down

