dev:
	go run -tags server -race main.go server

migrate:
	go run -tags cli -race main.go migrate

env:
	podman-compose up -d postgres

down:
	podman-compose down