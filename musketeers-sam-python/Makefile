ENV ?= dev
MY_NAME ?= yun
STACK_NAME ?= $(ENV)-tech-challenge-${MY_NAME}
AWS_CLI = docker-compose run aws-cli
AWS_GOOGLE_AUTH = docker-compose run aws-google-auth

.env:
	@echo ".env not found, creating new .env out of .env.template"
	cp .env.template .env

auth: .env
	$(AWS_GOOGLE_AUTH)
.PHONY: auth

test:

deploy:
	$(AWS_CLI) make _deploy

# Helpers
shellAWS:
	$(AWS_CLI) bash

# Inner Targets
_package:
	aws cloudformation package 

_validate:
	aws cloudformation validate-template \
		--template-body file://template.yaml

_deploy: _validate
	aws cloudformation deploy \
		--template-file ./template.yaml \
		--stack-name $(STACK_NAME) \
		--capabilities CAPABILITY_NAMED_IAM \
		--parameter-overrides MyName=$(MY_NAME)

_describe:
	aws cloudformation describe-stack-events --stack-name $(STACK_NAME)

_remove:
	aws cloudformation delete-stack --stack-name $(STACK_NAME)
