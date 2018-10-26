dev:
	@echo "=============starting locally============="
	docker-compose up --build

logs:
	docker-compose logs -f

down:
	docker-compose down

test:
	go test ./app/... -cover -coverprofile=coverage.out && go tool cover -html=coverage.out

clean: down
	@echo "=============cleaning up============="
	rm -f api
	docker system prune -f
	docker volume prune -f