dev:
	go run -tags server -race main.go server

migrate:
	podman-compose up -d postgres
	go run -tags cli -race main.go migrate

secret:
	go run -tags cli -race main.go secret

env:
	podman-compose up -d postgres

down:
	podman-compose down