.PHONY: run-multiple-swagger

run-multiple-swagger:
	docker run --rm -d -p 8080:8080 \
	  -e URLS_PRIMARY_NAME='API Gateway' \
	  -e 'URLS=[{"url":"/spec/internal-identity-api.yaml","name":"Identity"},{"url":"/spec/internal-purchase-api.yaml","name":"Purchase"},{"url":"/spec/internal-payments.yaml","name":"Payments"},{"url":"/spec/internal-calculation-api.yaml","name":"Calculation"},{"url":"/spec/internal-reports-api.yaml","name":"Reporting"},{"url":"/spec/internal-notification.yaml","name":"Notification"},{"url":"/spec/public-api.yaml","name":"API Gateway"}]' \
	  -v "$$PWD/docs/contracts:/usr/share/nginx/html/spec" \
	  swaggerapi/swagger-ui

compose-dev-up:
	docker compose -f docker-compose.dev.yml up -d

compose-dev-down:
	docker compose -f docker-compose.dev.yml down -v