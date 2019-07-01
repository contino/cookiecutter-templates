COMPOSE_RUN_COOKIECUTTER = docker-compose run --rm cookiecutter

all: generate test clean
.PHONY: all

generate:
	$(COMPOSE_RUN_COOKIECUTTER) cookiecutter --no-input cookiecutter-musketeers-sls-go
.PHONY: generate

test:
	cd musketeers-sls-go && make
.PHONY: test

shell:
	$(COMPOSE_RUN_COOKIECUTTER) sh
.PHONY: shell

clean:
	$(COMPOSE_RUN_COOKIECUTTER) rm -rf musketeers-sls-go
	docker-compose down --remove-orphans
.PHONY: clean