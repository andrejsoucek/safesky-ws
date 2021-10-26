.PHONY: help
help:
	@awk 'BEGIN {FS = ":.*##"; printf "Usage: make <target>\n"} /^[A-Za-z-]+:.*?##/ { printf "  %-15s %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

.PHONY: build run


build: ## Build container
	docker build -t safesky-ws .

run: ## Run app
	docker run \
	--name safesky-ws \
	--network host \
	--restart always \
	-v "$$(pwd)/config.env:/var/lib/safesky-ws/config/config.env" \
	-d \
	safesky-ws \
	safesky-ws --env-file /var/lib/safesky-ws/config/config.env
