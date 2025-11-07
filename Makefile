.PHONY: dev-up dev-down buf-generate doc-static-identity doc-static-gen-calculation doc-static-gen-all run-public-swagger stop-public-swagger

COMPOSE_DEV = docker compose -f ops/docker-compose.dev.yml

run-public-swagger:
	@docker rm -f swagger-public 2>/dev/null || true
	docker run --name swagger-public --rm -d -p 8080:8080 \
	  -e SWAGGER_JSON=/foo/public-api.yaml \
	  -v "$$PWD/docs/contracts:/foo:ro" \
	  swaggerapi/swagger-ui
	@echo "Public API docs → http://localhost:8080"

stop-public-swagger:
	@docker rm -f swagger-public 2>/dev/null || true

dev-up:
	$(COMPOSE_DEV) up -d

dev-down:
	$(COMPOSE_DEV) down -v

buf-generate:
	@cd proto/calculation && buf generate
	@cd proto/identity    && buf generate

doc-static-gen-calculation:
	@protoc -I proto   --doc_out=docs/contracts   --doc_opt=html,calculation.html   proto/calculation/calculation.proto

doc-static-gen-identity:
	@protoc -I proto   --doc_out=docs/contracts   --doc_opt=html,identity.html   proto/identity/identity.proto

doc-static-gen-all:
	@protoc -I proto   --doc_out=docs/contracts   --doc_opt=html,calculation.html   proto/calculation/calculation.proto
	@protoc -I proto   --doc_out=docs/contracts   --doc_opt=html,identity.html   proto/identity/identity.proto
