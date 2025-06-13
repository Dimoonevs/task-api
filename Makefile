task-api-linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/task-api-linux-amd64 ./cmd/main.go

task-api-run-docker: down task-api-linux
	docker compose up -d --build

down:
	docker compose down

clean:
	docker compose down -v --rmi all --remove-orphans

logs:
	docker compose logs -f