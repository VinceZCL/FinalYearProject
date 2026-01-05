# for Echo Backend

dev:
	go -C echo run -tags server -race main.go server

migrate:
	podman-compose up -d postgres
	go -C echo run -tags cli -race echo/main.go migrate

secret:
	go -C echo run -tags cli -race echo/main.go secret

env:
	podman-compose up -d postgres

down:
	podman-compose down

# for Angular Frontend

serve:
	npm --prefix angular run start