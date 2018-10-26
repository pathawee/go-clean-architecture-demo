dev:
	@echo "=============starting locally============="
	docker-compose up --build

logs:
	docker-compose logs -f

down:
	docker-compose down

test:
	go test ./app/... -cover -coverpkg=./app/... -coverprofile=coverage.out && go tool cover -html=coverage.out

clean: down
	@echo "=============cleaning up============="
	rm -f app
	docker system prune -f
	docker volume prune -f

format:
	go fmt ./app/...