# Final Year Project

**Title**: *A Web-Based Scrum Management Tool for Daily Check-Ins with Smart Suggestions and Jira Ticket Mapping*

## Components

* **PostgreSQL** Database
* **Go Echo** Backend
* **TypeScript Angular** Frontend
* **Liquibase** Database Management

## Prerequisites

* **Go**
* **npm**
* **Podman** with **Podman Compose**
* **Liquibase** *Optional*

## Project Structure

```text
.
├── Dockerfile.ci.arm64         # docker image for ECR
├── Makefile
├── README.md
├── client                      # Angular Frontend
│   ├── angular.json
│   ├── dist
│   ├── node_modules
│   ├── package-lock.json
│   ├── package.json
│   ├── public
│   ├── src
│   ├── tsconfig.app.json
│   ├── tsconfig.json
│   └── tsconfig.spec.json
├── docker-compose.yml
├── package-lock.json
├── package.json
├── postgres                    # Liquibase
│   ├── change.yaml
│   ├── indexes.sql
│   ├── liquibase.properties
│   └── migration.sql
└── server                      # Echo Backend
    ├── app
    ├── config
    ├── go.mod
    ├── go.sum
    ├── internal
    ├── main.go
    ├── tools
    └── types
```

## QuickStart

```sh
# download dependencies
make deps

# run migration
make migrate

# OR

# run liquibase migration
make update

# generate secret
make secret

# build static frontend
make build-client

# start backend server
make dev

# start frontend server
make serve
```

## Usage

* Run

```sh
make build-client
make dev
```

* navigate to `localhost:8080`

## Deployment

* Update `client/src/environments/environments.prod.ts` to follow K8s setup
* Implement deployment CI/CD

### Example

1. Build docker image via GitLab pipeline
2. Publish docker image to ECR via GitLab pipeline
3. Update K8s to fetch new ECR image
