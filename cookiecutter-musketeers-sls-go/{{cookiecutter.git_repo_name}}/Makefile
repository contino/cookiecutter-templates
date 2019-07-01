BIN_DIR = bin

COMPOSE_RUN_GO = docker-compose run --rm golang
COMPOSE_RUN_SERVERLESS = docker-compose run --rm serverless
COMPOSE_RUN_AUTH = docker-compose run --rm auth

# creates .env with .env.template if no ENVFILE is defined
.env:
ifdef ENVFILE
	cp -f $(ENVFILE) .env
else
	cp -f .env.template .env
endif

# Public Targets
test:
	$(COMPOSE_RUN_GO) make _test

build:
	$(COMPOSE_RUN_GO) make _build

deploy:
	$(COMPOSE_RUN_GO) make _deploy

auth:
	$(COMPOSE_RUN_AUTH) make auth

shellGo:
	$(COMPOSE_RUN_GO) sh

# Inner Targets
_test:
	go test ./...

_build:
	go build -ldflags="-w -s" -o bin/http.go handlers/http/main.go

_deploy:
	cd handlers/http
	