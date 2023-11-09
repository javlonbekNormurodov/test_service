CURRENT_DIR=$(shell pwd)

-include .env

PSQL_CONTAINER_NAME?=postgres_container
PROJECT_NAME?=test_service
PSQL_URI?=postgres://postgres:postgres@localhost:5432/${PROJECT_NAME}?sslmode=disable

TAG=latest

.PHONY: sqlc
sqlc:
	sqlc generate

.PHONY: createdb
createdb:
	docker exec -it ${PSQL_CONTAINER_NAME} createdb -U postgres ${PROJECT_NAME}

.PHONY: execdb
execdb:
	docker exec -it ${PSQL_CONTAINER_NAME} psql -U postgres ${PROJECT_NAME}

.PHONY: dropdb
dropdb:
	docker exec -it ${PSQL_CONTAINER_NAME} dropdb -U postgres ${PROJECT_NAME}

.PHONY: migrate_up
migrate_up:
	goose -dir migrate/migrations postgres "postgres://postgres:admin123@localhost:5434/test_service?sslmode=disable" up



.PHONY: migrate_down
migrate_down:
	goose -dir migrate/migrations postgres "${PSQL_URI}" down

.PHONY: migrate_status
migrate_status:
	goose -s -dir migrate/migrations postgres "${PSQL_URI}" status

.PHONY: migrate_create
migrate_create:
	goose -s -dir migrate/migrations create ${NAME} sql

build_image:
	docker build --rm -t "${REGISTRY_NAME}/${PROJECT_NAME}:${TAG}" .

push_image:
	docker push "${REGISTRY_NAME}/${PROJECT_NAME}:${TAG}"

proto:
	rm -f generated/**/*.go
	rm -f doc/swagger/*.swagger.json
	mkdir -p generated
	protoc \
		--proto_path=protos --go_out=generated --go_opt=paths=source_relative \
		--go-grpc_out=generated --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=generated --grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=swagger_docs,use_allof_for_refs=true,disable_service_tags=true \
			protos/**/*.proto
