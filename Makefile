PSQL_DB ?= postgres
PSQL_PASS ?= mysecretpass
PSQL_CONTAINER ?= psql-`date '+%Y-%m-%d-%H-%M-%S'`

DOCKER_REPO ?= pkprzekwas/fakeapp
DOCKER_TAG ?= `git rev-parse --short HEAD`

db-up:
	 docker run --name $(PSQL_CONTAINER) \
		 -p 5432:5432 \
		 -e POSTGRES_PASSWORD=$(PSQL_PASS) \
		 -d $(PSQL_DB)

build-image:
	docker build -t $(DOCKER_REPO):$(DOCKER_TAG) .

.PHONY: db-up
