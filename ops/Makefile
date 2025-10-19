.PHONY: run-public-swagger stop-public-swagger compose-dev-up compose-dev-down

COMPOSE_DEV = docker compose -f docker-compose.dev.yml

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