ENV ?= dev
ENV := $(or $(shell echo $$ENV),dev)
MIGRATE_URL := $(shell bash ./scripts/migrate.sh --url)

check-env:
ifeq ($(ENV),dev)
	@echo "Environment is set to dev"
else ifeq ($(ENV),test)
	@echo "Environment is set to test"
else ifeq ($(ENV),staging)
	@echo "Environment is set to staging"
else ifeq ($(ENV),prod)
	@echo "Environment is set to prod"
else
	$(error Unknown environment: $(ENV))
endif

migration:
	@echo "The script returned: $(MIGRATE_URL)"

migrate-ext:
	./bin/migrate create -ext sql -dir db/sql/migrations -seq $(sql_name)

migrate-up:
	@echo "The script returned: $(MIGRATE_URL)"
	./bin/migrate -database="$(MIGRATE_URL)" -path=./db/sql/migrations/ up

migration-down:
	@echo "The script returned: $(MIGRATE_URL)"
	./bin/migrate -database="$(MIGRATE_URL)" -path=./db/sql/migrations/ down

migration-force:
	@echo "The script returned: $(MIGRATE_URL)"
	./bin/migrate -database="$(MIGRATE_URL)" -path=./db/sql/migrations/ force $(version)


