PSQL_DB ?= postgres
PSQL_PASS ?= mysecretpass 
PSQL_CONTAINER ?= psql-`date '+%Y-%m-%d-%H-%M-%S'`

db-up:
	 docker run \
		 -p 5432:5432 \
		 --name $(PSQL_CONTAINER) \
		 -e POSTGRES_PASSWORD=$(PSQL_PASS) \
		 -d $(PSQL_DB)

.PHONY: db-up
