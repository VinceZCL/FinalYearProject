# for Dependencies

go-deps:
	cd ./server && go mod tidy

npm-deps:
	npm install
	cd ./client && npm install

deps: go-deps npm-deps

build-client:
	npm --prefix client run build -- --configuration development

# for Echo Backend

dev: go-deps
	go -C server run -tags server -race main.go server

migrate: go-deps env
	go -C server run -tags cli -race main.go migrate

secret: go-deps
	go -C server run -tags cli -race main.go secret

default: go-deps
	go -C server run -tags default -race main.go default

build: deps build-client
	go -C server build -tags server -o ../out/fyp-scrum

# for Postgres Database

env:
	podman-compose up -d postgres

down:
	podman-compose down

# for Angular Frontend

serve: npm-deps
	npm --prefix client run start

# for liquibase database
update:
	cd postgres && liquibase update