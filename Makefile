swagger:
	swag init --parseDependency --parseInternal
wire:
	wire ./internal

MIGRATIONS_DIR=migrations

DATETIME := $(shell date +%Y_%m_%d_%H%M%S)

migration:
	@if [ -z "$(name)" ]; then \
		echo "Error: You must specify a migration name. Usage: make migration name=your_migration_name"; \
		exit 1; \
	fi
	@mkdir -p $(MIGRATIONS_DIR)
	@touch $(MIGRATIONS_DIR)/$(DATETIME)_$(name).up.sql
	@touch $(MIGRATIONS_DIR)/$(DATETIME)_$(name).down.sql
	@echo "Created migration files: $(DATETIME)_$(name).up.sql and $(DATETIME)_$(name).down.sql"
