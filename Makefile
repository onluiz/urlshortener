include .env

DB_URI="mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_SCHEMA)"

start:
	air

migrations-up:
	migrate -path $(MIGRATIONS_PATH) -database $(DB_URI)?query up

migrations-down:
	migrate -path $(MIGRATIONS_PATH) -database $(DB_URI)?query down
