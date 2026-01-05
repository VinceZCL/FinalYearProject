# for Dependencies

go-deps:
	cd ./echo && go mod tidy

npm-deps:
	npm install
	cd ./angular && npm install

deps: go-deps npm-deps

# for Echo Backend

dev: go-deps
	go -C echo run -tags server -race main.go server

migrate: go-deps
	podman-compose up -d postgres
	go -C echo run -tags cli -race main.go migrate

secret: go-deps
	go -C echo run -tags cli -race main.go secret

env:
	podman-compose up -d postgres

down:
	podman-compose down

# for Angular Frontend

serve: npm-deps
	npm --prefix angular run start