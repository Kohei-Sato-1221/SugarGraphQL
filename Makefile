run-db:
	docker-compose up -d

seed:
	cd database && go run ./seed.go