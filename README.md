# Final Year Project

**Title**: *A Web-Based Scrum Management Tool for Daily Check-Ins with Smart Suggestions and Jira Ticket Mapping*

## Components

* **PostgreSQL** Database
* **Go Echo** Backend
* **TypeScript Angular** Frontend

## Prerequisites

* **Go**
* **npm**
* **Podman** with **Podman Compose**

## Project Structure

```text
.
├── Makefile
├── README.md
├── angular             # Angular Frontend
├── docker-compose.yml
├── echo                # Echo Backend
├── node_modules
├── package-lock.json
└── package.json
```

### Echo Backend

```text
.
├── app
├── config
├── go.mod
├── go.sum
├── internal
├── main.go
├── tools
└── types
```

### Angular Frontend

```text
.
├── angular.json
├── node_modules
├── package-lock.json
├── package.json
├── public
├── src
├── tsconfig.app.json
├── tsconfig.json
└── tsconfig.spec.json
```

## QuickStart

```sh
# download dependencies
make deps

# run migration
make migrate

# generate secret
make secret

# start database container
make env

# start backend server
make dev

# start frontend server
make serve
```
