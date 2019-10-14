PSQL_DB ?= postgres
PSQL_PASS ?= mysecretpass
PSQL_CONTAINER ?= psql-`date '+%Y-%m-%d-%H-%M-%S'`

DOCKER_REPO ?= pkprzekwas/fakeapp
DOCKER_TAG ?= `git rev-parse --short HEAD`

# Tested only on MacOS
LOCAL_IP_ADDR ?= `ifconfig | grep inet | awk '{ print $$2 }' | grep -E '^(192\.168\.)'`

db-up:
	 docker run \
		 --name $(PSQL_CONTAINER) \
		 -p 5432:5432 \
		 -e POSTGRES_PASSWORD=$(PSQL_PASS) \
		 -d $(PSQL_DB)

build-image:
	docker build -t $(DOCKER_REPO):$(DOCKER_TAG) .

run-container:
	docker run \
		-p 8080:8080 \
		-e POSTGRES_PASS=$(PSQL_PASS) \
		-e POSTGRES_HOST=$(LOCAL_IP_ADDR) \
		$(DOCKER_REPO):$(DOCKER_TAG)

.PHONY: db-up
